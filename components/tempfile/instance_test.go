package tempfile

import (
	"github.com/langwan/langgo"
	"testing"
)

func TestTempFile_Create(t *testing.T) {
	langgo.Run(&Instance{
		Base: "./temp",
	})
	Get().CreateFile([]byte{1, 2, 3}, 0666)
	Get().ReadFile("t1.txt", true)
}
