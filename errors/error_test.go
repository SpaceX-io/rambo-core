package errors

import (
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestError(t *testing.T) {
	logger, _ := zap.NewProduction()

	err := New("an err")

	logger.Info("wrap std err", zap.Error(Wrap(errors.New("std err"), "some error occurs")))

	// WrapF
	err = WrapF(err, "ip: %s port %d", "127.0.0.1", 80)
	logger.Info("WrapF", zap.Error(err))

	// Wrap
	err = Wrap(err, "ping timeout err")
	logger.Info("Wrap", zap.Error(err))

	// WithStack
	err = WithStack(err)
	logger.Info("WithStack", zap.Error(err))

	t.Logf("%+v", New("end err test"))
}
