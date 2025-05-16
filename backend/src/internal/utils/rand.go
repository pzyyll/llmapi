package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func RandBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	return bytes, err
}

func RandString(n int) (string, error) {
	bytes, err := RandBytes(n)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func RandBase64(n int) (string, error) {
	bytes, err := RandBytes(n)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func RandSaltWithBase64(n int) ([]byte, string, error) {
	bytes, err := RandBytes(n)
	if err != nil {
		return nil, "", err
	}
	return bytes, base64.RawURLEncoding.EncodeToString(bytes), nil
}
