package jwt

import (
	"time"

	"github.com/2hangpeng/lib/e"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const (
	GenerateTokenFailErrorCode = 101
	ParseTokenFailErrorCode    = 102
	InvalidTokenErrorCode      = 103
)

var GenerateTokenFailError = e.NewError(GenerateTokenFailErrorCode, "generate token fail")
var ParseTokenFailError = e.NewError(ParseTokenFailErrorCode, "parse token fail")
var InvalidTokenError = e.NewError(InvalidTokenErrorCode, "invalid token")

var key = []byte("key")

type Claims struct {
	UserID uint64 `json:"userID"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, errors.Wrap(ParseTokenFailError, err.Error())
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.Wrap(InvalidTokenError, "invalid token")
	}
	return claims, nil
}

func GenerateToken(userID uint64, expiresAt int64) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", errors.Wrap(GenerateTokenFailError, err.Error())
	}
	return tokenString, nil
}
