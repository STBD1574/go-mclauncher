package user

/*
X19 用户登录后的用户信息

	ID 用户ID
	Token 用户登录后的Token（应定期刷新，防止失效）
*/
type X19User struct {
	ID    string
	Token string
}
