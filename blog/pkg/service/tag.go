package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/opentracing/opentracing-go"

	el "github.com/emadghaffari/virgool/blog/database/elastic"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	model "github.com/emadghaffari/virgool/blog/model"
)

// CreateTag
// create and store new tag into database and elasticsearch
func (b *basicBlogService) CreateTag(ctx context.Context, name string, token string) (message string, status string, err error) {

	// start create-tag tracer
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("create-tag")
	defer span.Finish()

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	tag := model.Tag{
		Name: name,
	}

	// try to store new tag to database
	if err := tx.Save(&tag).Error; err != nil {
		return fmt.Sprintf("error in store a new tag: %s", err.Error()), "ERROR", err
	}

	// try to commit
	tx.Commit()

	// try to store tags into elastic
	if _, err := el.Database.Store(ctx, "tag", "_doc", strconv.Itoa(int(tag.ID)), tag); err != nil {
		tx.Rollback()
		return "error in store tag into search engine", "ERROR", fmt.Errorf(fmt.Sprintf("error in store tag into search engine %s", err.Error()))
	}

	return "new tag stored successfully", "SUCCESS", nil
}

// GetTag
// for get a tag with filters from elastic
func (b *basicBlogService) GetTag(ctx context.Context, filter []*model.Query, token string) (tags []*model.Tag, message string, status string, err error) {

	// start get-tag tracer
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("get-tag")
	defer span.Finish()

	// build query with filters
	query, err := el.Database.BuildQuery(
		el.FilterQuery(filter),
	)

	if err != nil {
		return nil, "Failed To Create a Search Query", "500", fmt.Errorf("Failed To Create a Search Query: %s", err)
	}

	// search tags with query
	result, err := el.Database.Search(ctx, "tag", query)
	if err != nil {
		return nil, "Failed To Search", "500", fmt.Errorf("Failed To Search: %s", err)
	}

	// loop over hits and unmarshal the tags
	for _, hit := range result.Hits.Hits {
		var tag model.Tag
		if err := json.Unmarshal(hit.Source, &tag); err != nil {
			return tags, "Failed To Unmarshal Data", "500", fmt.Errorf("Failed To Unmarshal Data: %v", hit)
		}

		// append tags
		tags = append(tags, &tag)
	}

	return tags, "tags searched successfully", "SUCCESS", nil
}

// UpdateTag
// replace old tag name and to new name
func (b *basicBlogService) UpdateTag(ctx context.Context, oldName string, newName string, token string) (message string, status string, err error) {

	// start get-tag tracer
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("update-tag")

	defer span.Finish()

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	var tag model.Tag

	// get tag if you are owner of tag
	err = tx.Table("tags").
		Where("name = ? ", oldName).
		First(&tag).Error
	if err != nil {
		tx.Rollback()
		return "tag not found", "ERROR", err
	}

	tag.Name = newName

	// update tag with new vars
	err = tx.Table("tags").
		Where("name = ? ", oldName).
		Select("*").
		Updates(tag).Error
	if err != nil {
		tx.Rollback()
		return "we can not update your tag", "ERROR", err
	}

	// update elastic with new document
	_, err = el.Database.Update(ctx, "tag", "_doc", strconv.Itoa(int(tag.ID)), tag)
	if err != nil {
		tx.Rollback()
		return "we can not update your tag", "ERROR", err
	}

	tx.Commit()

	return "tag updated successfully", "SUCCESS", err
}
func (b *basicBlogService) DeleteTag(ctx context.Context, name string, token string) (message string, status string, err error) {

	del, err := el.Database.Delete(ctx, "tag", "_doc", name)
	if err != nil {
		return fmt.Sprintf("error in delete a tag: %s - deleted ID: %s", err, del.Id), "ERROR", err
	}

	return message, status, err
}
