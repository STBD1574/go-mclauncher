package api

import (
	"go-mclauncher/api"
	"go-mclauncher/session/x19"
	"testing"
)

func parse(t *testing.T, rc *api.RequestConfig, serverList *api.ServerList) {
	url, method, headers, body, err := rc.Transform(serverList, false, "", "", &x19.X19Encryptor{})
	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	t.Logf("URL: %s", url)
	t.Logf("Method: %s", method)
	t.Logf("Headers: %v", headers)
	t.Logf("Body: %s", string(body))

}

func TestRequestConfig(t *testing.T) {
	serverList := &api.ServerList{
		APIGatewayURL: "https://api.gateway.url",
		ForumURL:      "https://forum.url",
	}

	t.Run("Get_with_url", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "GET",
		}

		parse(t, rc, serverList)
	})

	t.Run("Get_with_url_2", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ForumUrl",
			Method:   "GET",
		}

		parse(t, rc, serverList)
	})

	t.Run("Get_with_host", func(t *testing.T) {
		rc := &api.RequestConfig{
			Host:   "https://www.baidu.com/test",
			Method: "GET",
		}

		parse(t, rc, serverList)
	})

	t.Run("Get_with_host_and_url", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:    "/test",
			Host:   "https://www.baidu.com",
			Method: "GET",
		}

		parse(t, rc, serverList)
	})

	t.Run("Get_with_params", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "GET",
			Params: map[string]interface{}{
				"param1": "value1",
				"param2": "value2",
			},
		}

		parse(t, rc, serverList)
	})

	t.Run("Post_with_struct", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "POST",
			Params: struct {
				Key string `json:"key"`
			}{
				Key: "value",
			},
		}

		parse(t, rc, serverList)
	})

	t.Run("Post_with_bytes", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "POST",
			Params:   []byte(`{"key":"value"}`),
		}

		parse(t, rc, serverList)
	})

	t.Run("Post_encrypt", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "POST",
			Params: struct {
				Key string `json:"key"`
			}{
				Key: "value",
			},
			NeedEncrypt: true,
		}

		parse(t, rc, serverList)
	})
}
