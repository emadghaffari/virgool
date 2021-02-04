package service

import (
	"bytes"
	"context"

	"github.com/emadghaffari/virgool/blog/model"
)

// BlogService describes the service.
type BlogService interface {
	// CRUD Posts
	CreatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error)
	UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error)
	GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query, token string) (posts []model.Post, message string, status string, err error)
	DeletePost(ctx context.Context, filter []*model.Query, token string) (message string, status string, err error)

	// CRUD tags
	CreateTag(ctx context.Context, name string, token string) (message string, status string, err error)
	GetTag(ctx context.Context, filter []*model.Query, token string) (tags []*model.Tag, message string, status string, err error)
	UpdateTag(ctx context.Context, oldName, newName string, token string) (message string, status string, err error)
	DeleteTag(ctx context.Context, name string, token string) (message string, status string, err error)

	// CRUD media
	// multipart.File
	Upload(ctx context.Context, title, description, fileType string, file bytes.Buffer, token string) (message string, status string, err error)
}

type basicBlogService struct{}

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
