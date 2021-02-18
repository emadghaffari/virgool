package file

import "os"

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
	return os.MkdirAll(path, os.ModePerm)
}