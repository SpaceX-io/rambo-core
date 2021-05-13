package httpclient

import (
	"github.com/kirintang/rambo-core/trace"
	"go.uber.org/zap"
	"sync"
	"time"
)

var cache = &sync.Pool{
	New: func() interface{} {
		return &option{}
	},
}

type Mock func() (body []byte)

type Option func(*option)

type option struct {
	ttl         time.Duration
	header      map[string][]string
	trace       *trace.Trace
	dialog      *trace.Dialog
	logger      *zap.Logger
	retryTimes  int
	retryDelay  time.Duration
	retryVerify RetryVerify
	alarmTitle  string
	alarmObj    AlarmObj
	alarmVerify AlarmVerify
	mock        Mock
}

func (o *option) reset() {
	o.ttl = 0
	o.header = make(map[string][]string)
	o.trace = nil
	o.dialog = nil
	o.logger = nil
	o.retryTimes = 0
	o.retryDelay = 0
	o.retryVerify = nil
	o.alarmTitle = ""
	o.alarmObj = nil
	o.alarmVerify = nil
	o.mock = nil
}

func getOption() *option {
	return cache.Get().(*option)
}

func releaseOption(o *option) {
	o.reset()
	cache.Put(o)
}

func WithTTL(ttl time.Duration) Option {
	return func(o *option) {
		o.ttl = ttl
	}
}

func WithHeader(key, value string) Option {
	return func(o *option) {
		o.header[key] = []string{value}
	}
}

func WithTrace(t trace.T) Option {
	return func(o *option) {
		if t != nil {
			o.trace = t.(*trace.Trace)
			o.dialog = new(trace.Dialog)
		}
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(o *option) {
		o.logger = logger
	}
}

func WithMock(m Mock) Option {
	return func(o *option) {
		o.mock = m
	}
}

func WithOnFailedAlarm(alarmTitle string, alarmObj AlarmObj, alarmVerify AlarmVerify) Option {
	return func(o *option) {
		o.alarmTitle = alarmTitle
		o.alarmObj = alarmObj
		o.alarmVerify = alarmVerify
	}
}

func WithOnFailedRetry(retryTimes int, retryDelay time.Duration, retryVerify RetryVerify) Option {
	return func(o *option) {
		o.retryTimes = retryTimes
		o.retryDelay = retryDelay
		o.retryVerify = retryVerify
	}
}
