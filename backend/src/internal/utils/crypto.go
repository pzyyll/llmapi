package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/argon2"
)

func EncodeArgon2(
	secret []byte,
	salt []byte,
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
) string {
	// Hash the secret using Argon2
	hash := argon2.IDKey(
		secret,
		salt,
		time,
		memory,
		threads,
		keyLen,
	)
	// Encode the hash to base64
	encodedHash := base64.RawURLEncoding.EncodeToString(hash)
	return encodedHash
}

func DecodeBase64(
	encoded string,
) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(encoded)
}

func EncodeSha256(
	secret []byte,
) string {
	// Hash the secret using SHA-256
	sha256 := sha256.New()
	sha256.Write(secret)
	hash := sha256.Sum(nil)
	encodedHash := base64.RawURLEncoding.EncodeToString(hash)
	return encodedHash
}

func KeyBrief(apiKey string) string {
	const minLengthForNoPrefixAbbreviation = 7 // "..." + 4 个字符
	const lastCharsToKeep = 4

	parts := strings.SplitN(apiKey, "-", 2)
	if len(parts) == 2 && len(parts[0]) > 0 {
		prefix := parts[0] + "-"
		secretPart := parts[1]

		if len(secretPart) > lastCharsToKeep {
			return prefix + "..." + secretPart[len(secretPart)-lastCharsToKeep:]
		}
		return prefix + "..." + secretPart
	}

	if len(apiKey) > minLengthForNoPrefixAbbreviation {
		return "..." + apiKey[len(apiKey)-lastCharsToKeep:]
	}

	return apiKey
}
