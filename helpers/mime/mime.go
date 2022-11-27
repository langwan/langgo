package helper_mime

import (
	"errors"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"

	"os"
	"path/filepath"

	"strings"
)

// GetMimeType get file mime and file type
func GetMimeType(filename string) (mimeType string, fileType string, err error) {

	osOpen, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer osOpen.Close()
	return getFileContentType(filename, osOpen)

}

// IsRaw the file is a raw image
func IsRaw(buf []byte) bool {
	kind, _ := Raw(buf)
	return kind != types.Unknown
}

func getFileContentType(filename string, osFile *os.File) (mimeType string, fileType string, err error) {

	buffer := make([]byte, 512)
	_, err = osFile.Read(buffer)

	if err != nil {
		return "", "", errors.New("GetFileContentType error:" + err.Error())
	}

	contentType, err := filetype.Get(buffer)
	if err != nil {
		return "", "", errors.New("GetFileContentType error:" + err.Error())
	}

	fileType = "file"

	if contentType == filetype.Unknown {
		ext := filepath.Ext(filename)
		if strings.ToLower(ext) == ".md" {
			fileType = "md"
		}
	} else {

		if filetype.IsVideo(buffer) {
			fileType = "video"
		} else if IsRaw(buffer) {
			fileType = "raw"
		} else if filetype.IsImage(buffer) {
			fileType = "image"
		} else {
			fileType = "file"
		}

	}
	return contentType.MIME.Value, fileType, nil
}

func Raw(buf []byte) (types.Type, error) {
	return doMatchMap(buf, matchersRaw)
}

func doMatchMap(buf []byte, machers matchers.Map) (types.Type, error) {
	kind := filetype.MatchMap(buf, machers)
	if kind != types.Unknown {
		return kind, nil
	}
	return kind, filetype.ErrUnknownBuffer
}

var (
	TypeCR2  = types.NewType("cr2", "image/x-canon-cr2")
	TypeTiff = types.NewType("tif", "image/tiff")
)

var matchersRaw = matchers.Map{
	TypeCR2:  matchers.CR2,
	TypeTiff: matchers.Tiff,
}
