package auth

import (
	"fmt"
	"strings"
)

func GetAuthorizationToken(authorization string) (string, string, error) {
	if authorization == "" {
		return "", "", fmt.Errorf("authorization header is empty")
	}

	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid authorization format")
	}
	return parts[0], parts[1], nil
}

func GetAuthorizationTokenFromHeader(header string, proto string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("authorization header is empty")
	}

	checkProto, token, err := GetAuthorizationToken(header)
	if err != nil {
		return "", err
	}

	if token == "" {
		return "", fmt.Errorf("authorization token is empty")
	}

	if !strings.EqualFold(checkProto, proto) {
		return "", fmt.Errorf("authorization type not supported")
	}

	return token, nil
}
