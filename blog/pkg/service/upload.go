package service

import (
	"bytes"
	"context"

	"github.com/emadghaffari/virgool/blog/utils/upload"
)

func (b *basicBlogService) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {

	disk := upload.NewDiskFileStore("public/upload")
	disk.File.Title = &title
	disk.File.Description = &description

	if _, err := disk.Store(fileType, file); err != nil {
		return "error in upload image", "ERROR", err
	}

	return "successfully file uploaded", "SUCCESS", nil
}
