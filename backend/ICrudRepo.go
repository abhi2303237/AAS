package backend

type ICrudRepo[T any] interface {
	Get() ([]T, error)
	GetById(id string) (*T, error)
	Put(e *T, id string) (*T, error)
	Patch(e *T, id string) (*T, error)
	Delete(id string) (*T, error)
}
