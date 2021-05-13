package errno

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	i()
	WithErr(err error) Error
	GetBusinessNo() int
	GetHttpCode() int
	GetMsg() string
	GetErr() error
	ToString() string
}

type err struct {
	HttpCode   int
	BusinessNo int
	Message    string
	Err        error
}

func NewError(httpCode, businessNo int, msg string) Error {
	return &err{
		HttpCode:   httpCode,
		BusinessNo: businessNo,
		Message:    msg,
	}
}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessNo() int {
	return e.BusinessNo
}

func (e *err) GetMsg() string {
	return e.Message
}

func (e *err) GetErr() error {
	return e.Err
}

func (e *err) ToString() string {
	err := &struct {
		HttpCode     int    `json:"http_code"`
		BusinessCode int    `json:"business_code"`
		Message      string `json:"message"`
	}{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessNo,
		Message:      e.Message,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}

func (e *err) i() {}
