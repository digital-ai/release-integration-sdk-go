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

var SessionKey = os.Getenv("SESSION_KEY")

func Decrypt(content []byte) ([]byte, error) {
	if len(SessionKey) == 0 {
		return content, nil
	} else {
		// TODO remove this!
		klog.Infof("got from secret sessionKey base64: %s", SessionKey)
		klog.Infof("got from secret sessionKey base64: %s", content)
		key, _ := base64.StdEncoding.DecodeString(SessionKey)
		ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
		c, _ := aes.NewCipher(key)
		nonceSize := 16
		gcm, _ := cipher.NewGCMWithNonceSize(c, nonceSize)
		nonce, message := ciphertext[:nonceSize], ciphertext[nonceSize:]
		return gcm.Open(nil, nonce, message, nil)
	}
}

//	iv, encryptedData := encryptedData[:aes.BlockSize], encryptedData[aes.BlockSize:]
//
//	block, err := aes.NewCipher(sessionKey)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create AES cipher block: %s", err)
//	}
//
//	// GCM cipher with given nonce/iv, decrypt the encrypted data using the GCM cipher
//	gcm, err := cipher.NewGCMWithNonceSize(block, aes.BlockSize)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create GCM cipher: %s", err)
//	}
//	plaintext, err := gcm.Open(nil, iv, encryptedData, nil)
//	if err != nil {
//		return nil, fmt.Errorf("failed to decrypt ciphertext: %s", err)
//	}
//
//	return plaintext, nil
//

func Encrypt(content []byte) ([]byte, error) {
	if len(SessionKey) == 0 {
		return content, nil
	} else {
		key, _ := base64.StdEncoding.DecodeString(SessionKey)
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
