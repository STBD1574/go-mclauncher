package api

import (
	"go-mclauncher/api"
	"testing"
)

func getFullUrl(t *testing.T, rc *api.RequestConfig, serverList *api.ServerList) {
	fullURL, err := rc.GetFullURL(serverList)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Full URL: %s", fullURL)
}

func TestRequestConfig(t *testing.T) {
	serverList := &api.ServerList{
		APIGatewayURL: "https://api.gateway.url",
		ForumURL:      "https://forum.url",
	}

	t.Run("Get_1", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ApiGatewayUrl",
			Method:   "GET",
		}

		getFullUrl(t, rc, serverList)
	})

	t.Run("Get_2", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:      "/test",
			HostName: "ForumUrl",
			Method:   "GET",
		}

		getFullUrl(t, rc, serverList)
	})

	t.Run("Get_3", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:    "https://www.baidu.com/test",
			Method: "GET",
		}

		getFullUrl(t, rc, serverList)
	})

	t.Run("Get_by_host", func(t *testing.T) {
		rc := &api.RequestConfig{
			URL:    "/test",
			Host:   "https://www.baidu.com",
			Method: "GET",
		}

		getFullUrl(t, rc, serverList)
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

		getFullUrl(t, rc, serverList)
	})
}
