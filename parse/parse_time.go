package parse

import (
	"math"
	"net/http"
	"time"
)

var cst *time.Location

const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
}

func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

func CSTLayoutString() string {
	ts := time.Now()

	return ts.In(cst).Format(CSTLayout)
}

func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

func CSTLayoutStringToUnix(cstLayoutStr string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutStr, cst)
	if err != nil {
		return 0, err
	}

	return stamp.Unix(), nil
}

func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}
