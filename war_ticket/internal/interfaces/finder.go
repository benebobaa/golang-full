package interfaces

type Finder[T any] interface {
	FindByID(id int) (*T, error)
}
