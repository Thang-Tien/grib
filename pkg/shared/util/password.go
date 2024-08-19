package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

func HashMD5(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("chuỗi nhập vào không được rỗng")
	}
	hash := md5.New()
	io.WriteString(hash, input)
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
