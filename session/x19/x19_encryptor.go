package x19

import "go-mclauncher/security/x19"

type X19Encryptor struct {
}

func (e *X19Encryptor) UserTokenEncrypt(url string, body []byte, token string) string {
	return x19.X19ComputeDynamicToken(url, body, token)
}

func (e *X19Encryptor) HttpEncrypt(data []byte) ([]byte, error) {
	return x19.X19HttpEncrypt(data)
}
