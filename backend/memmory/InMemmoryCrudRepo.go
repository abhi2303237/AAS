package memmory

import (
	// "github.com/abhi2303237/AAS/backend"
	"net/http"

	// "github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/utils"
	"golang.org/x/exp/maps"
)

type InMemmoryCrudRepo[T types.IEntity] struct {
	Data map[string]T
}

func (sdb *InMemmoryCrudRepo[T]) Get() ([]T, error) {
	return maps.Values(sdb.Data), nil
}

func (sdb *InMemmoryCrudRepo[T]) GetById(id string) (*T, error) {
	if entity, ok := sdb.Data[id]; ok {
		return &entity, nil
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")
}

func (sdb *InMemmoryCrudRepo[T]) Put(e *T, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		return nil, utils.GetError(http.StatusConflict, "Already Exist")
	}
	sdb.Data[id] = *e
	return e, nil
}

func (sdb *InMemmoryCrudRepo[T]) Patch(e *T, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		sdb.Data[id] = *e
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *InMemmoryCrudRepo[T]) Delete(id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		delete(sdb.Data, id)
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}
