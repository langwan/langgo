package tempfile

import (
	helper_gen "github.com/langwan/langgo/helpers/gen"
	"os"
	"path/filepath"
)

type TempFile struct {
	Base string `yaml:"base"`
}

func (tf TempFile) CreateFile(name string, data []byte, perm os.FileMode) error {
	p := filepath.Join(tf.Base, name)
	return os.WriteFile(p, data, perm)
}

func (tf TempFile) CreateTempFile(data []byte, perm os.FileMode) (string, error) {
	p := filepath.Join(tf.Base, helper_gen.UuidNoSeparator())
	return p, os.WriteFile(p, data, perm)
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
