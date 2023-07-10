package task

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"k8s.io/klog/v2"
	"os"
)

// sessionKey is the environment variable for the session key.
var sessionKey = os.Getenv(SessionKey)

// Decrypt decrypts the content using the session key.
func Decrypt(content []byte) ([]byte, error) {
	if len(sessionKey) == 0 {
		return content, nil
	} else {
		key, _ := base64.StdEncoding.DecodeString(sessionKey)
		ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
		c, _ := aes.NewCipher(key)
		nonceSize := 16
		gcm, _ := cipher.NewGCMWithNonceSize(c, nonceSize)
		nonce, message := ciphertext[:nonceSize], ciphertext[nonceSize:]
		return gcm.Open(nil, nonce, message, nil)
	}
}

// Encrypt encrypts the content using the session key.
func Encrypt(content []byte) ([]byte, error) {
	if len(sessionKey) == 0 {
		return content, nil
	} else {
		key, _ := base64.StdEncoding.DecodeString(sessionKey)
		c, _ := aes.NewCipher(key)
		nonce := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			klog.Fatalf("Cannot encrypt data: [%v]", err)
			return nil, err
		}
		gcm, err := cipher.NewGCMWithNonceSize(c, 16)
		if err != nil {
			klog.Fatalf("Cannot encrypt data: [%v]", err)
			return nil, err
		}
		encryptedBytes := gcm.Seal(nil, nonce, content, nil)
		ciphertext := base64.StdEncoding.EncodeToString(append(nonce, encryptedBytes...))
		return []byte(ciphertext), nil
	}
}
