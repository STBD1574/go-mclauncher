package api

/*
为API接口提供加密功能的接口
*/
type Encryptor interface {
	UserTokenEncrypt(url string, body []byte, token string) string
	HttpEncrypt(data []byte) ([]byte, error)
}
