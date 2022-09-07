package uuid

import (
	"github.com/lithammer/shortuuid/v3"
	gouuid "github.com/satori/go.uuid"
	"strings"
)

func String() string {
	return gouuid.NewV4().String()
}

func Short() string {
	return shortuuid.New()
}

func String2() string {
	return strings.Replace(gouuid.NewV4().String(), "-", "", -1)
}
