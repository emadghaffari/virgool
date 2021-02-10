package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"

	el "github.com/emadghaffari/virgool/blog/database/elastic"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	model "github.com/emadghaffari/virgool/blog/model"
)

// CreatePost method for create new pos
func (b *basicBlogService) CreatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("create-post")
	defer span.Finish()

	user, ok := ctx.Value(model.User).(map[string]interface{})
	if !ok {
		return "user not found", "ERROR", fmt.Errorf(fmt.Sprintf("error user not found: %s", err.Error()))
	}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

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
		return "error in store new post", "ERROR", fmt.Errorf(fmt.Sprintf("error in store new post %s", err.Error()))
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
		return "error in append tags into post", "ERROR", fmt.Errorf(fmt.Sprintf("error in append tags into post %s", err.Error()))
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
		return "error in append params into post", "ERROR", fmt.Errorf(fmt.Sprintf("error in append params into post %s", err.Error()))
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
		return "error in append medias into post", "ERROR", fmt.Errorf(fmt.Sprintf("error in append medias into post %s", err.Error()))
	}

	tx.Commit()

	var ps model.Post
	if err := mysql.Database.GetDatabase().Table("posts").Preload("Media").Preload("Params").Preload("Tags").Where("slug=?", post.Slug).First(&ps).Error; err != nil {
		return "error in get new post stored by customer", "ERROR", fmt.Errorf(fmt.Sprintf("error in get new post stored by customer %s", err.Error()))
	}

	if _, err := el.Database.Store(ctx, "blog", ps); err != nil {
		tx.Rollback()
		return "error in store post into search engine", "ERROR", fmt.Errorf(fmt.Sprintf("error in store post into search engine %s", err.Error()))
	}

	return "post created successfully", "SUCCESS", nil
}

// update a post
func (b *basicBlogService) UpdatePost(ctx context.Context, title string, slug string, description string, text string, params []*model.Query, medias []uint64, Tags []uint64, Status model.StatusPost, token string) (message string, status string, err error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("update-post")
	defer span.Finish()

	user, ok := ctx.Value(model.User).(map[string]interface{})
	if !ok {
		return "user not found", "ERROR", err
	}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// blog model
	post := model.Post{
		PublishedAT: time.Now(),
	}

	err = tx.Table("posts").
		Where("user_id = ? AND slug = ?", uint64(user["id"].(float64)), slug).
		First(&post).Error
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

	// FIXME fix this section
	// _, err = el.Database.Update("blog", "_doc", strconv.Itoa(int(post.ID)), &elastic.Script{})
	// if err != nil {
	// 	tx.Rollback()
	// 	return "we can not update your post", "ERROR", err
	// }

	tx.Commit()

	message = "post updated successfully"
	status = "SUCCESS"

	return message, status, err
}

// get post
func (b *basicBlogService) GetPost(ctx context.Context, must []*model.Query, should []*model.Query, not []*model.Query, filter []*model.Query, token string) (posts []model.Post, message string, status string, err error) {
	query, err := el.Database.BuildQuery(must, should, not, filter)
	if err != nil {
		return nil, "Failed To Create a Search Query", "500", fmt.Errorf("Failed To Create a Search Query: %s", err)
	}
	result, err := el.Database.Search("blog", query)
	if err != nil {
		return nil, "Failed To Search", "500", fmt.Errorf("Failed To Search: %s", err)
	}

	for _, hit := range result.Hits.Hits {
		var post model.Post
		if err := json.Unmarshal(hit.Source, &post); err != nil {
			return posts, "Failed To Unmarshal Data", "500", fmt.Errorf("Failed To Unmarshal Data: %v", hit)
		}
		posts = append(posts, post)
	}

	return posts, " posts searched successfully", "SUCCESS", nil
}

// delete post
func (b *basicBlogService) DeletePost(ctx context.Context, filter []*model.Query, token string) (message string, status string, err error) {

	for _, id := range filter {
		id, ok := id.Value.(string)
		if ok {
			del, err := el.Database.Delete(ctx, "blog", "_doc", id)
			if err != nil {
				return fmt.Sprintf("error in delete a post: %s - deleted ID: %s", err, del.Id), "ERROR", err
			}
		}
	}

	return message, status, err
}
