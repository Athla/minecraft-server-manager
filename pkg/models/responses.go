package models

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Code    int    `json:"code"`
}
