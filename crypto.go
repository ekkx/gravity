package gravity

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	KEY = "baisimeji9262019"
	IV  = "qrstuvwxyz123456"
)

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, fmt.Errorf("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(KEY))
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(IV))
	content := []byte(text)
	content = pad(content)

	ciphertext := make([]byte, len(content))
	ecb.CryptBlocks(ciphertext, content)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(KEY))
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	if len(decoded) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	ecb := cipher.NewCBCDecrypter(block, []byte(IV))
	ecb.CryptBlocks(decoded, decoded)

	unpadData, err := unpad(decoded)
	if err != nil {
		return "", err
	}

	return string(unpadData), nil
}
