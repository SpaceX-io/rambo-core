package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"testing"
)

func TestJsonLogger(t *testing.T) {
	logger, err := NewJSONLogger(WithField("log_key", "log_value"))

	if err != nil {
		t.Fatal(err)
	}

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	err = errors.New("test pkg err")
	logger.Error("err occurred", WrapMeta(nil, NewMeta("log_key1", "log_value1"), NewMeta("log_key2", "log_value2"))...)
	logger.Error("err occurred", WrapMeta(err, NewMeta("log_key1", "log_value1"), NewMeta("log_key2", "log_value2"))...)
}
