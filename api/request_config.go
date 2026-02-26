package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
)

/*
RequestConfig API请求统一配置
*/
type RequestConfig struct {
	URL         string
	HostName    string // full URL
	Host        string
	Params      interface{}
	Method      string
	Headers     map[string]string
	HPack       bool
	NoBody      bool
	Signal      bool
	NeedEncrypt bool
}

func (rc *RequestConfig) String() string {
	return fmt.Sprintf("RequestConfig{URL: %s, HostName: %s, Params: %v, Method: %s, Headers: %v, HPack: %t, NoBody: %t, Signal: %t, NeedEncrypt: %t}",
		rc.URL, rc.HostName, rc.Params, rc.Method, rc.Headers, rc.HPack, rc.NoBody, rc.Signal, rc.NeedEncrypt)
}

/*
获取URL
*/
func (rc *RequestConfig) getURL() (string, error) {
	urlValue := rc.URL

	if rc.Host != "" {
		urlParser, err := url.Parse(rc.Host)
		if err != nil {
			return "", err
		}

		urlValue = urlParser.Path
	}

	if rc.Method == "GET" && rc.Params != nil {
		query := ""

		q := url.Values{}
		for k, v := range rc.Params.(map[string]interface{}) {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		query = q.Encode()

		if query != "" {
			if strings.Contains(urlValue, "?") {
				urlValue += "&" + query
			} else {
				urlValue += "?" + query
			}
		}
	}

	return urlValue, nil
}

/*
获取请求域名
*/
func (rc *RequestConfig) getHostName(serverList *ServerList) (string, error) {
	var err error
	hostName := rc.HostName

	if rc.Host != "" {
		urlParser, err := url.Parse(rc.Host)
		if err != nil {
			return "", err
		}

		return urlParser.Scheme + "://" + urlParser.Host, nil
	}

	if strings.HasPrefix(hostName, "http") {
		return hostName, nil
	}

	switch hostName {
	case "momentHost":
		hostName = "MomentUrl"
	case "HomeServerUrl", "ApiGatewayUrl":
		hostName = serverList.APIGatewayURL
	case "WebServerUrl":
		hostName = serverList.WebServerURL
	default:
		hostName, err = serverList.Get(hostName)
		if err != nil {
			return "", err
		}
	}

	return hostName, nil
}

/*
获取请求头
*/
func (rc *RequestConfig) getHeaders(isGray bool) map[string]string {
	headers := map[string]string{
		"user-token":   "",
		"user-id":      "",
		"Content-Type": "application/json",
		"X_TRACE_ID":   generateTraceID(),
	}

	for k, v := range rc.Headers {
		headers[k] = v
	}

	if rc.HPack {
		headers["hpack"] = "1"
	}

	if rc.NoBody {
		headers["no_body"] = "1"
	}

	if rc.HostName == "HomeServerUrl" {
		headers["home"] = "1"
	}

	if isGray {
		headers["is_gray"] = "1"
	}

	return headers
}

/*
转换为HTTP请求所需的信息
*/
func (rc *RequestConfig) Transform(serverList *ServerList, isGray bool, userID string, userToken string, encryptor Encryptor) (string, string, map[string]string, []byte, error) {
	var err error
	url, err := rc.getURL()
	if err != nil {
		return "", "", nil, nil, err
	}

	hostName, err := rc.getHostName(serverList)
	if err != nil {
		return "", "", nil, nil, err
	}

	headers := rc.getHeaders(isGray)
	body := []byte{}
	if rc.Method == "POST" && rc.Params != nil {
		body, err = json.Marshal(rc.Params)
		if err != nil {
			return "", "", nil, nil, err
		}
	}

	encryptedToken, err := encryptor.UserTokenEncrypt(url, body)
	if err != nil {
		return "", "", nil, nil, err
	}

	headers["user-token"] = string(encryptedToken)
	headers["user-id"] = userID
	if rc.NeedEncrypt {
		body, err = encryptor.HttpEncrypt(body)
		if err != nil {
			return "", "", nil, nil, err
		}
	}

	return hostName + url, rc.Method, headers, body, nil
}

/*
进行加密
*/
func encrypt(string url, body, Encryptor encryptor) (string, string, []byte, error) {
	encryptedToken, err := encryptor.UserTokenEncrypt(url, body)
	if err != nil {
		return "", "", nil, nil, err
	}

}

/*
随机生成TraceID
*/
func generateTraceID() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var n = 32

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
