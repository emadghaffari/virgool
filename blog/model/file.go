package model

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// FileStore struct
type FileStore struct {
	mutex  sync.RWMutex
	folder string
	files  map[string]*FileInfo
}

// FileInfo struct
type FileInfo struct {
	ID   string
	Path string
	Type string
}

// NewFileStore struct
func NewFileStore(folder string) *FileStore {
	return &FileStore{
		folder: folder,
		files:  make(map[string]*FileInfo),
	}
}

// Store meth for store a new file into blog service
func (f *FileStore) Store(folder string, fileType string, file bytes.Buffer) (*FileInfo, error) {
	// generate new UUID
	id, err := uuid.NewRandom()
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error in Store a new File ID: %s - Type: %s - Error: %v", id, fileType, err))
		return nil, err
	}

	// take default path for file
	path := fmt.Sprintf("%s/%s%s", folder, id, fileType)

	// create file
	oFile, err := os.Create(path)
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error in os.Create for a new File ID: %s - Type: %s - Error: %v", id, fileType, err))
		return nil, err
	}

	// write and upload file
	if _, err := file.WriteTo(oFile); err != nil {
		logrus.Warn(fmt.Sprintf("Error in file.WriteTo os.Created File ID: %s - Type: %s - Error: %v", id, fileType, err))
		return nil, err
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	// make new file result
	f.files[id.String()] = &FileInfo{
		ID:   id.String(),
		Type: fileType,
		Path: path,
	}

	return f.files[id.String()], nil
}
