package security

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Sum(data []byte) []byte {
	result := md5.Sum(data)
	return result[:]
}

func MD5Hex(input []byte) string {
	return hex.EncodeToString(MD5Sum(input))
}
