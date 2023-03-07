package lib

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomBase64String(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	fmt.Println(base64.URLEncoding.EncodeToString(bytes))

	return base64.URLEncoding.EncodeToString(bytes), nil
}
