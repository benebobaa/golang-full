package errr

import "errors"

var (
	ErrDuplicateID         = errors.New("error duplicate id")
	ErrNotFound            = errors.New("error not found")
	ErrTicketOutOfStock    = errors.New("error ticket out of stock")
	ErrTicketsRequestEmpty = errors.New("error tickets request empty")
)
