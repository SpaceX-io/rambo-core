package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"runtime"
)

var _ Error = (*item)(nil)
var _ fmt.Formatter = (*item)(nil)

type item struct {
	msg   string
	stack []uintptr
}

type Error interface {
	i()
	error
}

func callers() []uintptr {
	var pcs [32]uintptr
	l := runtime.Callers(3, pcs[:])

	return pcs[:l]
}

func (i *item) Error() string {
	return i.msg
}

func (i *item) Format(s fmt.State, verb rune) {
	_, _ = io.WriteString(s, i.msg)
	_, _ = io.WriteString(s, "\n")

	for _, pc := range i.stack {
		_, _ = fmt.Fprintf(s, "%+v\n", errors.Frame(pc))
	}
}

func (i *item) i() {}

func New(msg string) Error {
	return &item{
		msg:   msg,
		stack: callers(),
	}
}

func ErrorF(format string, args ...interface{}) Error {
	return &item{msg: fmt.Sprintf(format, args...), stack: callers()}
}

func Wrap(err error, msg string) Error {
	if err == nil {
		return nil
	}

	e, ok := err.(*item)
	if !ok {
		return &item{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	}

	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)

	return e
}

func WrapF(err error, format string, args ...interface{}) Error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	e, ok := err.(*item)
	if !ok {
		return &item{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	}

	e.msg = fmt.Sprintf("%s; %s", msg, e.msg)

	return e
}

func WithStack(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*item); ok {
		return e
	}

	return &item{msg: err.Error(), stack: callers()}
}
