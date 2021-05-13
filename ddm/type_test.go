package ddm

import (
	"encoding/json"
	"testing"
)

type message struct {
	Name     IDName   `json:"name"`
	Mobile   Mobile   `json:"mobile"`
	IDCard   IDCard   `json:"id_card"`
	Password Password `json:"password"`
	Email    Email    `json:"email"`
	BankCard BandCard `json:"bank_card"`
}

func TestMarshalJson(t *testing.T) {
	msg := new(message)
	msg.Name = IDName("李连杰")
	msg.IDCard = IDCard("110120999912345678")
	msg.Mobile = Mobile("18888888888")
	msg.Password = Password("123456")
	msg.BankCard = BandCard("61022799390971599")
	msg.Email = Email("tangchunlinit@google.com")

	marshal, _ := json.Marshal(msg)
	t.Log(string(marshal))
}
