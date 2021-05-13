package token

import (
	"github.com/dgrijalva/jwt-go"
	"net/url"
	"time"
)

var _ Token = (*token)(nil)

type Token interface {
	i()

	JwtSign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error)
	JwtParse(tokenString string) (*claims, error)

	UrlSign(path, method string, params url.Values) (tokenString string, err error)
}

type token struct {
	secret string
}

type claims struct {
	UserId   int64
	UserName string
	jwt.StandardClaims
}

func New(secret string) Token {
	return &token{
		secret: secret,
	}
}

func (t *token) i() {}
