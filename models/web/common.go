package web

type WithDataResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}