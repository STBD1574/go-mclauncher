package launcher

type MPaySAuthToken struct {
	GameID          string `json:"GameID"`
	LoginChannel    string `json:"LoginChannel"`
	AppChannel      string `json:"AppChannel"`
	Platform        string `json:"Platform"`
	SDKUserID       string `json:"SDKUserID"`
	SessionID       string `json:"SessionID"`
	SDKVersion      string `json:"SDKVersion"`
	Udid            string `json:"UDID"`
	DeviceID        string `json:"DeviceID"`
	ClientLoginSign string `json:"ClientLoginSign"`
	GasToken        string `json:"GasToken"`
	SourcePlatform  string `json:"SourcePlatform"`
	IP              string `json:"IP"`
}
