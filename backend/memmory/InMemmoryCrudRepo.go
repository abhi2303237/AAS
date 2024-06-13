package memmory

import (
	"context"
	"net/http"

	"github.com/abhi2303237/AAS/utils"
	"golang.org/x/exp/maps"
)

type InMemmoryCrudRepo[T any] struct {
	Db map[string]T
}

func (sdb *InMemmoryCrudRepo[T]) Get(ctx context.Context) ([]T, error) {
	return maps.Values(sdb.Db), nil
}

func (sdb *InMemmoryCrudRepo[T]) GetById(ctx context.Context, id string) (*T, error) {
	if entity, ok := sdb.Db[id]; ok {
		return &entity, nil
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")
}

func (sdb *InMemmoryCrudRepo[T]) Put(ctx context.Context, e *T, id string) (*T, error) {
	if _, ok := sdb.Db[id]; ok {
		return nil, utils.GetError(http.StatusConflict, "Already Exist")
	}
	sdb.Db[id] = *e
	return e, nil
}

func (sdb *InMemmoryCrudRepo[T]) Patch(ctx context.Context, e *T, id string) (*T, error) {
	if _, ok := sdb.Db[id]; ok {
		sdb.Db[id] = *e
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *InMemmoryCrudRepo[T]) Delete(ctx context.Context, id string) (*T, error) {
	if _, ok := sdb.Db[id]; ok {
		delete(sdb.Db, id)
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *InMemmoryCrudRepo[T]) Init() bool {
	return true
}
