package grpc

import (
	"context"
	"errors"
	endpoint "github.com/emadghaffari/virgool/notification/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/notification/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeSMSHandler creates the handler logic
func makeSMSHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SMSEndpoint, decodeSMSRequest, encodeSMSResponse, options...)
}

// decodeSMSResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SMS request.
// TODO implement the decoder
func decodeSMSRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Decoder is not impelemented")
}

// encodeSMSResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSMSResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Encoder is not impelemented")
}
func (g *grpcServer) SMS(ctx context1.Context, req *pb.SMSRequest) (*pb.SMSReply, error) {
	_, rep, err := g.sMS.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SMSReply), nil
}

// makeEmailHandler creates the handler logic
func makeEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.EmailEndpoint, decodeEmailRequest, encodeEmailResponse, options...)
}

// decodeEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Email request.
// TODO implement the decoder
func decodeEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Decoder is not impelemented")
}

// encodeEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Encoder is not impelemented")
}
func (g *grpcServer) Email(ctx context1.Context, req *pb.EmailRequest) (*pb.EmailReply, error) {
	_, rep, err := g.email.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.EmailReply), nil
}

// makeVerifyHandler creates the handler logic
func makeVerifyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.VerifyEndpoint, decodeVerifyRequest, encodeVerifyResponse, options...)
}

// decodeVerifyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Verify request.
// TODO implement the decoder
func decodeVerifyRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Decoder is not impelemented")
}

// encodeVerifyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeVerifyResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notification' Encoder is not impelemented")
}
func (g *grpcServer) Verify(ctx context1.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	_, rep, err := g.verify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.VerifyReply), nil
}
