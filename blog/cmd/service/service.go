package service

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	http1 "github.com/go-kit/kit/transport/http"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	group "github.com/oklog/oklog/pkg/group"
	"github.com/opentracing/opentracing-go"
	opentracinggo "github.com/opentracing/opentracing-go"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/database/elastic"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	"github.com/emadghaffari/virgool/blog/database/redis"
	"github.com/emadghaffari/virgool/blog/env"
	endpoint "github.com/emadghaffari/virgool/blog/pkg/endpoint"
	grpc "github.com/emadghaffari/virgool/blog/pkg/grpc"
	pb "github.com/emadghaffari/virgool/blog/pkg/grpc/pb"
	pkghttp "github.com/emadghaffari/virgool/blog/pkg/http"
	service "github.com/emadghaffari/virgool/blog/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("auth", flag.ExitOnError)

// Run func
func Run() {
	err := fs.Parse(os.Args[1:])
	if err != nil {
		logrus.Warn(err.Error())
	}

	// Read configs
	if err := initConfigs(); err != nil {
		if err := logger.Log("exit"); err != nil {
			logrus.Warn(err.Error())
		}
		return
	}

	// conf logger
	conf.ConfigureLogging(&conf.GlobalConfigs.Log)

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	closer, err := initJaeger()
	if err != nil {
		if err := logger.Log("exit"); err != nil {
			logrus.Warn(err.Error())
		}
		return
	}
	defer closer.Close()

	// connect to local database
	if err := initDatabase(); err != nil {
		if err := logger.Log("exit"); err != nil {
			logrus.Warn(err.Error())
		}
		return
	}

	// connect to local initRedis
	if err := initRedis(); err != nil {
		if err := logger.Log("exit"); err != nil {
			logrus.Warn(err.Error())
		}
		return
	}

	if err := elastic.Database.Connect(&conf.GlobalConfigs, logrus.StandardLogger()); err != nil {
		logrus.Error(err.Error())
		return
	}

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency

	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	if err := logger.Log("exit", g.Run()); err != nil {
		logrus.Warn(err.Error())
	}

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", conf.GlobalConfigs.GRPC.Port)
	if err != nil {
		if err := logger.Log("transport", "gRPC", "during", "Listen", "err", err); err != nil {
			logrus.Warn(err.Error())
		}
	}
	g.Add(func() error {
		if err := logger.Log("transport", "gRPC", "addr", conf.GlobalConfigs.GRPC.Port); err != nil {
			logrus.Warn(err.Error())
		}

		// UnaryInterceptor and OpenTracingServerInterceptor for tracer
		baseServer := grpc1.NewServer(
			grpc1.UnaryInterceptor(
				otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
			),
		)

		// reflection for evans
		reflection.Register(baseServer)

		pb.RegisterBlogServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}

// initHttpHandler func
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	httpHandler := pkghttp.NewHTTPHandler(endpoints, map[string][]http1.ServerOption{})
	httpListener, err := net.Listen("tcp", conf.GlobalConfigs.HTTP.Port)
	if err != nil {
		if err := logger.Log("transport", "HTTP", "during", "Listen", "err", err); err != nil {
			logrus.Warn(err.Error())
		}
	}
	g.Add(func() error {
		if err := logger.Log("transport", "HTTP", "addr", conf.GlobalConfigs.HTTP.Port); err != nil {
			logrus.Warn(err.Error())
		}
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = append(mw, service.LoggingMiddleware(logger))
	// Append your middleware here

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
	// 	Help:      "Request duration in seconds.",
	// 	Name:      "request_duration_seconds",
	// 	Namespace: "example",
	// 	Subsystem: "blog",
	// }, []string{"method", "success"})
	addEndpointMiddlewareToAllMethods(mw, endpoint.LoggingMiddleware(logger))
	// Add you endpoint middleware here

	return
}
func initMetricsEndpoint(g *group.Group) {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", conf.GlobalConfigs.DEBUG.Port)
	if err != nil {
		if err := logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err); err != nil {
			logrus.Warn(err.Error())
		}
	}
	g.Add(func() error {
		if err := logger.Log("transport", "debug/HTTP", "addr", conf.GlobalConfigs.DEBUG.Port); err != nil {
			logrus.Warn(err.Error())
		}
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initDatabase() error {
	return mysql.Database.Connect(&conf.GlobalConfigs, logrus.New())
}

func initConfigs() error {
	// Current working directory
	dir, err := os.Getwd()
	if err != nil {
		logrus.Warn(err.Error())
	}
	// read from file
	return env.LoadGlobalConfiguration(dir + "/config.yaml")
}

func initJaeger() (io.Closer, error) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: conf.GlobalConfigs.Service.Name,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           conf.GlobalConfigs.Jaeger.LogSpans,
			LocalAgentHostPort: conf.GlobalConfigs.Jaeger.HostPort,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	var closer io.Closer
	var err error
	tracer, closer, err = cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		if err := logger.Log("during", "Listen", "jaeger", "err", err); err != nil {
			logrus.Warn(err.Error())
		}
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}

func initRedis() error {
	return redis.Database.Connect(&conf.GlobalConfigs)
}
