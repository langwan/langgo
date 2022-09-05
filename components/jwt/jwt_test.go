package jwt

import "testing"

func TestSign(t *testing.T) {
	instance = &Instance{Secret: "123456"}
	payload := struct {
		Name string
	}{Name: "langwan"}
	sign, err := Sign(payload)
	if err != nil {
		t.Error(err)
		return
	}
	err = Verify(sign)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sign)
}
