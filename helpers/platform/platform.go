package platform

import (
	"github.com/langwan/langgo/helpers/io"
	"os/user"
	"path"
)

func GetDefaultDocumentFolderPath() (document string) {
	u, _ := user.Current()
	p := path.Join(u.HomeDir, "Documents")
	if io.FolderExists(p) == false {
		return ""
	}
	return p
}
