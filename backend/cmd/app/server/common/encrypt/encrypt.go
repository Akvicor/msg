package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"msg/cmd/config"
)

func Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func Decode(s string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	return data, err
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(config.Global.Encrypt.Key))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(config.Global.Encrypt.Iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(config.Global.Encrypt.Key))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, []byte(config.Global.Encrypt.Iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
