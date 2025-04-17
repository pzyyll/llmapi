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
