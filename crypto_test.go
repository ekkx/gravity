package gravity

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	text := "hello world"

	encrypted := encrypt(text)
	decrypted, _ := decrypt(encrypted)

	if text != decrypted {
		t.Fatal("encrypt(), plan text and decrypted text do not match.")
	}
}

func TestGenerateSignature(t *testing.T) {
	data := map[string]string{
		"address":          "ux8S1csO8VY/H44REw/LuA==",
		"sim_country":      "JP",
		"device":           "android",
		"system_version":   "7.1.2",
		"sign":             "",
		"referrer":         "Organic",
		"zone":             "9",
		"idfa":             "eOl0xGlfS0H4qZ7ajwg7WlXMb2xJ7ec6u4C0D1QbtmhSFB58J4FMzKOgyqRTpdzP",
		"app_version_code": "375",
		"ts":               "1708600965",
		"sys_lang":         "ja",
		"app_version":      "9.2.0",
		"languageV2":       "ja",
		"uwd":              "juOwtzKZsQHwMov+aUW9MQ==",
		"country":          "JP",
		"brand":            "samsung",
		"sdk_version":      "25",
		"user_country":     "",
		"pkg":              "anonymous.sns.community.gravity",
		"product":          "gravity",
		"model":            "SM-G965N",
	}

	sign, err := generateSignature(data)
	if err != nil {
		t.Fatal(err)
	}

	if sign != "5b2babe96ec1933c6aed9f1025670512" {
		t.Fatal("generateSignature(), sign check failed.")
	}
}
