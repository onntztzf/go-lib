package jwt

import (
	"time"

	"github.com/gh-zhangpeng/lib/e"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

const (
	GenerateTokenFailErrorCode = 101
	ParseTokenFailErrorCode    = 102
	InvalidTokenErrorCode      = 103
	ExpiredTokenErrorCode      = 104
)

var GenerateTokenFailError = e.NewError(GenerateTokenFailErrorCode, "generate token fail")
var ParseTokenFailError = e.NewError(ParseTokenFailErrorCode, "parse token fail")
var InvalidTokenError = e.NewError(InvalidTokenErrorCode, "invalid token")
var ExpiredTokenError = e.NewError(ExpiredTokenErrorCode, "expired token")

var key = []byte("box-key")

type claims struct {
	UserID uint64 `json:"userID"`
	jwt.StandardClaims
}

func ParseToken(token string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, ParseTokenFailError
	}
	claims, ok := tokenClaims.Claims.(*claims)
	if !ok {
		return nil, ParseTokenFailError
	}
	if tokenClaims.Valid {
		return claims, nil
	} else {
		e, ok := err.(*jwt.ValidationError)
		if ok {
			if e.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return claims, errors.Wrapf(ExpiredTokenError, "token has not taken effect or has expired, claims: %+v, token: %s", claims, token)
			} else if e.Errors&(jwt.ValidationErrorAudience) != 0 {
				return claims, errors.Wrapf(InvalidTokenError, "token has no permission to use the current application, claims: %+v, token: %s", claims, token)
			} else if e.Errors&(jwt.ValidationErrorIssuer) != 0 {
				return claims, errors.Wrapf(InvalidTokenError, "token issuer is incorrect, claims: %+v, token: %s", claims, token)
			} else {
				return claims, errors.Wrapf(InvalidTokenError, "parse token fail, err: %s, claims: %+v, token: %s", err.Error(), claims, token)
			}
		} else {
			return claims, errors.Wrapf(ParseTokenFailError, "parse token fail, err: %s, token: %s", err.Error(), token)
		}
	}
}

func GenerateToken(userID uint64, expiresAt int64) (string, error) {
	claims := claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(key)
	if err != nil {
		return "", errors.Wrapf(GenerateTokenFailError, "signed string fail, err: %s, claims: %+v", err.Error(), claims)
	}
	return signedString, nil
}
