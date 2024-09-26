package dto

type ApiResponse struct {
	Key     string      `json:"key"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
