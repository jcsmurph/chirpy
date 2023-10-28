package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrNoAuthHeaderIncluded = errors.New("not auth header included in request")

type JwtClaim struct {
	jwt.RegisteredClaims
}

func MakeJWT(userID int, tokenSecret string) (string, string, error) {
	// Access Token
	mySigningKey := []byte(tokenSecret)
	access_claims := JwtClaim{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(1 * time.Hour))),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "chirpy-access",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
	accessSS, err := accessToken.SignedString(mySigningKey)

	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := JwtClaim{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(1440 * time.Hour))),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "chirpy-refresh",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSS, err := refreshToken.SignedString(mySigningKey)

	if err != nil {
		return "", "", err
	}

	return accessSS, refreshSS, nil
}

func ValidateJWT(tokenString, tokenSecret string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return "", err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return userIDString, nil
}

// GetBearerToken -
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func ValidateAccessToken(tokenString, tokenSecret string) error {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)

	if err != nil {
		return err
	}
	tokenIssuer, err := token.Claims.GetIssuer()
	if tokenIssuer != "chirpy-access" {
		return errors.New("Not an access token")
	}

	return nil
}

