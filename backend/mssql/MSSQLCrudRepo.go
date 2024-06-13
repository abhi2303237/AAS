package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/abhi2303237/AAS/utils"
	"golang.org/x/exp/maps"
)

type MSSQLCrudRepo[T any] struct {
	EntityName string
	Db         *sql.DB
	Data       map[string]T
}

func (sdb *MSSQLCrudRepo[T]) Get(ctx context.Context) ([]T, error) {
	return maps.Values(sdb.Data), nil
}

func (sdb *MSSQLCrudRepo[T]) GetById(ctx context.Context, id string) (*T, error) {
	if entity, ok := sdb.Data[id]; ok {
		return &entity, nil
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")
}

func (sdb *MSSQLCrudRepo[T]) Put(ctx context.Context, e *T, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		return nil, utils.GetError(http.StatusConflict, "Already Exist")
	}
	sdb.Data[id] = *e
	return e, nil
}

func (sdb *MSSQLCrudRepo[T]) Patch(ctx context.Context, e *T, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		sdb.Data[id] = *e
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *MSSQLCrudRepo[T]) Delete(ctx context.Context, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		delete(sdb.Data, id)
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *MSSQLCrudRepo[T]) Init() bool {
	var exists bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE 1=0", sdb.EntityName)
	err := sdb.Db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		//table doesn't exist create table
		tsql := fmt.Sprintf("create table %s ([documnt_id] varchar(max) primary key identity, [content] json)", sdb.EntityName)
		if _, err = sdb.Db.Exec(tsql); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return exists
}
