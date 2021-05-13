package token

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func (t *token) UrlSign(path, method string, params url.Values) (tokenString string, err error) {
	methods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"PATH":    true,
		"DELETE":  true,
		"HEAD":    true,
		"OPTIONS": true,
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = errors.New("method param error")
		return
	}

	sortParamsEncode := params.Encode()

	encryptStr := fmt.Sprintf("%s%s%s%s", path, methodName, sortParamsEncode, t.secret)

	s := md5.New()
	s.Write([]byte(encryptStr))
	md5Str := hex.EncodeToString(s.Sum(nil))

	tokenString = base64.StdEncoding.EncodeToString([]byte(md5Str))

	return
}
