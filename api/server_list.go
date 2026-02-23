package api

/*
服务器列表
键对应RequestConfig中HostName
*/
type ServerList struct {
	CdnURL                     string `json:"CdnUrl"`
	StaticWebVersionURL        string `json:"StaticWebVersionUrl"`
	StaticWeb2VersionURL       string `json:"StaticWeb2VersionUrl"`
	SeadraURL                  string `json:"SeadraUrl"`
	EmbedWebPageURL            string `json:"EmbedWebPageUrl"`
	NewsVideo                  string `json:"NewsVideo"`
	GameCenter                 string `json:"GameCenter"`
	VideoPrefix                string `json:"VideoPrefix"`
	ComponentCenter            string `json:"ComponentCenter"`
	GameDetail                 string `json:"GameDetail"`
	CompDetail                 string `json:"CompDetail"`
	LiveURL                    string `json:"LiveUrl"`
	ForumURL                   string `json:"ForumUrl"`
	WebServerURL               string `json:"WebServerUrl"`
	WebServerGrayURL           string `json:"WebServerGrayUrl"`
	CoreServerURL              string `json:"CoreServerUrl"`
	TransferServerURL          string `json:"TransferServerUrl"`
	PeTransferServerURL        string `json:"PeTransferServerUrl"`
	PeTransferServerHTTPURL    string `json:"PeTransferServerHttpUrl"`
	TransferServerHTTPURL      string `json:"TransferServerHttpUrl"`
	PeTransferServerNewHTTPURL string `json:"PeTransferServerNewHttpUrl"`
	AuthServerURL              string `json:"AuthServerUrl"`
	AuthServerCppURL           string `json:"AuthServerCppUrl"`
	AuthorityURL               string `json:"AuthorityUrl"`
	CustomerServiceURL         string `json:"CustomerServiceUrl"`
	ChatServerURL              string `json:"ChatServerUrl"`
	PathNURL                   string `json:"PathNUrl"`
	PePathNURL                 string `json:"PePathNUrl"`
	MgbSdkURL                  string `json:"MgbSdkUrl"`
	DCWebURL                   string `json:"DCWebUrl"`
	APIGatewayURL              string `json:"ApiGatewayUrl"`
	APIGatewayGrayURL          string `json:"ApiGatewayGrayUrl"`
	PlatformURL                string `json:"PlatformUrl"`
	RentalTransferURL          string `json:"RentalTransferUrl"`
}
