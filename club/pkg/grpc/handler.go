package grpc

import (
	"context"
	"errors"
	endpoint "github.com/emadghaffari/virgool/club/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/club/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Get request.
// TODO implement the decoder
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Club' Decoder is not impelemented")
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Club' Encoder is not impelemented")
}
func (g *grpcServer) Get(ctx context1.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}
