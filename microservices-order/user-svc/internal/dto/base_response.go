package dto

type BaseResponse[T any] struct {
	Error string `json:"error"`
	Data  T      `json:"data"`
}
