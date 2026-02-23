package x19

import (
	"go-mclauncher/api"
	"go-mclauncher/session/x19/user"
)

type X19Session struct {
	ServerList *api.ServerList

	// 登录后的用户信息
	user      *user.X19User
	apiClient *X19APIClient
}

func (s *X19Session) Login() {

}

func NewX19Session(serverList *api.ServerList) *X19Session {
	return &X19Session{
		ServerList: serverList,
	}
}
