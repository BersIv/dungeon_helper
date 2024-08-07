package util

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type MyJWTClaims struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

type TokenGetter interface {
	GetIdFromToken(r *http.Request) (int64, error)
	GetNickNameFromToken(r *http.Request) (string, error)
}

type JWTTokenGetter struct{}

func (JWTTokenGetter) GetIdFromToken(r *http.Request) (int64, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("authorization header is missing")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return 0, errors.New("invalid token format")
	}

	tokenString := splitToken[1]

	secret := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*MyJWTClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.Id, nil
}

func (JWTTokenGetter) GetNickNameFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return "", errors.New("invalid token format")
	}

	tokenString := splitToken[1]

	secret := os.Getenv("SECRET_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*MyJWTClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Nickname, nil
}
