package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultLevel      = zapcore.InfoLevel
	DefaultTimeFormat = time.RFC3339
)

type option struct {
	level          zapcore.Level
	fields         map[string]string
	file           io.Writer
	timeFormat     string
	disableConsole bool
}

type Option func(*option)

func WithInfoLevel() Option {
	return func(o *option) {
		o.level = zapcore.InfoLevel
	}
}

func WithDebugLevel() Option {
	return func(o *option) {
		o.level = zapcore.DebugLevel
	}
}

func WithWarnLevel() Option {
	return func(o *option) {
		o.level = zapcore.WarnLevel
	}
}

func WithErrorLevel() Option {
	return func(o *option) {
		o.level = zapcore.ErrorLevel
	}
}

func WithField(key, value string) Option {
	return func(o *option) {
		o.fields[key] = value
	}
}

func WithFile(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0776)
	if err != nil {
		panic(err)
	}

	return func(o *option) {
		o.file = zapcore.Lock(f)
	}
}

func WithFileRotation(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0776); err != nil {
		panic(err)
	}

	return func(o *option) {
		o.file = &lumberjack.Logger{
			Filename:   file,
			MaxSize:    128,
			MaxBackups: 256,
			MaxAge:     30,
			LocalTime:  true,
			Compress:   true,
		}
	}
}

func WithTimeFormat(timeFormat string) Option {
	return func(o *option) {
		o.timeFormat = timeFormat
	}
}

func WithDisableConsole() Option {
	return func(o *option) {
		o.disableConsole = true
	}
}

func NewJSONLogger(opts ...Option) (*zap.Logger, error) {
	opt := &option{level: DefaultLevel, fields: make(map[string]string)}
	for _, f := range opts {
		f(opt)
	}

	timeFormat := DefaultTimeFormat
	if opt.timeFormat != "" {
		timeFormat = opt.timeFormat
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(timeFormat))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= opt.level && level < zapcore.ErrorLevel
	})

	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= opt.level && level >= zapcore.ErrorLevel
	})

	stdout := zapcore.Lock(os.Stdout)
	stderr := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee()

	if !opt.disableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stdout),
				lowPriority,
			),
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stderr),
				highPriority,
			),
		)
	}

	if opt.file != nil {
		core = zapcore.NewTee(core,
			zapcore.NewCore(jsonEncoder,
				zapcore.AddSync(opt.file),
				zap.LevelEnablerFunc(func(level zapcore.Level) bool {
					return level >= opt.level
				}),
			),
		)
	}

	logger := zap.New(core, zap.AddCaller(), zap.ErrorOutput(stderr))

	for key, value := range opt.fields {
		logger = logger.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
	}

	return logger, nil
}

var _Meta = (*meta)(nil)

type Meta interface {
	Key() string
	Value() interface{}
	m()
}

type meta struct {
	key   string
	value interface{}
}

func (m *meta) Key() string {
	return m.key
}

func (m *meta) Value() interface{} {
	return m.value
}

func (m *meta) m() {

}

func NewMeta(key string, value interface{}) Meta {
	return &meta{
		key:   key,
		value: value,
	}
}

func WrapMeta(err error, metas ...Meta) (fields []zap.Field) {
	capacity := len(metas) + 1
	if err != nil {
		capacity++
	}

	fields = make([]zap.Field, 0, capacity)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	fields = append(fields, zap.Namespace("meta"))
	for _, meta := range metas {
		fields = append(fields, zap.Any(meta.Key(), meta.Value()))
	}

	return
}
