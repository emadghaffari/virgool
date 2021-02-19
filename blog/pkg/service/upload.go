package service

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/opentracing/opentracing-go"

	el "github.com/emadghaffari/virgool/blog/database/elastic"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	"github.com/emadghaffari/virgool/blog/utils/upload"
)

func (b *basicBlogService) Upload(ctx context.Context, title string, description string, fileType string, file bytes.Buffer, token string) (message string, status string, err error) {

	// start get-tag tracer
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("upload-media")

	defer span.Finish()

	disk := upload.NewDiskFileStore(time.Now().UTC().String())
	disk.File.Title = &title
	disk.File.Description = &description

	if _, err := disk.Store(fileType, file); err != nil {
		return "error in upload image", "ERROR", err
	}

	// begins a transaction
	tx := mysql.Database.GetDatabase().Begin()

	// try to store new tag to database
	if err := tx.Save(&disk.File).Error; err != nil {
		return fmt.Sprintf("error in store a new media: %s", err.Error()), "ERROR", err
	}

	// try to commit
	tx.Commit()

	// try to store tags into elastic
	if _, err := el.Database.Store(ctx, "media", "_doc", strconv.Itoa(int(disk.File.ID)), disk.File); err != nil {
		tx.Rollback()
		return "error in store media into search engine", "ERROR", fmt.Errorf(fmt.Sprintf("error in store media into search engine %s", err.Error()))
	}

	return "successfully file uploaded", "SUCCESS", nil
}
