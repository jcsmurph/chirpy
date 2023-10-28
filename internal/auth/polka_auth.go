package auth

import (
	"net/http"
	"strings"
)


func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}

    apiKey := strings.Split(authHeader, " ")

	return apiKey[1], nil
}
