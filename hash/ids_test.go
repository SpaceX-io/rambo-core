package hash

import "testing"

const (
	secret = "242BB39A78E1251A"
	length = 16
)

func TestHashIdsEncode(t *testing.T) {
	str, _ := New(secret, length).HashIdsEncode([]int{100})
	t.Log(str) // O182oARa7nQJgjwK
}

func TestHashIdsDecode(t *testing.T) {
	ids, _ := New(secret, length).HashIdsDecode("O182oARa7nQJgjwK")
	t.Log(ids) // [100]
}
