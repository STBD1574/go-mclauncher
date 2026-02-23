package launcher

type LoginOTPEntity struct {
	Otp      int    `json:"otp"`
	OtpToken string `json:"otp_token"`
	Aid      string `json:"aid"`
	LockTime int    `json:"lock_time"`
	OpenOtp  int    `json:"open_otp"`
}
