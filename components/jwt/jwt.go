package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/langwan/langgo/core"
	"strings"
)

type Instance struct {
	Secret string `yaml:"secret"`
}

const name = "jwt"

var instance *Instance

func (i *Instance) Load() error {
	instance = i
	core.GetComponentConfiguration(name, i)
	return nil
}

func (i *Instance) GetName() string {
	return name
}

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

const (
	HS256 = "HS256"
)

var alg = HS256

func hs256(secret, data []byte) (ret string, err error) {
	hasher := hmac.New(sha256.New, secret)
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}
	r := hasher.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(r), nil
}

func Sign(payload interface{}) (ret string, err error) {
	h := header{
		Alg: alg,
		Typ: "JWT",
	}
	marshal, err := json.Marshal(h)
	if err != nil {
		return "", err
	}

	bh := base64.RawURLEncoding.EncodeToString(marshal)

	marshal, err = json.Marshal(payload)
	if err != nil {
		return "", err
	}

	bp := base64.RawURLEncoding.EncodeToString(marshal)

	s := fmt.Sprintf("%s.%s", bh, bp)

	ret, err = hs256([]byte(instance.Secret), []byte(s))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s.%s", bh, bp, ret), nil
}

func Verify(token string) (err error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("parts len error")
	}
	data := strings.Join(parts[0:2], ".")
	hasher := hmac.New(sha256.New, []byte(instance.Secret))
	_, err = hasher.Write([]byte(data))
	if err != nil {
		return err
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return err
	}
	if hmac.Equal(sig, hasher.Sum(nil)) {
		return nil
	}
	return errors.New("verify is invalid")
}

func GetPayload(token string) ([]byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("parts len error")
	}
	de, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	return de, nil
}
