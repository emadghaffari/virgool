package grpc

import (
	"context"
	"errors"
	endpoint "github.com/emadghaffari/virgool/blog/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/blog/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCreatePostHandler creates the handler logic
func makeCreatePostHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreatePostEndpoint, decodeCreatePostRequest, encodeCreatePostResponse, options...)
}

// decodeCreatePostResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreatePost request.
// TODO implement the decoder
func decodeCreatePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeCreatePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreatePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) CreatePost(ctx context1.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	_, rep, err := g.createPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreatePostReply), nil
}

// makeUpdatePostHandler creates the handler logic
func makeUpdatePostHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdatePostEndpoint, decodeUpdatePostRequest, encodeUpdatePostResponse, options...)
}

// decodeUpdatePostResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdatePost request.
// TODO implement the decoder
func decodeUpdatePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeUpdatePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdatePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) UpdatePost(ctx context1.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	_, rep, err := g.updatePost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdatePostReply), nil
}

// makeGetPostHandler creates the handler logic
func makeGetPostHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetPostEndpoint, decodeGetPostRequest, encodeGetPostResponse, options...)
}

// decodeGetPostResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetPost request.
// TODO implement the decoder
func decodeGetPostRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeGetPostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetPostResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) GetPost(ctx context1.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	_, rep, err := g.getPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetPostReply), nil
}

// makeDeletePostHandler creates the handler logic
func makeDeletePostHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeletePostEndpoint, decodeDeletePostRequest, encodeDeletePostResponse, options...)
}

// decodeDeletePostResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeletePost request.
// TODO implement the decoder
func decodeDeletePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeDeletePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeletePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) DeletePost(ctx context1.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	_, rep, err := g.deletePost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeletePostReply), nil
}

// makeCreateTagHandler creates the handler logic
func makeCreateTagHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateTagEndpoint, decodeCreateTagRequest, encodeCreateTagResponse, options...)
}

// decodeCreateTagResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateTag request.
// TODO implement the decoder
func decodeCreateTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeCreateTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) CreateTag(ctx context1.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	_, rep, err := g.createTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateTagReply), nil
}

// makeGetTagHandler creates the handler logic
func makeGetTagHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTagEndpoint, decodeGetTagRequest, encodeGetTagResponse, options...)
}

// decodeGetTagResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTag request.
// TODO implement the decoder
func decodeGetTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeGetTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) GetTag(ctx context1.Context, req *pb.GetTagRequest) (*pb.GetTagReply, error) {
	_, rep, err := g.getTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTagReply), nil
}

// makeUpdateTagHandler creates the handler logic
func makeUpdateTagHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTagEndpoint, decodeUpdateTagRequest, encodeUpdateTagResponse, options...)
}

// decodeUpdateTagResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTag request.
// TODO implement the decoder
func decodeUpdateTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeUpdateTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTag(ctx context1.Context, req *pb.UpdateTagRequest) (*pb.UpdateTagReply, error) {
	_, rep, err := g.updateTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTagReply), nil
}

// makeDeleteTagHandler creates the handler logic
func makeDeleteTagHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTagEndpoint, decodeDeleteTagRequest, encodeDeleteTagResponse, options...)
}

// decodeDeleteTagResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTag request.
// TODO implement the decoder
func decodeDeleteTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeDeleteTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTag(ctx context1.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagReply, error) {
	_, rep, err := g.deleteTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTagReply), nil
}

// makeUploadHandler creates the handler logic
func makeUploadHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UploadEndpoint, decodeUploadRequest, encodeUploadResponse, options...)
}

// decodeUploadResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Upload request.
// TODO implement the decoder
func decodeUploadRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Decoder is not impelemented")
}

// encodeUploadResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUploadResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Blog' Encoder is not impelemented")
}
func (g *grpcServer) Upload(ctx context1.Context, req *pb.UploadRequest) (*pb.UploadReply, error) {
	_, rep, err := g.upload.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UploadReply), nil
}
