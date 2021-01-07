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
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/mysql"
	"github.com/emadghaffari/virgool/auth/env"
	endpoint "github.com/emadghaffari/virgool/auth/pkg/endpoint"
	grpc "github.com/emadghaffari/virgool/auth/pkg/grpc"
	pb "github.com/emadghaffari/virgool/auth/pkg/grpc/pb"
	service "github.com/emadghaffari/virgool/auth/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("auth", flag.ExitOnError)
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var thriftAddr = fs.String("thrift-addr", ":8083", "Thrift listen address")
var thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
var thriftBuffer = fs.Int("thrift-buffer", 0, "0 for unbuffered")
var thriftFramed = fs.Bool("thrift-framed", false, "true to enable framing")

// Run func
func Run() {

	// Read configs
	if err := initConfigs(); err != nil {
		logger.Log("exit")
		return
	}

	conf.ConfigureLogging(&conf.GlobalConfigs.Log)

	// connect to local database
	if err := initDatabase(); err != nil {
		logger.Log("exit")
		return
	}

	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	closer, err := initJaeger()
	if err != nil {
		logger.Log("exit")
		return
	}
	defer closer.Close()

	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}

// initHttpHandler func
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", conf.GlobalConfigs.GRPC.Port)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", conf.GlobalConfigs.GRPC.Port)

		// UnaryInterceptor and OpenTracingServerInterceptor for tracer
		baseServer := grpc1.NewServer(
			grpc1.UnaryInterceptor(
				otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
			),
		)

		// reflection for evans
		reflection.Register(baseServer)

		pb.RegisterAuthServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	// mw = []service.Middleware{}
	// mw = addDefaultServiceMiddleware(logger, mw)
	// Append your middleware here

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
	// 	Help:      "Request duration in seconds.",
	// 	Name:      "request_duration_seconds",
	// 	Namespace: "example",
	// 	Subsystem: "auth",
	// }, []string{"method", "success"})
	// addDefaultEndpointMiddleware(logger, duration, mw)
	// Add you endpoint middleware here

	return
}
func initMetricsEndpoint(g *group.Group) {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", conf.GlobalConfigs.DEBUG.Port)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", conf.GlobalConfigs.DEBUG.Port)
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
	return mysql.Database.Connect(&conf.GlobalConfigs, conf.Logger)
}

// FIXME fix the config file path
func initConfigs() error {
	return env.LoadGlobalConfiguration("auth/config.yaml")
	// os.Getenv("config_file")
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
		logger.Log("during", "Listen", "jaeger", "err", err)
		return nil, err
	}

	opentracinggo.SetGlobalTracer(tracer)
	return closer, nil
}
