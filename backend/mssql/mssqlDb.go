// Go connection Sample Code:
package mssql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/backend/types"
	"github.com/golobby/container/v3"
	_ "github.com/microsoft/go-mssqldb"
)

// var db *sql.DB
// var server = "iapps-in-dev-sql-server.database.windows.net"
// var port = 1433
// var user = "iappsadmin"
// var password = "<your_password>"
// var database = "getmaildb"

type MssqlDb struct {
	conn     *sql.DB
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

func (db *MssqlDb) Connect() bool {
	// Build connection string
	var conn *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		db.Server, db.User, db.Password, db.Port, db.Database)
	var err error
	// Create connection pool
	conn, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return false
	}
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	db.conn = conn
	fmt.Printf("Database Connected!")
	return true
}

func (db *MssqlDb) GetDatabase() any {
	if db.conn == nil {
		db.Connect()
	}
	return db.conn
}

func NewMssqlDatabase() *types.IDatabase {
	var mssqlDatabase types.IDatabase
	if err := container.Resolve(&mssqlDatabase); err != nil {
		log.Fatal(err)
	}
	return &mssqlDatabase
}
