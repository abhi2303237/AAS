// Go connection Sample Code:
package memmory

import (
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

type InMemmoryDB[T any] struct {
	Data map[string]T
}

func (db *InMemmoryDB[T]) Connect() bool {

	fmt.Printf("Database Connected!")
	return true
}

func (db *InMemmoryDB[T]) GetDatabase() any {
	return db.Data
}

// func NewMssqlDatabase() *backend.IDatabase {
// 	var mssqlDatabase backend.IDatabase
// 	if err := container.Resolve(&mssqlDatabase); err != nil {
// 		log.Fatal(err)
// 	}
// 	return &mssqlDatabase
// }
