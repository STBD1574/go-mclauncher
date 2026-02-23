package security

import (
	"bytes"
	"math/rand"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func ToBinaryString(data []byte) string {
	var buf bytes.Buffer
	for i := 0; i < len(data); i++ {
		for j := 0; j < 8; j++ {
			buf.WriteByte('0' + (data[i] >> (7 - j) & 1))
		}
	}
	return buf.String()
}
