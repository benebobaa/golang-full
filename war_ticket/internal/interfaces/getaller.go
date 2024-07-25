package interfaces

type Getaller[T any] interface {
	GetAll() []T
}
