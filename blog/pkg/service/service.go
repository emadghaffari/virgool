package service

import (
	"bytes"
	"context"

	"github.com/emadghaffari/virgool/blog/model"
)

// BlogService describes the service.
type BlogService interface {
	// CRUD Posts
	CreatePost(ctx context.Context, userID uint64, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error)
	UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error)
	GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query) (posts []model.Post, message string, status string, err error)
	DeletePost(ctx context.Context, filter []*model.Query) (message string, status string, err error)

	// CRUD tags
	CreateTag(ctx context.Context, name string) (message string, status string, err error)
	GetTag(ctx context.Context, filter []*model.Query) (tags []*model.Tag, message string, status string, err error)
	UpdateTag(ctx context.Context, oldName, newName string) (message string, status string, err error)
	DeleteTag(ctx context.Context, name string) (message string, status string, err error)

	// CRUD media
	// multipart.File
	Upload(ctx context.Context, title, description, fileType string, file bytes.Buffer) (message string, status string, err error)
}

type basicBlogService struct{}

func (b *basicBlogService) CreatePost(ctx context.Context, userID uint64, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error) {
	// TODO implement the business logic of CreatePost
	return message, status, err
}
func (b *basicBlogService) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error) {
	// TODO implement the business logic of UpdatePost
	return message, status, err
}
func (b *basicBlogService) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query) (posts []model.Post, message string, status string, err error) {
	// TODO implement the business logic of GetPost
	return posts, message, status, err
}
func (b *basicBlogService) DeletePost(ctx context.Context, filter []*model.Query) (message string, status string, err error) {
	// TODO implement the business logic of DeletePost
	return message, status, err
}
func (b *basicBlogService) CreateTag(ctx context.Context, name string) (message string, status string, err error) {
	// TODO implement the business logic of CreateTag
	return message, status, err
}
func (b *basicBlogService) GetTag(ctx context.Context, filter []*model.Query) (tags []*model.Tag, message string, status string, err error) {
	// TODO implement the business logic of GetTag
	return tags, message, status, err
}
func (b *basicBlogService) UpdateTag(ctx context.Context, oldName string, newName string) (message string, status string, err error) {
	// TODO implement the business logic of UpdateTag
	return message, status, err
}
func (b *basicBlogService) DeleteTag(ctx context.Context, name string) (message string, status string, err error) {
	// TODO implement the business logic of DeleteTag
	return message, status, err
}
func (b *basicBlogService) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer) (message string, status string, err error) {
	// TODO implement the business logic of Upload
	return message, status, err
}

// NewBasicBlogService returns a naive, stateless implementation of BlogService.
func NewBasicBlogService() BlogService {
	return &basicBlogService{}
}

// New returns a BlogService with all of the expected middleware wired in.
func New(middleware []Middleware) BlogService {
	var svc BlogService = NewBasicBlogService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
