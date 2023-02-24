package helper_machine

import (
	"errors"
	"github.com/denisbrodbeck/machineid"
)

func GetId(name string) (string, error) {
	id, err := machineid.ProtectedID(name)
	if err != nil {
		return "", errors.New("get machine id faild")
	}
	return id, nil
}
