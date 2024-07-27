package interfaces

type Updater[T any] interface {
	Update(value *T) (*T, error)
}
