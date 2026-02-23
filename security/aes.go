package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func AesCBCEncrypt(key []byte, data []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, len(data))
	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, data)

	return encrypted, nil
}

func AesCBCDecrypt(key []byte, data []byte, iv []byte, unPadding bool) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(data))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, data)

	if unPadding {
		num := len(decrypted) - 1
		num2 := 0
		for decrypted[num] == 0 && num > 0 {
			num--
			num2++
		}

		fmt.Printf("num: %d, num2: %d\n", num, num2)

		return decrypted[:(num-127)-20], nil
	}

	return decrypted, nil
}
