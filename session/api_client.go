package session

import (
	"context"
	"go-mclauncher/api"
	"go-mclauncher/entities"
	"net/http"
)

/*
APIClient 与HTTP层进行交互的接口
*/
type APIClient interface {
	DoRequest(ctx context.Context, rc *api.RequestConfig) (resp *http.Response, err error)
	EntityRequest(ctx context.Context, rc *api.RequestConfig) (resp *entities.EntityResponse[any], err error)
	EntityListRequest(ctx context.Context, rc *api.RequestConfig) (resp *entities.EntityListResponse[any], err error)
}
