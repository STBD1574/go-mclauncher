package user

type UserUpdater struct {
	// 创建一个定时schedule，每隔x分钟执行一次
	Delay int
	User  *X19User
}

/*
UserUpdater实例化
delay 单位为秒
user 用户信息指针
*/
func NewUserUpdater(delay int, user *X19User) *UserUpdater {
	return &UserUpdater{
		Delay: delay,
		User:  user,
	}
}

func (u *UserUpdater) Start() {

}

func (u *UserUpdater) Update() {
	// 待实现
}

func (u *UserUpdater) Stop() {

}
