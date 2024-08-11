package conceptDescription

import (
	"database/sql"

	"github.com/abhi2303237/AAS/backend/memmory"
	"github.com/abhi2303237/AAS/backend/mongodb"
	"github.com/abhi2303237/AAS/backend/mssql"
	"github.com/abhi2303237/AAS/backend/types"
	"github.com/abhi2303237/AAS/conceptDescription/config"
	"github.com/abhi2303237/AAS/utils"
	"github.com/golobby/container/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewConceptDescriptionCrudRepo() types.ICrudRepo[ConceptDescription] {
	var config config.Config
	utils.HandleFatal(container.Resolve(&config))
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

	case "mongodb":
		db = &mongodb.MongoDb{ConnectionString: config.MongoConnectionString}
		mongoCrudRep := &mongodb.MongoCrudRepo[ConceptDescription]{
			EntityName:   "ConceptDescription",
			DatabaseName: "AAS",
			Db:           db.GetDatabase().(*mongo.Client),
		}
		mongoCrudRep.Init()
		return mongoCrudRep

	default:
		db = &memmory.InMemmoryDB[ConceptDescription]{
			Data: make(map[string]ConceptDescription),
		}
		return &memmory.InMemmoryCrudRepo[ConceptDescription]{
			Db: db.GetDatabase().(map[string]ConceptDescription),
		}
	}
}
