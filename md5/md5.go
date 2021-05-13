package md5

import (
	_md5 "crypto/md5"
	"encoding/hex"
)

var _ MD5 = (*md5)(nil)

type md5 struct{}

type MD5 interface {
	i()
	Encrypt(encryptStr string) string
}

func New() MD5 {
	return &md5{}
}

func (m *md5) i() {}

func (m *md5) Encrypt(encrypt string) string {
	s := _md5.New()
	s.Write([]byte(encrypt))

	return hex.EncodeToString(s.Sum(nil))
}
