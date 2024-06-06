package backend

import "github.com/abhi2303237/AAS/backend/types"

// import "github.com/abhi2303237/AAS/backend/types"

type ICrudRepo[T types.IEntity] interface {
	Get() ([]T, error)
	GetById(id string) (*T, error)
	Put(e *T, id string) (*T, error)
	Patch(e *T, id string) (*T, error)
	Delete(id string) (*T, error)
}
