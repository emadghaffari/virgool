package grpc

import (
	"bytes"
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	"github.com/emadghaffari/virgool/blog/model"
	endpoint "github.com/emadghaffari/virgool/blog/pkg/endpoint"
	pb "github.com/emadghaffari/virgool/blog/pkg/grpc/pb"
)

// makeCreatePostHandler creates the handler logic
func makeCreatePostHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreatePostEndpoint, decodeCreatePostRequest, encodeCreatePostResponse, options...)
}

// decodeCreatePostResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreatePost request.
func decodeCreatePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.CreatePostRequest)

	var params []*model.Query
	for _, v := range rq.Params {
		params = append(params, &model.Query{Name: v.Key, Value: v.Value})
	}

	return endpoint.CreatePostRequest{
		UserID:      rq.UserID,
		Title:       rq.Title,
		Slug:        rq.Slug,
		Description: rq.Description,
		Text:        rq.Text,
		Params:      params,
		Medias:      rq.Medias,
		Tags:        rq.Tags,
		Status:      model.ScanStatus(rq.Status),
		Token:       rq.Token,
	}, nil
}

// encodeCreatePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeCreatePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.CreatePostResponse)

	if rs.Err != nil {
		return &pb.CreatePostReply{}, rs.Err
	}

	return &pb.CreatePostReply{
		Message: rs.Message,
		Status:  rs.Status,
	}, nil

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
func decodeUpdatePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.UpdatePostRequest)

	var params []*model.Query
	for _, v := range rq.Params {
		params = append(params, &model.Query{Name: v.Key, Value: v.Value})
	}

	return endpoint.UpdatePostRequest{
		Title:       rq.Title,
		Slug:        rq.Slug,
		Description: rq.Description,
		Text:        rq.Text,
		Params:      params,
		Medias:      rq.Medias,
		Tags:        rq.Tags,
		Status:      model.ScanStatus(rq.Status),
		Token:       rq.Token,
	}, nil
}

// encodeUpdatePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeUpdatePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.UpdatePostResponse)

	if rs.Err != nil {
		return &pb.UpdatePostReply{}, rs.Err
	}

	return &pb.UpdatePostReply{
		Message: rs.Message,
		Status:  rs.Status,
	}, nil
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
func decodeGetPostRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.GetPostRequest)

	var should []*model.Query
	for _, v := range rq.Should {
		should = append(should, &model.Query{Name: v.Key, Value: v.Value})
	}

	var must []*model.Query
	for _, v := range rq.Must {
		must = append(must, &model.Query{Name: v.Key, Value: v.Value})
	}

	var not []*model.Query
	for _, v := range rq.Not {
		not = append(not, &model.Query{Name: v.Key, Value: v.Value})
	}

	var filter []*model.Query
	for _, v := range rq.Filter {
		filter = append(filter, &model.Query{Name: v.Key, Value: v.Value})
	}

	return endpoint.GetPostRequest{
		Must:   must,
		Should: should,
		Not:    not,
		Filter: filter,
		Token:  rq.Token,
	}, nil
}

// encodeGetPostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetPostResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.GetPostResponse)

	if rs.Err != nil {
		return &pb.GetPostReply{}, rs.Err
	}

	var posts []*pb.Posts

	for _, v := range rs.Posts {
		var media []*pb.Media
		var query []*pb.Query
		var tag []*pb.Tag

		for _, m := range v.Media {
			media = append(media, &pb.Media{Url: m.URL, Type: m.Type, Title: *m.Title, Description: *m.Description})
		}

		for _, p := range v.Params {
			query = append(query, &pb.Query{Key: p.Name, Value: p.Value})
		}

		for _, t := range v.Tags {
			tag = append(tag, &pb.Tag{Name: t.Name})
		}

		posts = append(posts, &pb.Posts{
			Title:       v.Title,
			Slug:        v.Slug,
			Description: v.Description,
			Text:        v.Text,
			Status:      string(v.Status),
			Medias:      media,
			Params:      query,
			Tags:        tag,
			PublishedAT: v.PublishedAT.UTC().String(),
		})
	}

	return &pb.GetPostReply{
		Posts:   posts,
		Message: rs.Message,
		Status:  rs.Status,
	}, nil
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
func decodeDeletePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.DeletePostRequest)

	var filter []*model.Query
	for _, v := range rq.Filter {
		filter = append(filter, &model.Query{Name: v.Key, Value: v.Value})
	}

	return endpoint.DeletePostRequest{
		Filter: filter,
		Token:  rq.Token,
	}, nil
}

// encodeDeletePostResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeDeletePostResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.DeletePostResponse)

	if rs.Err != nil {
		return &pb.DeletePostReply{}, rs.Err
	}

	return &pb.DeletePostReply{Message: rs.Message, Status: rs.Status}, nil
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
func decodeCreateTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.CreateTagRequest)

	return endpoint.CreateTagRequest{
		Name:  rq.Name,
		Token: rq.Token,
	}, nil
}

// encodeCreateTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeCreateTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.CreateTagResponse)

	if rs.Err != nil {
		return &pb.CreateTagReply{}, rs.Err
	}

	return &pb.CreateTagReply{Message: rs.Message, Status: rs.Status}, nil
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
func decodeGetTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.GetTagRequest)

	var filter []*model.Query
	for _, v := range rq.Filter {
		filter = append(filter, &model.Query{Name: v.Key, Value: v.Value})
	}

	return endpoint.GetTagRequest{
		Filter: filter,
		Token:  rq.Token,
	}, nil
}

// encodeGetTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.GetTagResponse)

	if rs.Err != nil {
		return &pb.GetTagReply{}, rs.Err
	}

	return &pb.GetTagReply{Message: rs.Message, Status: rs.Status}, nil
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
func decodeUpdateTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.UpdateTagRequest)

	return endpoint.UpdateTagRequest{
		OldName: rq.OldName,
		NewName: rq.NewName,
		Token:   rq.Token,
	}, nil
}

// encodeUpdateTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeUpdateTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.UpdateTagResponse)

	if rs.Err != nil {
		return &pb.UpdateTagReply{}, rs.Err
	}

	return &pb.UpdateTagReply{Message: rs.Message, Status: rs.Status}, nil
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
func decodeDeleteTagRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.DeleteTagRequest)

	return endpoint.DeleteTagRequest{
		Name:  rq.Name,
		Token: rq.Token,
	}, nil
}

// encodeDeleteTagResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeDeleteTagResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.DeleteTagResponse)

	if rs.Err != nil {
		return &pb.DeleteTagReply{}, rs.Err
	}

	return &pb.DeleteTagReply{Message: rs.Message, Status: rs.Status}, nil
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
func decodeUploadRequest(_ context.Context, r interface{}) (interface{}, error) {
	rq := r.(*pb.UploadRequest)

	return endpoint.UploadRequest{
		Title:       rq.Title,
		Description: rq.Description,
		FileType:    rq.FileType,
		File:        *bytes.NewBuffer(rq.Data),
		Token:       rq.Token,
	}, nil
}

// encodeUploadResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeUploadResponse(_ context.Context, r interface{}) (interface{}, error) {
	rs := r.(endpoint.UploadResponse)

	if rs.Err != nil {
		return &pb.UploadReply{}, rs.Err
	}

	return &pb.UploadReply{Message: rs.Message, Status: rs.Status}, nil
}
func (g *grpcServer) Upload(ctx context1.Context, req *pb.UploadRequest) (*pb.UploadReply, error) {
	_, rep, err := g.upload.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UploadReply), nil
}
