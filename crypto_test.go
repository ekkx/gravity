package gravity

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	text := "hello world"

	encrypted, _ := encrypt(text)
	decrypted, _ := decrypt(encrypted)

	if text != decrypted {
		t.Fatal("Encrypt(), plan text and decrypted text do not match.")
	}
}
