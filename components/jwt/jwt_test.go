package jwt

import (
	"runtime"
	"testing"
)

func TestGetFilename(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)
}
