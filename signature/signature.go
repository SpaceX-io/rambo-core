package signature

import (
	"net/http"
	"net/url"
	"time"
)

var _ Signature = (*signature)(nil)

const delimiter = "|"

var methods = map[string]bool{
	http.MethodGet:     true,
	http.MethodPost:    true,
	http.MethodDelete:  true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodConnect: true,
	http.MethodPatch:   true,
	http.MethodTrace:   true,
	http.MethodPut:     true,
}

type Signature interface {
	i()

	Generate(path, method string, params url.Values) (authorization, date string, err error)
	Verify(authorization, date, path, method string, params url.Values) (ok bool, err error)
}

type signature struct {
	key    string
	secret string
	ttl    time.Duration
}

func New(key, secret string, ttl time.Duration) Signature {
	return &signature{
		key:    key,
		secret: secret,
		ttl:    ttl,
	}
}

func (s *signature) i() {

}
