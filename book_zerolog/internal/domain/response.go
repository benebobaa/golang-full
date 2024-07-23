package domain

type BaseResponse[T any] struct {
	Data  T       `json:"data"`
	Error *string `json:"error"`
}
