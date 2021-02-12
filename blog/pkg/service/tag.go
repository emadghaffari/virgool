package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/opentracing/opentracing-go"

	el "github.com/emadghaffari/virgool/blog/database/elastic"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	model "github.com/emadghaffari/virgool/blog/model"
)

func (b *basicBlogService) CreateTag(ctx context.Context, name string, token string) (message string, status string, err error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("create-tag")
	defer span.Finish()

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	tag := model.Tag{
		Name: name,
	}

	if err:= tx.Save(&tag).Error; err != nil{
		return fmt.Sprintf("error in store a new tag: %s",err.Error()),"ERROR",err
	}

	tx.Commit()

	if _, err := el.Database.Store(ctx, "tag", "_doc", strconv.Itoa(int(tag.ID)), tag); err != nil {
		tx.Rollback()
		return "error in store tag into search engine", "ERROR", fmt.Errorf(fmt.Sprintf("error in store tag into search engine %s", err.Error()))
	}


	return "new tag stored successfully", "SUCCESS", nil
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

	del, err := el.Database.Delete(ctx, "tag", "_doc", name)
	if err != nil {
		return fmt.Sprintf("error in delete a tag: %s - deleted ID: %s", err, del.Id), "ERROR", err
	}

	return message, status, err
}
