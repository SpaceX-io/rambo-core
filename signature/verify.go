package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/kirintang/rambo-core/parse"
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"time"
)

func (s *signature) Verify(authorization, date, path, method string, params url.Values) (ok bool, err error) {
	if date == "" {
		err = errors.New("date required")
		return
	}

	if path == "" {
		err = errors.New("path required")
		return
	}

	if method == "" {
		err = errors.New("method required")
		return
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = errors.New("method param error")
		return
	}

	ts, err := parse.ParseCSTInLocation(date)
	if err != nil {
		err = errors.New("date format error")
		return
	}

	if parse.SubInLocation(ts) > float64(s.ttl/time.Second) {
		err = errors.Errorf("date exceeds limit %v", s.ttl)
		return
	}

	sortParamsEncode, err := url.QueryUnescape(params.Encode())
	if err != nil {
		err = errors.Errorf("url QueryUnescape error %v", err)
		return
	}

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(path)
	buffer.WriteString(delimiter)
	buffer.WriteString(methodName)
	buffer.WriteString(delimiter)
	buffer.WriteString(sortParamsEncode)
	buffer.WriteString(delimiter)
	buffer.WriteString(date)

	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	ok = authorization == fmt.Sprintf("%s %s", s.key, digest)
	return
}
