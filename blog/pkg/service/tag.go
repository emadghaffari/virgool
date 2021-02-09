package service

import (
	"context"
	"fmt"

	el "github.com/emadghaffari/virgool/blog/database/elastic"
	model "github.com/emadghaffari/virgool/blog/model"
)

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

	del, err := el.Database.Delete(ctx, "tag", "_doc", name)
	if err != nil {
		return fmt.Sprintf("error in delete a post: %s - deleted ID: %s", err, del.Id), "ERROR", err
	}

	return message, status, err
}
