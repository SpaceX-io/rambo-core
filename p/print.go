package p

import (
	"fmt"
	"time"

	"github.com/SpaceX-io/rambo-core/trace"
)

type Option func(*option)

type Trace = trace.T

type option struct {
	Trace *trace.Trace
	Debug *trace.Debug
}

func newOption() *option {
	return &option{}
}

func Println(key string, value interface{}, options ...Option) {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Debug.Key = key
			opt.Debug.Value = value
			opt.Debug.Cost = time.Since(ts).Seconds()
			opt.Trace.AppendDebug(opt.Debug)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	fmt.Println(fmt.Sprintf("KEY: %s | VALUE: %v", key, value))
}

func WithTrace(t Trace) Option {
	return func(o *option) {
		if t != nil {
			o.Trace = t.(*trace.Trace)
			o.Debug = new(trace.Debug)
		}
	}
}
