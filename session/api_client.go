package session

import (
	"context"
	"go-mclauncher/api"
	"go-mclauncher/entities"
	"net/http"
)

/*
APIClient 与API进行交互的接口

	DoRequest 以此APIClient的方式进行API请求，返回HTTP响应
	EntityRequest 同上，返回实体响应
	EntityListRequest 同上，返回实体列表响应
*/
type APIClient interface {
	DoRequest(ctx context.Context, rc *api.RequestConfig) (resp *http.Response, err error)
	EntityRequest(ctx context.Context, rc *api.RequestConfig) (resp *entities.EntityResponse[any], err error)
	EntityListRequest(ctx context.Context, rc *api.RequestConfig) (resp *entities.EntityListResponse[any], err error)
}
