package helper_os

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
)

func GetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

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

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if info == nil {
		return false
	}
	return true
}

func MoveFile(source string, dest string) error {
	err := CopyFile(source, dest)
	if err != nil {
		return err
	}
	err = os.Remove(source)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}

func CopyFile(source string, dest string) error {
	inputFile, err := os.Open(source)
	defer inputFile.Close()
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(dest)
	defer outputFile.Close()
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	return nil
}
