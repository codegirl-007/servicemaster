package crypto

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestEncryptorRoundTrip(t *testing.T) {
	key := base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{1}, 32))

	encryptor, err := NewEncryptorFromBase64Key(key)
	if err != nil {
		t.Fatalf("new encryptor: %v", err)
	}

	plaintext := []byte("secret-token")

	encrypted, err := encryptor.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}

	if bytes.Contains(encrypted, plaintext) {
		t.Fatal("encrypted value contains plaintext")
	}

	decrypted, err := encryptor.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("decrypt: %v", err)
	}

	if !bytes.Equal(decrypted, plaintext) {
		t.Fatal("decrypted value mismatch")
	}
}

func TestNewEncryptorRejectsShortKey(t *testing.T) {
	key := base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{1}, 16))

	_, err := NewEncryptorFromBase64Key(key)
	if err == nil {
		t.Fatal("expected error")
	}
}
