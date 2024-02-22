package gravity

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
)

// getstrts() converets the given timestamp into a string.
func getstrts(ts int64) string {
	return strconv.FormatUint(uint64(ts), 10)
}

var (
	LoginTypeInvalid = -1
	LoginTypeOther   = 0
	LoginTypePnum    = 1
	LoginTypeEmail   = 2
)

// getLoginType() returns the type of identifier.
func getLoginType(identifier string) int {
	return LoginTypeEmail
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func structToMapWithJSON(data interface{}) map[string]string {
	result := make(map[string]interface{})
	strResult := make(map[string]string)

	b, _ := Marshal(data)
	Unmarshal(b, &result)

	for key, value := range result {
		if value != nil && value != "" {
			strResult[key] = fmt.Sprintf("%v", value)
		}
	}

	return strResult
}
