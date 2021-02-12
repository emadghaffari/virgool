package service

import (
	"bytes"
	"context"
	"strings"

	log "github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"

	model "github.com/emadghaffari/virgool/blog/model"
	"github.com/emadghaffari/virgool/blog/utils/str"
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

func (l loggingMiddleware) CreatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "CreatePost", "title", title, "slug", slug, "description", description, "text", text, "params", params, "medias", medias, "Tags", Tags, "Status", Status, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", "ERROR", err
	}

	if slug == ""{
		slug = title
	}

	slug,err = str.RemoveSymbols(slug)
	if err != nil{
		return "invalid slug", "ERROR", err
	}
	slug = strings.ReplaceAll(slug," ","-")

	return l.next.CreatePost(context.WithValue(ctx, model.User, user), title, slug, description, text, params, medias, Tags, Status, token)
}
func (l loggingMiddleware) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "UpdatePost", "title", title, "slug", slug, "description", description, "text", text, "params", params, "medias", medias, "Tags", Tags, "Status", Status, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", "ERROR", err
	}

	if slug == ""{
		slug = title
	}

	slug,err = str.RemoveSymbols(slug)
	if err != nil{
		return "invalid slug", "ERROR", err
	}
	slug = strings.ReplaceAll(slug," ","-")

	return l.next.UpdatePost(context.WithValue(ctx, model.User, user), title, slug, description, text, params, medias, Tags, Status, token)
}
func (l loggingMiddleware) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query, token string) (posts []model.Post, message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "GetPost", "must", must, "should", should, "not", not, "filter", filter, "token", token, "posts", posts, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()
	return l.next.GetPost(ctx, must, should, not, filter, token)
}
func (l loggingMiddleware) DeletePost(ctx context.Context, filter []*model.Query, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "DeletePost", "filter", filter, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()
	return l.next.DeletePost(ctx, filter, token)
}
func (l loggingMiddleware) CreateTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "CreateTag", "name", name, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", "ERROR", err
	}

	name,err = str.RemoveSymbols(name)
	if err != nil{
		return "invalid name", "ERROR", err
	}

	return l.next.CreateTag(ctx, name, token)
}
func (l loggingMiddleware) GetTag(ctx context.Context, filter []*model.Query, token string) (tags []*model.Tag, message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "GetTag", "filter", filter, "token", token, "tags", tags, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()
	return l.next.GetTag(ctx, filter, token)
}
func (l loggingMiddleware) UpdateTag(ctx context.Context, oldName string, newName string, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "UpdateTag", "oldName", oldName, "newName", newName, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", "ERROR", err
	}

	oldName,err = str.RemoveSymbols(oldName)
	if err != nil{
		return "invalid old name", "ERROR", err
	}

	newName,err = str.RemoveSymbols(newName)
	if err != nil{
		return "invalid new name for tag", "ERROR", err
	}

	return l.next.UpdateTag(ctx, oldName, newName, token)
}
func (l loggingMiddleware) DeleteTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "DeleteTag", "name", name, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()

	var user interface{}
	if err := model.JWT.Get(ctx, token, &user); err != nil {
		err := l.logger.Log("user", "not", "found")
		if err != nil {
			logrus.Warn(err.Error())
		}
		return "user not found", "ERROR", err
	}

	name,err = str.RemoveSymbols(name)
	if err != nil{
		return "invalid name", "ERROR", err
	}

	return l.next.DeleteTag(ctx, name, token)
}
func (l loggingMiddleware) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {
	defer func() {
		err := l.logger.Log("method", "Upload", "title", title, "description", description, "fileType", fileType, "file", file, "token", token, "message", message, "status", status, "err", err)
		if err != nil {
			logrus.Warn(err.Error())
		}
	}()
	return l.next.Upload(ctx, title, description, fileType, file, token)
}
