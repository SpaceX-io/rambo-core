package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (t *token) JwtSign(userId int64, userName string, expireDuration time.Duration) (tokenString string, err error) {
	claims := claims{
		userId,
		userName,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}

	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.secret))
	return
}

func (t *token) JwtParse(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
