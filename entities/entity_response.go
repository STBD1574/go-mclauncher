package entities

/*
统一实体响应
*/
type EntityResponse[T any] struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
	Entity  T      `json:"entity"`
	Message string `json:"message"`
}

/*
统一实体列表响应
*/
type EntityListResponse[T any] struct {
	Code     int    `json:"code"`
	Details  string `json:"details"`
	Entities T      `json:"entity"`
	Message  string `json:"message"`
	Total    int    `json:"total"`
}
