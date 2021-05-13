package signature

import (
	"net/url"
	"testing"
	"time"
)

const (
	key    = "rambo"
	secret = "242BB39A78E1251A"
	ttl    = time.Minute * 10
)

func TestGenerate(t *testing.T) {
	path := "/user/add"
	method := "POST"

	params := url.Values{}
	params.Add("user_name", "SpaceX-io")
	params.Add("email", "xxxx@163.com")
	params.Add("age", "25")

	authorization, date, err := New(key, secret, ttl).Generate(path, method, params)
	t.Log("authorization", authorization)
	t.Log("date", date)
	t.Log("err:", err)
}

func TestVerify(t *testing.T) {
	authorization := "rambo 3ePsgN6FSv8xV+YrRD9jwPJvggDPsvRTM1u9oZcmMjI="
	date := "2021-05-06 13:28:20"

	path := "/user/add"
	method := "post"
	params := url.Values{}
	params.Add("user_name", "SpaceX-io")
	params.Add("email", "xxxx@163.com")
	params.Add("age", "25")

	ok, err := New(key, secret, ttl).Verify(authorization, date, path, method, params)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(ok)
}
