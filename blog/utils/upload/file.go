package upload

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/model"
	f "github.com/emadghaffari/virgool/blog/utils/file"
)

type FileStore interface {
	Store(fileType string, file bytes.Buffer) (string, error)
}

type DiskFileStore struct {
	mutex      sync.RWMutex
	fileFolder string
	File       *model.Media
}

// NewDiskFileStore is path to file should store
// example: /path
func NewDiskFileStore(path string) *DiskFileStore {
	if string(path[0]) != "/" {
		path = "/" + path 
	}

	path = conf.GlobalConfigs.General.Upload+path

	// check base directory exists for upload file
	if !f.Exists(path,true){
		err := f.CreateDir(path)
		if err != nil {
			logrus.Warn(err)
		}
	}

	return &DiskFileStore{
		fileFolder: path,
	}
}

// Store method, store a file into path you want
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
