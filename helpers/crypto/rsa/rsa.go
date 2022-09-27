package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func Encrypt(key *rsa.PublicKey, src []byte) (data []byte, err error) {
	h := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(h, rand.Reader, key, src, nil)

	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func Decrypt(key *rsa.PrivateKey, src []byte) (data []byte, err error) {
	h := sha256.New()
	oaep, err := rsa.DecryptOAEP(h, rand.Reader, key, src, nil)
	if err != nil {
		return nil, err
	}
	return oaep, nil
}

func Sign(key *rsa.PrivateKey, src []byte) (sign []byte, err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.SignPSS(rand.Reader, key, h, sum, nil)
}

func Verify(key *rsa.PublicKey, sign, src []byte) (err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.VerifyPSS(key, h, sum, sign, nil)
}

func CreateKeyX509PKCS1(bits int) (pub string, pri string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	publicKey := &privateKey.PublicKey
	bytePri := x509.MarshalPKCS1PrivateKey(privateKey)
	pri = base64.StdEncoding.EncodeToString(bytePri)
	bytePub := x509.MarshalPKCS1PublicKey(publicKey)
	pub = base64.StdEncoding.EncodeToString(bytePub)
	return pub, pri
}

func PrivateKeyFromX509PKCS1(pri string) (*rsa.PrivateKey, error) {
	data, err := base64.StdEncoding.DecodeString(pri)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(data)
}

func PublicKeyFromX509PKCS1(pub string) (*rsa.PublicKey, error) {
	data, err := base64.StdEncoding.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(data)
}
