package aes

import "testing"

const (
	key = "242BB39A78E1251A"
	iv  = "242BB39A78E1251A"
	str = "rambo"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key, iv).Encrypt(str))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key, iv).Decrypt("UQRV5hVzz-5tKuEcsPVg5Q=="))
}

func BenchMarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key, iv)
	for i := 0; i < b.N; i++ {
		encryptStr, _ := aes.Encrypt(str)
		_, _ = aes.Decrypt(encryptStr)
	}
}
