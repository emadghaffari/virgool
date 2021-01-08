package grpc

import (
	"context"
	"fmt"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	endpoint "github.com/emadghaffari/virgool/auth/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/auth/pkg/grpc/pb"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)
}

// decodeRegisterResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Register request.
func decodeRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.RegisterRequest)
	return endpoint.RegisterRequest{
		Username: rq.Username,
		Password: rq.Password,
		Name:     rq.Name,
		LastName: rq.LastName,
		Phone:    rq.Phone,
		Email:    rq.Email,
	}, nil
}

// encodeRegisterResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.RegisterResponse)

	if rs.Err != nil {
		return &pb.RegisterReply{
			Message: rs.Err.Error(),
			Status:  "ERROR",
		}, fmt.Errorf("Message: %v With Status: %v", rs.Err.Error(), rs.Err)
	}

	return &pb.RegisterReply{
		Message: fmt.Sprintf("Hi %s We Send a SMS for verify your Phone!", rs.Response.Username),
		Status:  "SUCCESS",
	}, nil
}

func (g *grpcServer) Register(ctx context1.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	_, rep, err := g.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterReply), nil
}

// makeLoginUPHandler creates the handler logic
func makeLoginUPHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginUPEndpoint, decodeLoginUPRequest, encodeLoginUPResponse, options...)
}

// decodeLoginUPResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain LoginUP request.
func decodeLoginUPRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.LoginUPRequest)
	return endpoint.LoginUPRequest{
		Username: rq.Username,
		Password: rq.Password,
	}, nil
}

// encodeLoginUPResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeLoginUPResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.LoginUPResponse)

	if rs.Err != nil {
		return &pb.LoginUPReply{}, fmt.Errorf("Error: %s ", rs.Err.Error())
	}

	permissions := make([]*pb.Permission, len(rs.Response.Role.Permissions))
	for _, v := range rs.Response.Role.Permissions {
		permissions = append(permissions, &pb.Permission{Name: v.Name})
	}

	return &pb.LoginUPReply{
		Username: rs.Response.Username,
		Name:     rs.Response.Name,
		LastName: rs.Response.LastName,
		Phone:    rs.Response.Phone,
		Email:    rs.Response.Email,
		Token:    rs.Response.Token,
		Role: &pb.Role{
			Name:        rs.Response.Role.Name,
			Permissions: permissions,
		},
	}, nil
}
func (g *grpcServer) LoginUP(ctx context1.Context, req *pb.LoginUPRequest) (*pb.LoginUPReply, error) {
	_, rep, err := g.loginUP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginUPReply), nil
}

// makeLoginPHandler creates the handler logic
func makeLoginPHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginPEndpoint, decodeLoginPRequest, encodeLoginPResponse, options...)
}

// decodeLoginPResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain LoginP request.
func decodeLoginPRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.LoginPRequest)

	return endpoint.LoginPRequest{
		Phone: rq.Phone,
	}, nil
}

// encodeLoginPResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeLoginPResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.LoginPResponse)

	if rs.Err != nil {
		return &pb.LoginPReply{}, fmt.Errorf("Error: %s ", rs.Err.Error())
	}

	return &pb.LoginPReply{
		Message: fmt.Sprintf("Hi %s We Send a SMS for verify your Phone!", rs.Response.Username),
		Status:  "SUCCESS",
	}, nil
}
func (g *grpcServer) LoginP(ctx context1.Context, req *pb.LoginPRequest) (*pb.LoginPReply, error) {
	_, rep, err := g.loginP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginPReply), nil
}

// makeVerifyHandler creates the handler logic
func makeVerifyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.VerifyEndpoint, decodeVerifyRequest, encodeVerifyResponse, options...)
}

// decodeVerifyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Verify request.
func decodeVerifyRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.VerifyRequest)

	return endpoint.VerifyRequest{
		Token: rq.Token,
		Type:  rq.Type,
		Code:  rq.Code,
	}, nil
}

// encodeVerifyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeVerifyResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.VerifyResponse)

	if rs.Err != nil {
		return &pb.VerifyReply{}, fmt.Errorf("Error: %s ", rs.Err.Error())
	}

	permissions := make([]*pb.Permission, len(rs.Response.Role.Permissions))
	for _, v := range rs.Response.Role.Permissions {
		permissions = append(permissions, &pb.Permission{Name: v.Name})
	}

	return &pb.VerifyReply{
		Username: rs.Response.Username,
		Name:     rs.Response.Name,
		LastName: rs.Response.LastName,
		Phone:    rs.Response.Phone,
		Email:    rs.Response.Email,
		Token:    rs.Response.Token,
		Role: &pb.Role{
			Name:        rs.Response.Role.Name,
			Permissions: permissions,
		},
	}, nil
}
func (g *grpcServer) Verify(ctx context1.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	_, rep, err := g.verify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.VerifyReply), nil
}
