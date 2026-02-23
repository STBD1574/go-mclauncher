package api

import (
	"fmt"
	"maps"
	"math/rand"
	"reflect"
	"strings"
)

/*
RequestConfig API请求统一配置

	URL 请求URL
	HostName ReleaseInfo中的URL字段常量
	Host 直接指定域名，优先级高于HostName
	Params 请求参数 当Method为GET时会拼接URL 当Method为POST时会作为Json请求体发送
	Method HTTP方法
	Headers 自定义请求头
	HPack 是否使用HPack编码
	NoBody 是否不发送请求体
	Signal 是否请求取消/中断信号
	NeedEncrypt 是否需要加密请求体
*/
type RequestConfig struct {
	URL         string
	HostName    string // such as "WebServerUrl"
	Host        string
	Params      interface{} // Json body or URL query params
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
获取完整URL（根据HostName从ServerList获取基础URL并拼接URL）
*/
func (rc *RequestConfig) GetFullURL(serverList *ServerList) (string, error) {
	var err error
	fullURL := rc.URL

	if rc.Method == "GET" && rc.Params != nil {
		query := ""

		for key, value := range rc.Params.(map[string]interface{}) {
			query += fmt.Sprintf("%s=%v&", key, value)
		}

		if query != "" {
			query = query[:len(query)-1] // Remove the trailing '&'
			if strings.Contains(fullURL, "?") {
				fullURL += "&" + query
			} else {
				fullURL += "?" + query
			}
		}
	}

	if !strings.HasPrefix(fullURL, "http") {
		if rc.Host != "" {
			return rc.Host + fullURL, nil
		}

		hostName := rc.HostName

		switch hostName {
		case "momentHost":
			hostName = "MomentUrl"
		case "HomeServerUrl":
		case "ApiGatewayUrl":
			hostName = serverList.APIGatewayURL
		case "WebServerUrl":
			hostName = serverList.WebServerURL
		default:
			hostName, err = getURLByHostName(serverList, hostName)
			if err != nil {
				return "", err
			}
		}

		fullURL = hostName + fullURL
	}

	return fullURL, nil
}

/*
获取请求头
*/
func (rc *RequestConfig) GetHeaders(isGray bool) map[string]string {
	headers := map[string]string{
		"user-token":   "",
		"user-id":      "",
		"Content-Type": "application/json",
		"X_TRACE_ID":   generateTraceID(),
	}
	maps.Copy(rc.Headers, headers)

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
通过反射获取hostName
*/
func getURLByHostName(serverList *ServerList, hostName string) (string, error) {
	serverListValue := reflect.ValueOf(serverList).Elem()
	serverListType := serverListValue.Type()

	for i := 0; i < serverListType.NumField(); i++ {
		field := serverListType.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == hostName {
			return serverListValue.Field(i).String(), nil
		}
	}

	return "", fmt.Errorf("HostName %s not found in ServerList", hostName)
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
