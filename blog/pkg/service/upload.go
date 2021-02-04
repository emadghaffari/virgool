package service

import (
	"bytes"
	"context"
)

func (b *basicBlogService) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {
	// TODO implement the business logic of Upload
	return message, status, err
}
