package api

import "fmt"

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

/*
获取指定服务器URL
*/
func (s *ServerList) Get(hostName string) (string, error) {
	switch hostName {

	case CdnURL:
		return s.CdnURL, nil
	case StaticWebVersionURL:
		return s.StaticWebVersionURL, nil
	case StaticWeb2VersionURL:
		return s.StaticWeb2VersionURL, nil
	case SeadraURL:
		return s.SeadraURL, nil
	case EmbedWebPageURL:
		return s.EmbedWebPageURL, nil
	case NewsVideo:
		return s.NewsVideo, nil
	case GameCenter:
		return s.GameCenter, nil
	case VideoPrefix:
		return s.VideoPrefix, nil
	case ComponentCenter:
		return s.ComponentCenter, nil
	case GameDetail:
		return s.GameDetail, nil
	case CompDetail:
		return s.CompDetail, nil
	case LiveURL:
		return s.LiveURL, nil
	case ForumURL:
		return s.ForumURL, nil
	case WebServerURL:
		return s.WebServerURL, nil
	case WebServerGrayURL:
		return s.WebServerGrayURL, nil
	case CoreServerURL:
		return s.CoreServerURL, nil
	case TransferServerURL:
		return s.TransferServerURL, nil
	case PeTransferServerURL:
		return s.PeTransferServerURL, nil
	case PeTransferServerHTTPURL:
		return s.PeTransferServerHTTPURL, nil
	case TransferServerHTTPURL:
		return s.TransferServerHTTPURL, nil
	case PeTransferServerNewHTTPURL:
		return s.PeTransferServerNewHTTPURL, nil
	case AuthServerURL:
		return s.AuthServerURL, nil
	case AuthServerCppURL:
		return s.AuthServerCppURL, nil
	case AuthorityURL:
		return s.AuthorityURL, nil
	case CustomerServiceURL:
		return s.CustomerServiceURL, nil
	case ChatServerURL:
		return s.ChatServerURL, nil
	case PathNURL:
		return s.PathNURL, nil
	case PePathNURL:
		return s.PePathNURL, nil
	case MgbSdkURL:
		return s.MgbSdkURL, nil
	case DCWebURL:
		return s.DCWebURL, nil
	case APIGatewayURL:
		return s.APIGatewayURL, nil
	case APIGatewayGrayURL:
		return s.APIGatewayGrayURL, nil
	case PlatformURL:
		return s.PlatformURL, nil
	case RentalTransferURL:
		return s.RentalTransferURL, nil
	}

	return "", fmt.Errorf("HostName %s not found in ServerList", hostName)
}
