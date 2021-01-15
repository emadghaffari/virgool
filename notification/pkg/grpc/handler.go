package grpc

import (
	"context"
	"encoding/json"
	"fmt"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	endpoint "github.com/emadghaffari/virgool/notification/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/notification/pkg/grpc/pb"
)

// makeSMSHandler creates the handler logic
func makeSMSHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SMSEndpoint, decodeSMSRequest, encodeSMSResponse, options...)
}

// decodeSMSResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SMS request.
// TODO implement the decoder
func decodeSMSRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.SMSRequest)
	return endpoint.EmailRequest{To: rq.To, Body: rq.Body, Data: rq.Data}, nil
}

// encodeSMSResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSMSResponse(_ context.Context, r interface{}) (interface{}, error) {
	rsp := r.(endpoint.SMSResponse)

	if rsp.Err != nil {
		return pb.SMSReply{Message: rsp.Message, Status: rsp.Status}, rsp.Err
	}

	return pb.EmailReply{Message: rsp.Message, Status: rsp.Status}, nil
}
func (g *grpcServer) SMS(ctx context1.Context, req *pb.SMSRequest) (*pb.SMSReply, error) {
	_, rep, err := g.sMS.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SMSReply), nil
}

// makeSMSTHandler creates the handler logic
func makeSMSTHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SMSTEndpoint, decodeSMSTRequest, encodeSMSTResponse, options...)
}

// decodeSMSResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SMS request.
func decodeSMSTRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.SMSTRequest)

	vars := make(map[string]string, len(rq.Params))
	for _, k := range rq.Params {
		vars[k.Key] = k.Value
	}

	return endpoint.SMSTRequest{To: rq.To, Params: vars, Template: rq.Template, Data: rq.Data}, nil
}

// encodeSMSResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSMSTResponse(_ context.Context, r interface{}) (interface{}, error) {
	rsp := r.(endpoint.SMSTResponse)

	if rsp.Err != nil {
		return pb.SMSReply{Message: rsp.Message, Status: rsp.Status}, rsp.Err
	}

	return pb.EmailReply{Message: rsp.Message, Status: rsp.Status}, nil
}
func (g *grpcServer) SMST(ctx context1.Context, req *pb.SMSTRequest) (*pb.SMSTReply, error) {
	_, rep, err := g.sMST.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SMSTReply), nil
}

// makeEmailHandler creates the handler logic
func makeEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.EmailEndpoint, decodeEmailRequest, encodeEmailResponse, options...)
}

// decodeEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Email request.
func decodeEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.EmailRequest)
	return endpoint.EmailRequest{To: rq.To, Body: rq.Body, Data: rq.Data}, nil
}

// encodeEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	rsp := r.(endpoint.EmailResponse)

	if rsp.Err != nil {
		return pb.EmailReply{Message: rsp.Message, Status: rsp.Status}, rsp.Err
	}

	return pb.EmailReply{Message: rsp.Message, Status: rsp.Status}, nil
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
func decodeVerifyRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.VerifyRequest)
	return endpoint.VerifyRequest{Phone: rq.Phone, Code: rq.Code}, nil
}

// encodeVerifyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeVerifyResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.VerifyResponse)

	if rs.Err != nil {
		return &pb.VerifyReply{Message: rs.Message, Status: rs.Status}, fmt.Errorf("Error: %s", rs.Err)
	}

	b, err := json.Marshal(rs.Data)
	if err != nil {
		return &pb.VerifyReply{Message: err.Error(), Status: "ERROR"}, fmt.Errorf("Error: %s", rs.Err)
	}

	var any []*pb.Any
	any = append(any, &pb.Any{Key: "Message", Value: rs.Message})
	any = append(any, &pb.Any{Key: "Status", Value: rs.Status})
	any = append(any, &pb.Any{Key: "Data", Value: string(b)})

	return &pb.VerifyReply{Message: rs.Message, Status: rs.Status, Data: any}, nil
}

func (g *grpcServer) Verify(ctx context1.Context, req *pb.VerifyRequest) (*pb.VerifyReply, error) {
	_, rep, err := g.verify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.VerifyReply), nil
}
