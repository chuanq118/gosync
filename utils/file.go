package utils

import "os"

// EnsureDir ensures the directory exists.
func EnsureDir(dirPath string) error {
	// check is dir exists
	if _, err := os.Stat(dirPath); err == nil {
		// dir exists, do nothing
		return nil
	} else if !os.IsNotExist(err) {
		// some other error, return it
		return err
	}
	// if not exists, create it
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		// failed to create dir, return error
		return err
	}
	return nil
}
