package tempfile

import (
	"github.com/langwan/langgo/helpers/gen"
	"os"
	"path/filepath"
)

type TempFile struct {
	Base string `yaml:"base"`
}

func (tf TempFile) CreateFile(data []byte, perm os.FileMode) (string, error) {
	filename := helper_gen.UuidNoSeparator()
	p := filepath.Join(tf.Base, filename)
	return filename, os.WriteFile(p, data, perm)
}

func (tf TempFile) ReadFile(name string, remove bool) ([]byte, error) {
	p := filepath.Join(tf.Base, name)
	data, err := os.ReadFile(p)
	if remove {
		defer os.Remove(p)
	}
	return data, err
}

func (tf TempFile) RemoveFile(name string) error {
	p := filepath.Join(tf.Base, name)
	return os.Remove(p)
}
