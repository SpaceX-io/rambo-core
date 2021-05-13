package mail

import "testing"

func TestSend(t *testing.T) {
	options := &Options{
		Host:    "smtp.163.com",
		Port:    465,
		User:    "xxx@163.com",
		Pass:    "",
		To:      "xxx@qq.com",
		Subject: "subject",
		Body:    "body",
	}

	err := Send(options)
	if err != nil {
		t.Error("send mail err", err)
		return
	}

	t.Log("send success")
}
