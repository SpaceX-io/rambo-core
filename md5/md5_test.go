package md5

import "testing"

func TestEncrypt(t *testing.T) {
	t.Log(New().Encrypt("rambo"))
}

func BenchmarkEncrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New().Encrypt("rambo")
	}
}
