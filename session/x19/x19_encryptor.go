package x19

import "go-mclauncher/security/x19"

type X19Encryptor struct {
}

func (e *X19Encryptor) UserTokenEncrypt(data []byte) ([]byte, error) {
	x19.X19ComputeDynamicToken()

}
