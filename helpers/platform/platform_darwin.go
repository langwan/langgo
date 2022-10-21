package helper_platform

import (
	"os/exec"
)

// OpenFileExplorer open file Explorer by path, support osx and windows
func OpenFileExplorer(path string) {
	exec.Command("open", "-R", path).Run()
}
