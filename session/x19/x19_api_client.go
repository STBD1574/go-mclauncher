package x19

import (
	"bytes"
	"context"
	"encoding/json"
	"go-mclauncher/api"
	"go-mclauncher/entities"
	"io"
	"net/http"
)

type X19APIClient struct {
	Client     *http.Client
	ServerList *api.ServerList
}

func (c *X19APIClient) DoRequest(ctx context.Context, reqConfig *api.RequestConfig) (*http.Response, error) {
	var buf io.Reader
	if reqConfig.Params != nil {
		jsonData, err := json.Marshal(reqConfig.Params)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, reqConfig.Method, "", buf)
	if err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *X19APIClient) EntityRequest(ctx context.Context, reqConfig *api.RequestConfig) (resp *entities.EntityResponse[any], err error) {
	httpResp, err := c.DoRequest(ctx, reqConfig)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var entityResp entities.EntityResponse[any]
	decoder := json.NewDecoder(httpResp.Body)
	err = decoder.Decode(&entityResp)
	if err != nil {
		return nil, err
	}

	return &entityResp, nil
}

func (c *X19APIClient) EntityListRequest(ctx context.Context, reqConfig *api.RequestConfig) (resp *entities.EntityListResponse[any], err error) {
	httpResp, err := c.DoRequest(ctx, reqConfig)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var entityListResp entities.EntityListResponse[any]
	decoder := json.NewDecoder(httpResp.Body)
	err = decoder.Decode(&entityListResp)
	if err != nil {
		return nil, err
	}

	return &entityListResp, nil
}

func NewX19APIClient(client *http.Client, serverList *api.ServerList) *X19APIClient {
	return &X19APIClient{
		Client:     client,
		ServerList: serverList,
	}
}
