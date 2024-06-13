package types

import (
	"context"
)

type ICrudRepo[T any] interface {
	Get(ctx context.Context) ([]T, error)
	GetById(ctx context.Context, id string) (*T, error)
	Put(ctx context.Context, e *T, id string) (*T, error)
	Patch(ctx context.Context, e *T, id string) (*T, error)
	Delete(ctx context.Context, id string) (*T, error)
	Init() bool
}
