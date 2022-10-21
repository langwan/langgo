package helper_platform

import (
	"fmt"
	"os/exec"
)

// OpenFileExplorer open file Explorer by path, support osx and windows
func OpenFileExplorer(path string) {
	exec.Command("cmd", "/C", fmt.Sprintf("start %s", path)).Run()
}
