package io

import (
	"errors"
	"os"
)

func CreateFolder(p string, ignoreExists bool) error {
	if FolderExists(p) == true && ignoreExists == false {
		err := errors.New("folder exists")
		return err
	}
	if FolderExists(p) == false {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func FolderExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if info == nil {
		return false
	}
	return info.IsDir()
}
