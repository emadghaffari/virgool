package service

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/emadghaffari/virgool/blog/database/mysql"
	"github.com/emadghaffari/virgool/blog/model"
	"github.com/emadghaffari/virgool/blog/utils/str"
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

// CreatePost method for create new pos
func (b *basicBlogService) CreatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("create-post")
	defer span.Finish()

	user := ctx.Value(model.User).(map[string]interface{})

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// remove symbols
	slug, err = str.RemoveSymbols(slug)
	if err != nil {
		return "Extra Symbols into slug", "ERROR", err
	}

	// blog model
	post := model.Post{
		UserID:      uint64(user["id"].(float64)),
		Title:       title,
		Slug:        strings.ReplaceAll(slug, " ", "-"),
		Description: description,
		Text:        text,
		Status:      Status,
		PublishedAT: time.Now(),
	}

	// try to store post with model
	if gm := tx.Save(&post); gm.Error != nil {
		tx.Rollback()
		return message, "ERROR", fmt.Errorf(err.Error())
	}

	t := make([]*model.Tag, len(Tags))
	for kt, vt := range Tags {
		t[kt] = &model.Tag{
			ID: vt,
		}
	}
	if err := tx.Model(&post).Association("Tags").Append(t); err != nil {
		tx.Rollback()
		return message, "ERROR", fmt.Errorf(err.Error())
	}

	p := make([]*model.Param, len(params))
	for kp, vp := range params {
		p[kp] = &model.Param{
			Query: model.Query{Name: vp.Name, Value: vp.Value},
		}
	}
	if err := tx.Model(&post).Association("Params").Append(p); err != nil {
		tx.Rollback()
		return message, "ERROR", fmt.Errorf(err.Error())
	}

	m := make([]*model.Media, len(medias))
	for km, vm := range medias {
		m[km] = &model.Media{
			ID: vm,
		}
	}
	if err := tx.Model(&post).Association("Media").Append(m); err != nil {
		tx.Rollback()
		return message, "ERROR", fmt.Errorf(err.Error())
	}

	tx.Commit()

	return "post created successfully", "SUCCESS", err
}
func (b *basicBlogService) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("update-post")
	defer span.Finish()

	user := ctx.Value(model.User).(map[string]interface{})

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// blog model
	post := model.Post{
		PublishedAT: time.Now(),
	}

	err = tx.Table("posts").Preload("Params").Preload("Media").Preload("Tags").Where("user_id = ? AND slug = ?", uint64(user["id"].(float64)), slug).First(&post).Error
	if err != nil {
		tx.Rollback()
		return "post not found", "ERROR", err
	}

	t := make([]*model.Tag, len(Tags))
	for kt, vt := range Tags {
		t[kt] = &model.Tag{
			ID: vt,
		}
	}

	p := make([]*model.Param, len(params))
	for kp, vp := range params {
		p[kp] = &model.Param{
			Query: model.Query{Name: vp.Name, Value: vp.Value},
		}
	}

	m := make([]*model.Media, len(medias))
	for km, vm := range medias {
		m[km] = &model.Media{
			ID: vm,
		}
	}

	post.Tags = t
	post.Params = p
	post.Media = m
	post.Title = title
	post.Slug = slug
	post.Description = description
	post.Status = Status
	post.Text = text

	// FIXME
	// change update model
	err = tx.Table("posts").
		Preload("Params").
		Preload("Media").
		Preload("Tags").
		Where("user_id = ? AND slug = ?", uint64(user["id"].(float64)), slug).
		Select("*").
		Updates(post).Error
	if err != nil {
		tx.Rollback()
		return "we can not update your post", "ERROR", err
	}

	tx.Commit()

	message = "post updated successfully"
	status = "SUCCESS"

	return message, status, err
}
func (b *basicBlogService) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query, token string) (posts []model.Post, message string, status string, err error) {
	// TODO implement the business logic of GetPost
	return posts, message, status, err
}
func (b *basicBlogService) DeletePost(ctx context.Context, filter []*model.Query, token string) (message string, status string, err error) {
	// TODO implement the business logic of DeletePost
	return message, status, err
}
func (b *basicBlogService) CreateTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	// TODO implement the business logic of CreateTag
	return message, status, err
}
func (b *basicBlogService) GetTag(ctx context.Context, filter []*model.Query, token string) (tags []*model.Tag, message string, status string, err error) {
	// TODO implement the business logic of GetTag
	return tags, message, status, err
}
func (b *basicBlogService) UpdateTag(ctx context.Context, oldName string, newName string, token string) (message string, status string, err error) {
	// TODO implement the business logic of UpdateTag
	return message, status, err
}
func (b *basicBlogService) DeleteTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	// TODO implement the business logic of DeleteTag
	return message, status, err
}
func (b *basicBlogService) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {
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
