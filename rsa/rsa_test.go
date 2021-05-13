package rsa

import (
	"testing"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANjppjjOuhkOr2aNaD4r6Ko9OxWOo7dy
mwujQyqWdOW8OsZx3kYgBS7Q1e0B3ffmWH6gook7wJnNUvOjfPYNrn0CAwEAAQ==
-----END PUBLIC KEY-----`

const privateKey = `-----BEGIN PRIVATE KEY-----
MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA2OmmOM66GQ6vZo1o
Pivoqj07FY6jt3KbC6NDKpZ05bw6xnHeRiAFLtDV7QHd9+ZYfqCiiTvAmc1S86N8
9g2ufQIDAQABAkAkPEPsUXx9GxrqAs1bNXKUnc309/MZfiewdgGOZ7v3dH9qbrzX
wJwHeWC+CZIL1eSR4R1KFw1jxiJFdAYy1Q1xAiEA+tpv89/mfac9Q/U/QO0n0a/+
Uq8NQPjB2aJleEywkocCIQDdXPHwO9H+UIvNSA0QVdpL4vwfg6fStsFxAKgXxOFD
2wIgeiP9urrcGXZiqEIzaEOQzdJpfIzrYSU+Dd+6lFaS6uUCIFe2LGd0TJDoeXyt
v/9pBUZselpCYI0tvRh5miFQ8bFhAiAdNlg/JnXhuKD0tsPi0TSrAtKrM3uTxjjZ
CwBvmT98Rg==
-----END PRIVATE KEY-----`

func TestEncrypt(t *testing.T) {
	str, err := NewPublic(publicKey).Encrypt("rambo")
	if err != nil {
		t.Error("rsa encrypt error", err)
		return
	}

	t.Log(str)
}

func TestDecrypt(t *testing.T) {
	decryptStr := "Z2JIO8_kBIAfr0bNaVUhA_CDvlWXvI1TXAcPbgN6CkpZbl2xGZYVidl56mwTNyKOa-NcFYNrdA1OrPo8OLmo6A=="

	str, err := NewPrivate(privateKey).Decrypt(decryptStr)
	if err != nil {
		t.Error("rsa decrypt error", err)
		return
	}

	t.Log(str)
}
