package endpoint

import (
	"bytes"
	"context"

	model "github.com/emadghaffari/virgool/blog/model"
	service "github.com/emadghaffari/virgool/blog/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreatePostRequest collects the request parameters for the CreatePost method.
type CreatePostRequest struct {
	UserID      uint64           `json:"user_id"`
	Title       string           `json:"title"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Text        string           `json:"text"`
	Params      []*model.Query   `json:"params"`
	Medias      []int64          `json:"medias"`
	Tags        []int64          `json:"tags"`
	Status      model.StatusPost `json:"status"`
	Token       string           `json:"token"`
}

// CreatePostResponse collects the response parameters for the CreatePost method.
type CreatePostResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeCreatePostEndpoint returns an endpoint that invokes CreatePost on the service.
func MakeCreatePostEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePostRequest)
		message, status, err := s.CreatePost(ctx, req.UserID, req.Title, req.Slug, req.Description, req.Text, req.Params, req.Medias, req.Tags, req.Status, req.Token)
		return CreatePostResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r CreatePostResponse) Failed() error {
	return r.Err
}

// UpdatePostRequest collects the request parameters for the UpdatePost method.
type UpdatePostRequest struct {
	Title       string           `json:"title"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Text        string           `json:"text"`
	Params      []*model.Query   `json:"params"`
	Medias      []int64          `json:"medias"`
	Tags        []int64          `json:"tags"`
	Status      model.StatusPost `json:"status"`
	Token       string           `json:"token"`
}

// UpdatePostResponse collects the response parameters for the UpdatePost method.
type UpdatePostResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeUpdatePostEndpoint returns an endpoint that invokes UpdatePost on the service.
func MakeUpdatePostEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePostRequest)
		message, status, err := s.UpdatePost(ctx, req.Title, req.Slug, req.Description, req.Text, req.Params, req.Medias, req.Tags, req.Status, req.Token)
		return UpdatePostResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdatePostResponse) Failed() error {
	return r.Err
}

// GetPostRequest collects the request parameters for the GetPost method.
type GetPostRequest struct {
	Must   []*model.Query `json:"must"`
	Should []*model.Query `json:"should"`
	Not    []*model.Query `json:"not"`
	Filter []*model.Query `json:"filter"`
	Token  string         `json:"token"`
}

// GetPostResponse collects the response parameters for the GetPost method.
type GetPostResponse struct {
	Posts   []model.Post `json:"posts"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
	Err     error        `json:"err"`
}

// MakeGetPostEndpoint returns an endpoint that invokes GetPost on the service.
func MakeGetPostEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPostRequest)
		posts, message, status, err := s.GetPost(ctx, req.Must, req.Should, req.Not, req.Filter, req.Token)
		return GetPostResponse{
			Err:     err,
			Message: message,
			Posts:   posts,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r GetPostResponse) Failed() error {
	return r.Err
}

// DeletePostRequest collects the request parameters for the DeletePost method.
type DeletePostRequest struct {
	Filter []*model.Query `json:"filter"`
	Token  string         `json:"token"`
}

// DeletePostResponse collects the response parameters for the DeletePost method.
type DeletePostResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeDeletePostEndpoint returns an endpoint that invokes DeletePost on the service.
func MakeDeletePostEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePostRequest)
		message, status, err := s.DeletePost(ctx, req.Filter, req.Token)
		return DeletePostResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r DeletePostResponse) Failed() error {
	return r.Err
}

// CreateTagRequest collects the request parameters for the CreateTag method.
type CreateTagRequest struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// CreateTagResponse collects the response parameters for the CreateTag method.
type CreateTagResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeCreateTagEndpoint returns an endpoint that invokes CreateTag on the service.
func MakeCreateTagEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTagRequest)
		message, status, err := s.CreateTag(ctx, req.Name, req.Token)
		return CreateTagResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateTagResponse) Failed() error {
	return r.Err
}

// GetTagRequest collects the request parameters for the GetTag method.
type GetTagRequest struct {
	Filter []*model.Query `json:"filter"`
	Token  string         `json:"token"`
}

// GetTagResponse collects the response parameters for the GetTag method.
type GetTagResponse struct {
	Tags    []*model.Tag `json:"tags"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
	Err     error        `json:"err"`
}

// MakeGetTagEndpoint returns an endpoint that invokes GetTag on the service.
func MakeGetTagEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTagRequest)
		tags, message, status, err := s.GetTag(ctx, req.Filter, req.Token)
		return GetTagResponse{
			Err:     err,
			Message: message,
			Status:  status,
			Tags:    tags,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTagResponse) Failed() error {
	return r.Err
}

// UpdateTagRequest collects the request parameters for the UpdateTag method.
type UpdateTagRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
	Token   string `json:"token"`
}

// UpdateTagResponse collects the response parameters for the UpdateTag method.
type UpdateTagResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeUpdateTagEndpoint returns an endpoint that invokes UpdateTag on the service.
func MakeUpdateTagEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTagRequest)
		message, status, err := s.UpdateTag(ctx, req.OldName, req.NewName, req.Token)
		return UpdateTagResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTagResponse) Failed() error {
	return r.Err
}

// DeleteTagRequest collects the request parameters for the DeleteTag method.
type DeleteTagRequest struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// DeleteTagResponse collects the response parameters for the DeleteTag method.
type DeleteTagResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeDeleteTagEndpoint returns an endpoint that invokes DeleteTag on the service.
func MakeDeleteTagEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTagRequest)
		message, status, err := s.DeleteTag(ctx, req.Name, req.Token)
		return DeleteTagResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteTagResponse) Failed() error {
	return r.Err
}

// UploadRequest collects the request parameters for the Upload method.
type UploadRequest struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	FileType    string       `json:"file_type"`
	File        bytes.Buffer `json:"file"`
	Token       string       `json:"token"`
}

// UploadResponse collects the response parameters for the Upload method.
type UploadResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Err     error  `json:"err"`
}

// MakeUploadEndpoint returns an endpoint that invokes Upload on the service.
func MakeUploadEndpoint(s service.BlogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadRequest)
		message, status, err := s.Upload(ctx, req.Title, req.Description, req.FileType, req.File, req.Token)
		return UploadResponse{
			Err:     err,
			Message: message,
			Status:  status,
		}, nil
	}
}

// Failed implements Failer.
func (r UploadResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreatePost implements Service. Primarily useful in a client.
func (e Endpoints) CreatePost(ctx context.Context, userID uint64, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost, token string) (message string, status string, err error) {
	request := CreatePostRequest{
		Description: description,
		Medias:      medias,
		Params:      params,
		Slug:        slug,
		Status:      Status,
		Tags:        Tags,
		Text:        text,
		Title:       title,
		Token:       token,
		UserID:      userID,
	}
	response, err := e.CreatePostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreatePostResponse).Message, response.(CreatePostResponse).Status, response.(CreatePostResponse).Err
}

// UpdatePost implements Service. Primarily useful in a client.
func (e Endpoints) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost, token string) (message string, status string, err error) {
	request := UpdatePostRequest{
		Description: description,
		Medias:      medias,
		Params:      params,
		Slug:        slug,
		Status:      Status,
		Tags:        Tags,
		Text:        text,
		Title:       title,
		Token:       token,
	}
	response, err := e.UpdatePostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdatePostResponse).Message, response.(UpdatePostResponse).Status, response.(UpdatePostResponse).Err
}

// GetPost implements Service. Primarily useful in a client.
func (e Endpoints) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query, token string) (posts []model.Post, message string, status string, err error) {
	request := GetPostRequest{
		Filter: filter,
		Must:   must,
		Not:    not,
		Should: should,
		Token:  token,
	}
	response, err := e.GetPostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetPostResponse).Posts, response.(GetPostResponse).Message, response.(GetPostResponse).Status, response.(GetPostResponse).Err
}

// DeletePost implements Service. Primarily useful in a client.
func (e Endpoints) DeletePost(ctx context.Context, filter []*model.Query, token string) (message string, status string, err error) {
	request := DeletePostRequest{
		Filter: filter,
		Token:  token,
	}
	response, err := e.DeletePostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeletePostResponse).Message, response.(DeletePostResponse).Status, response.(DeletePostResponse).Err
}

// CreateTag implements Service. Primarily useful in a client.
func (e Endpoints) CreateTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	request := CreateTagRequest{
		Name:  name,
		Token: token,
	}
	response, err := e.CreateTagEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateTagResponse).Message, response.(CreateTagResponse).Status, response.(CreateTagResponse).Err
}

// GetTag implements Service. Primarily useful in a client.
func (e Endpoints) GetTag(ctx context.Context, filter []*model.Query, token string) (tags []*model.Tag, message string, status string, err error) {
	request := GetTagRequest{
		Filter: filter,
		Token:  token,
	}
	response, err := e.GetTagEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTagResponse).Tags, response.(GetTagResponse).Message, response.(GetTagResponse).Status, response.(GetTagResponse).Err
}

// UpdateTag implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTag(ctx context.Context, oldName string, newName string, token string) (message string, status string, err error) {
	request := UpdateTagRequest{
		NewName: newName,
		OldName: oldName,
		Token:   token,
	}
	response, err := e.UpdateTagEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTagResponse).Message, response.(UpdateTagResponse).Status, response.(UpdateTagResponse).Err
}

// DeleteTag implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	request := DeleteTagRequest{
		Name:  name,
		Token: token,
	}
	response, err := e.DeleteTagEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTagResponse).Message, response.(DeleteTagResponse).Status, response.(DeleteTagResponse).Err
}

// Upload implements Service. Primarily useful in a client.
func (e Endpoints) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {
	request := UploadRequest{
		Description: description,
		File:        file,
		FileType:    fileType,
		Title:       title,
		Token:       token,
	}
	response, err := e.UploadEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UploadResponse).Message, response.(UploadResponse).Status, response.(UploadResponse).Err
}
