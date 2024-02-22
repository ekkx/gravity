package gravity

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
)

const (
	KEY      = "baisimeji9262019"
	IV       = "qrstuvwxyz123456"
	T_SECRET = 999983
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

func encrypt(text string) string {
	block, _ := aes.NewCipher([]byte(KEY))

	ecb := cipher.NewCBCEncrypter(block, []byte(IV))
	content := []byte(text)
	content = pad(content)

	ciphertext := make([]byte, len(content))
	ecb.CryptBlocks(ciphertext, content)

	return base64.StdEncoding.EncodeToString(ciphertext)
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

func generateSignature(data map[string]string) (string, error) {
	delete(data, "sign")
	ts, err := strconv.Atoi(data["ts"])
	if err != nil {
		return "", err
	}

	data["t_secret"] = strconv.Itoa(ts % T_SECRET)
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	text := ""
	for _, k := range keys {
		text += k + "=" + data[k]
	}

	hasher := md5.New()
	hasher.Write([]byte(text))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
