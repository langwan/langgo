package helper_platform

import (
	"github.com/langwan/langgo/helpers/os"
	"os/user"
	"path/filepath"
)

func GetDefaultDocumentFolderPath() (document string) {
	u, _ := user.Current()
	p := filepath.Join(u.HomeDir, "Documents")
	if helper_os.FolderExists(p) == false {
		return ""
	}
	return p
}
