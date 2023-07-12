package jwt

import (
	"fmt"
	"time"

	"github.com/2hangpeng/go-lib/e"
	"github.com/golang-jwt/jwt"
)

const key = "key"

type Claims struct {
	UserID uint64 `json:"userID"`
	jwt.StandardClaims
}

// 定义错误类型和错误代码
var (
	ErrGenerateTokenFail = e.SystemError.ReplaceMsg("generate token fail")
	ErrParseTokenFail    = e.SystemError.ReplaceMsg("parse token fail")
	ErrInvalidToken      = e.SystemError.ReplaceMsg("invalid token")
)

// ValidateToken 解析并验证 JWT Token
func ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, ErrParseTokenFail
	}
	if tokenClaims.Valid {
		return claims, nil
	}
	validationErr, ok := err.(*jwt.ValidationError)
	if ok {
		return claims, ErrInvalidToken.ReplaceMsg(fmt.Sprintf("%s, %s", ErrInvalidToken.Msg, validationErr.Error()))
	}
	return claims, ErrInvalidToken
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID uint64, expiresAt int64) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", ErrGenerateTokenFail
	}
	return signedString, nil
}
