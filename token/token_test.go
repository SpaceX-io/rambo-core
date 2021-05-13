package token

import (
	"net/url"
	"testing"
	"time"
)

const (
	secret   = "242BB39A78E1251A"
	ttl      = 24 * time.Hour
	userId   = 111111111
	userName = "SpaceX-io"
)

func TestJwtSign(t *testing.T) {
	tokenStr, err := New(secret).JwtSign(userId, userName, ttl)
	if err != nil {
		t.Error("sign error", err)
		return
	}
	t.Log(tokenStr)
}

func TestJwtParse(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjExMTExMTExMSwiVXNlck5hbWUiOiJraXJpbnRhbmciLCJleHAiOjE2MjAzNjQ0NDcsImlhdCI6MTYyMDI3ODA0NywibmJmIjoxNjIwMjc4MDQ3fQ.Sw6qtW-8NNeKVhAS7j9HCeeb_udDKWCxNyRFXc9OyAM"
	user, err := New(secret).JwtParse(tokenStr)
	if err != nil {
		t.Error("token parse error", err)
		return
	}
	t.Log(user)
}

func TestUrlSign(t *testing.T) {
	urlPath := "/user/add"
	method := "post"
	params := url.Values{}
	params.Add("user_name", "SpaceX-io")
	params.Add("age", "25")
	params.Add("email", "xxxx@163.com")

	tokenStr, err := New(secret).UrlSign(urlPath, method, params)
	if err != nil {
		t.Error("url sign error", err)
		return
	}
	t.Log(tokenStr)
}

func BenchMarkJwtSignAndParse(b *testing.B) {
	b.ResetTimer()
	token := New(secret)
	for i := 0; i < b.N; i++ {
		tokenStr, _ := token.JwtSign(userId, userName, ttl)
		_, _ = token.JwtParse(tokenStr)
	}
}
