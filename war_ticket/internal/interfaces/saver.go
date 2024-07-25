package interfaces

type Saver[T any] interface {
	Save(value *T) (*T, error)
}
