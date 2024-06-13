package conceptDescription

import (
	"database/sql"

	"github.com/abhi2303237/AAS/backend/memmory"
	"github.com/abhi2303237/AAS/backend/mssql"
	"github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/utils"
)

func NewConceptDescriptionCrudRepo() types.ICrudRepo[ConceptDescription] {
	config := utils.GetConfig()
	var db types.IDatabase
	switch config.Backend {
	case "mssql":
		db = &mssql.MssqlDb{
			Server:   config.DBServer,
			Port:     config.DBPort,
			User:     config.DBUser,
			Password: config.DBPassword,
			Database: config.DBDatabase,
		}
		mssqlCrudRepo := &mssql.MSSQLCrudRepo[ConceptDescription]{
			EntityName: "ConceptDescription",
			Db:         db.GetDatabase().(*sql.DB),
		}
		mssqlCrudRepo.Init()
		return mssqlCrudRepo
	default:
		db = &memmory.InMemmoryDB[ConceptDescription]{
			Data: make(map[string]ConceptDescription),
		}
		return &memmory.InMemmoryCrudRepo[ConceptDescription]{
			Db: db.GetDatabase().(map[string]ConceptDescription),
		}
	}
}
