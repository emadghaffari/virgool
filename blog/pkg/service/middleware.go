package service

import (
	"bytes"
	"context"

	model "github.com/emadghaffari/virgool/blog/model"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BlogService) BlogService

type loggingMiddleware struct {
	logger log.Logger
	next   BlogService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BlogService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BlogService) BlogService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreatePost(ctx context.Context, userID uint64, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "CreatePost", "userID", userID, "title", title, "slug", slug, "description", description, "text", text, "params", params, "medias", medias, "Tags", Tags, "Status", Status, "message", message, "status", status, "err", err)
	}()
	return l.next.CreatePost(ctx, userID, title, slug, description, text, params, medias, Tags, Status)
}
func (l loggingMiddleware) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []int64, Tags []int64, Status model.StatusPost) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "UpdatePost", "title", title, "slug", slug, "description", description, "text", text, "params", params, "medias", medias, "Tags", Tags, "Status", Status, "message", message, "status", status, "err", err)
	}()
	return l.next.UpdatePost(ctx, title, slug, description, text, params, medias, Tags, Status)
}
func (l loggingMiddleware) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query) (posts []model.Post, message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "GetPost", "must", must, "should", should, "not", not, "filter", filter, "posts", posts, "message", message, "status", status, "err", err)
	}()
	return l.next.GetPost(ctx, must, should, not, filter)
}
func (l loggingMiddleware) DeletePost(ctx context.Context, filter []*model.Query) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "DeletePost", "filter", filter, "message", message, "status", status, "err", err)
	}()
	return l.next.DeletePost(ctx, filter)
}
func (l loggingMiddleware) CreateTag(ctx context.Context, name string) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "CreateTag", "name", name, "message", message, "status", status, "err", err)
	}()
	return l.next.CreateTag(ctx, name)
}
func (l loggingMiddleware) GetTag(ctx context.Context, filter []*model.Query) (tags []*model.Tag, message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "GetTag", "filter", filter, "tags", tags, "message", message, "status", status, "err", err)
	}()
	return l.next.GetTag(ctx, filter)
}
func (l loggingMiddleware) UpdateTag(ctx context.Context, oldName string, newName string) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateTag", "oldName", oldName, "newName", newName, "message", message, "status", status, "err", err)
	}()
	return l.next.UpdateTag(ctx, oldName, newName)
}
func (l loggingMiddleware) DeleteTag(ctx context.Context, name string) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteTag", "name", name, "message", message, "status", status, "err", err)
	}()
	return l.next.DeleteTag(ctx, name)
}
func (l loggingMiddleware) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer) (message string, status string, err error) {
	defer func() {
		l.logger.Log("method", "Upload", "title", title, "description", description, "fileType", fileType, "file", file, "message", message, "status", status, "err", err)
	}()
	return l.next.Upload(ctx, title, description, fileType, file)
}
