package file

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Exists checks if a file exists
// try using it to prevent further errors.
func Exists(filename string,isDirectory bool) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }

	if isDirectory{
		return info.IsDir()

	}
    return !info.IsDir()

}

// CreateDir create nested directories
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
    if err != nil {
        logrus.Warn(fmt.Sprintf("error in create new directory: %s",err))
    }
    
    return err
}