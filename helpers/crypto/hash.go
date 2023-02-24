package helper_hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(orig string, salt string) string {
	hn := sha1.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}

func Md5(orig string, salt string) string {
	hn := md5.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}
