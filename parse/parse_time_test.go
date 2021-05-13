package parse

import (
	"testing"
)

func TestRFC3339ToCSTLayout(t *testing.T) {
	t.Log(RFC3339ToCSTLayout("2021-05-05T20:25:00+08:00")) // 2021-05-05 20:25:00 <nil>
}

func TestCSTLayoutString(t *testing.T) {
	t.Log(CSTLayoutString()) // 2021-05-05 20:29:13
}

func TestCSTLayoutStringToUnix(t *testing.T) {
	t.Log(CSTLayoutStringToUnix("2021-05-05 20:26:01")) // 1620217561 <nil>
}

func TestGMTLayoutString(t *testing.T) {
	t.Log(GMTLayoutString()) // Wed, 05 May 2021 20:28:51 GMT
}
