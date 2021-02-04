package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/emadghaffari/virgool/blog/database/mysql"
	model "github.com/emadghaffari/virgool/blog/model"
	"github.com/emadghaffari/virgool/blog/utils/str"
)

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

	// make slice of tags len
	// change format of tags
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

	// make slice of params len
	// change format of params
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

	// make slice of medias len
	// change format of medias
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
