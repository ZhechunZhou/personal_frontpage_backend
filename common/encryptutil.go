package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io"
)

func Encode(data []byte) string {
	if data == nil {
		return ""
	}
	return b64.StdEncoding.EncodeToString(data)
}

func Decode(data string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(data)
}

func GetHashCode(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func EncryptFile(data []byte) ([]byte, string, error) {
	key := GetHashCode(data)
	hashcode := fmt.Sprintf("%x", key)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, "", err
	}

	cipher := gcm.Seal(nonce, nonce, data, nil)
	// Save back to file
	return cipher, hashcode, nil
}

func DecryptFile(data []byte, hashcode string) ([]byte, error) {
	key := []byte(hashcode)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := data[:gcm.NonceSize()]
	data = data[gcm.NonceSize():]
	plain, err := gcm.Open(nil, nonce, data, nil)
	return plain, err
}
