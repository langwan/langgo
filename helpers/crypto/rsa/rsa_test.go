package helper_rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsa(t *testing.T) {
	bits := 2048
	spub, spri := CreateKeyX509PKCS1(bits)
	data := "0123456789"
	pub, err := PublicKeyFromX509PKCS1(spub)
	assert.NoError(t, err)

	pri, err := PrivateKeyFromX509PKCS1(spri)
	assert.NoError(t, err)

	encrypt, err := Encrypt(pub, []byte(data))
	assert.NoError(t, err)
	decrypt, err := Decrypt(pri, encrypt)
	assert.NoError(t, err)
	assert.Equal(t, data, string(decrypt))
}

func TestSign(t *testing.T) {
	bits := 2048
	spub, spri := CreateKeyX509PKCS1(bits)
	data := "0123456789"
	pub, err := PublicKeyFromX509PKCS1(spub)
	assert.NoError(t, err)
	t.Log(spub, spri)
	pri, err := PrivateKeyFromX509PKCS1(spri)
	assert.NoError(t, err)

	sign, err := Sign(pri, []byte(data))
	if err != nil {
		return
	}
	err = Verify(pub, sign, []byte(data))
	assert.NoError(t, err)
}

func TestPublicKeyToPKCS1(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	type args struct {
		pub *rsa.PublicKey
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				pub: publicKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, PublicKeyToPKCS1(tt.args.pub), "PublicKeyToPKCS1(%v)", tt.args.pub)
		})
	}
}

func TestPrivateKeyToPKCS1(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	type args struct {
		pri *rsa.PrivateKey
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				pri: privateKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, PrivateKeyToPKCS1(tt.args.pri), "PrivateKeyToPKCS1(%v)", tt.args.pri)
		})
	}
}
