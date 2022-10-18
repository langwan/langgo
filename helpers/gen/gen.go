package helper_gen

import (
	"crypto/rand"
	"github.com/lithammer/shortuuid/v3"
	gouuid "github.com/satori/go.uuid"
	"math/big"
	"strings"
)

func Uuid() string {
	return gouuid.NewV4().String()
}

func UuidShort() string {
	return shortuuid.New()
}

func UuidNoSeparator() string {
	return strings.Replace(Uuid(), "-", "", -1)
}

func RandInt(min, max int64) int64 {
	// calculate the max we will be using
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}

const (
	LettersLetter          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersUpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersNumber          = "0123456789"
	LettersNumberNoZero    = "23456789"
	LettersSymbol          = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

func RandString(n int, letters ...string) (string, error) {

	lettersDefaultValue := LettersLetter + LettersNumber + LettersSymbol

	if len(letters) > 0 {
		lettersDefaultValue = ""
		for _, letter := range letters {
			lettersDefaultValue = lettersDefaultValue + letter
		}
	}

	bytes := make([]byte, n)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = lettersDefaultValue[b%byte(len(lettersDefaultValue))]
	}

	return string(bytes), nil
}
