package interfaces

import "context"

type Saver[T any] interface {
	Save(ctx context.Context, value *T) (*T, error)
}
