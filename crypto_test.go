package gravity

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	text := "hello world"

	encrypted, _ := encrypt(text)
	decrypted, _ := decrypt(encrypted)

	if text != decrypted {
		t.Fatal("encrypt(), plan text and decrypted text do not match.")
	}
}

func TestGenerateSignature(t *testing.T) {
	data := map[string]string{
		"address":          "8q/ux729nh0NyK+8o+wSp4YcUzqFdPvTMIWuQYmdmG0=",
		"sim_country":      "JP",
		"device":           "android",
		"system_version":   "7.1.2",
		"sign":             "df8dddc15430b83284314fc36528648d",
		"referrer":         "Organic",
		"zone":             "9",
		"idfa":             "oP6bmU4SPL/o+93TSNKGJYa+SVqsNnHgif8N7MzDbwQ9iIOD0089hb74OnFigAch",
		"app_version_code": "375",
		"ts":               "1708260764",
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

	if sign != "d01bc8090834ef38d5e4888a97c91f0c" {
		t.Fatal("generateSignature(), sign check failed.")
	}
}
