package upload

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"

	"github.com/emadghaffari/virgool/blog/model"
)

type FileStore interface {
	Store(fileType string, file bytes.Buffer) (string, error)
}

type DiskFileStore struct {
	mutex      sync.RWMutex
	fileFolder string
	File       *model.Media
}

func NewDiskFileStore(path string) *DiskFileStore {
	return &DiskFileStore{
		fileFolder: path,
	}
}

func (i *DiskFileStore) Store(fileType string, file bytes.Buffer) (string, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("Error in create new UUID for store File: %w", err)
	}

	filePath := fmt.Sprintf("%s/%s%s", i.fileFolder, fileID, fileType)

	ofile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("Error in create file for store File: %w", err)
	}

	if _, err := file.WriteTo(ofile); err != nil {
		return "", fmt.Errorf("Error in write to created file and store File: %w", err)
	}

	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.File = &model.Media{
		URL:  filePath,
		Type: fileType,
	}

	return fileID.String(), nil

}
