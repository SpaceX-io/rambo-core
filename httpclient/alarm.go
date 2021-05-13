package httpclient

import (
	"bufio"
	"bytes"
	"go.uber.org/zap"
)

type AlarmVerify func(body []byte) (shouldAlarm bool)

type AlarmObj interface {
	Send(subject, body string) error
}

func onFailedAlarm(title string, raw []byte, logger *zap.Logger, alarmObj AlarmObj) {
	buf := bytes.NewBuffer(nil)

	scanner := bufio.NewScanner(bytes.NewReader(raw))
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
		buf.WriteString("<br />")
	}

	if err := alarmObj.Send(title, buf.String()); err != nil && logger != nil {
		logger.Error("calls failed alarm error", zap.Error(err))
	}
}
