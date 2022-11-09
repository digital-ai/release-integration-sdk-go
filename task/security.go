package task

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

var SessionKey = os.Getenv("SESSION_KEY")

func Decrypt(content []byte) ([]byte, error) {
	if len(SessionKey) == 0 {
		return content, nil
	} else {
		key, _ := base64.StdEncoding.DecodeString(SessionKey)
		ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
		c, _ := aes.NewCipher(key)
		nonceSize := 16
		gcm, _ := cipher.NewGCMWithNonceSize(c, nonceSize)
		nonce, message := ciphertext[:nonceSize], ciphertext[nonceSize:]
		return gcm.Open(nil, nonce, message, nil)
	}
}

func Encrypt(content []byte) []byte {
	if len(SessionKey) == 0 {
		return content
	} else {
		key, _ := base64.StdEncoding.DecodeString(SessionKey)
		c, _ := aes.NewCipher(key)
		nonce := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}
		gcm, err := cipher.NewGCMWithNonceSize(c, 16)
		if err != nil {
			panic(err.Error())
		}
		encryptedBytes := gcm.Seal(nil, nonce, content, nil)
		ciphertext := base64.StdEncoding.EncodeToString(append(nonce, encryptedBytes...))
		return []byte(ciphertext)
	}
}
