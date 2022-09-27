package aes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAes(t *testing.T) {
	s := "0123456789"
	key := "01234567890123456789012345678901"
	encrypt, err := Encrypt([]byte(key), []byte(s))

	if err != nil {
		assert.NoError(t, err)
	}
	decrypt, err := Decrypt([]byte(key), encrypt)
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, s, string(decrypt))

}
