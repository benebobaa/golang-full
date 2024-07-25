package err

import "errors"

var (
	ErrDuplicateID = errors.New("error duplicate id")
	ErrNotFound    = errors.New("error not found")
)
