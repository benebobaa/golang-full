package interfaces

type Getter[T any] interface {
	GetAll() []T
}
