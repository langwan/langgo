package helper_crypto_file

import (
	helper_aes "github.com/langwan/langgo/helpers/crypto/aes"
	"io"
	"os"
)

const (
	DefaultBufferSize = 204800
	DefaultBlockSize  = 204832
)

type File struct {
	Secret          []byte
	IsEncryptedFile bool
	Handle          *os.File
}

func (f *File) Open(name string) error {
	var err error
	f.Handle, err = os.Open(name)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Close() error {
	return f.Handle.Close()
}

func (f *File) Create(name string) error {
	var err error
	f.Handle, err = os.Create(name)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Seek(offset int64, whence int) (ret int64, err error) {
	return f.Handle.Seek(offset, whence)
}

func (f *File) WriteAt(data []byte, off int64) (n int, err error) {
	encrypt, err := helper_aes.Encrypt(f.Secret, data)
	if err != nil {
		return n, err
	}
	return f.Handle.WriteAt(encrypt, off)
}

func (f *File) Write(b []byte) (n int, err error) {
	encrypt, err := helper_aes.Encrypt(f.Secret, b)
	if err != nil {
		return n, err
	}
	return f.Handle.Write(encrypt)
}

func (f *File) ReadAt(b []byte, off int64) (de []byte, n int, err error) {
	n, err = f.Handle.ReadAt(b, off)
	if err != nil && err != io.EOF {
		return de, n, err
	}

	de, err = helper_aes.Decrypt(f.Secret, b[:n])
	if err != nil {
		return de, n, err
	}
	return de, n, nil
}

func (f *File) Read(b []byte) (de []byte, n int, err error) {
	n, err = f.Handle.Read(b)
	if err != nil {
		return de, n, err
	}
	de, err = helper_aes.Decrypt(f.Secret, b)
	if err != nil {
		return de, n, err
	}
	return de, n, nil
}

func WriteFile(secret []byte, name string, data []byte, perm os.FileMode) error {
	encrypt, err := helper_aes.Encrypt(secret, data)
	if err != nil {
		return err
	}
	return os.WriteFile(name, encrypt, perm)
}

func ReadFile(secret []byte, name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	decrypt, err := helper_aes.Decrypt(secret, data)
	if err != nil {
		return nil, err
	}
	return decrypt, nil
}
