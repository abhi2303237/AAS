package mongodb

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/backend/types"
	"github.com/golobby/container/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	conn             *mongo.Client
	ConnectionString string
}

func (db *MongoDb) Connect() bool {
	clientOptions := options.Client().ApplyURI(db.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db.conn = client
	fmt.Println("Connected to MongoDB!")
	return true
}

func (db *MongoDb) GetDatabase() any {
	if db.conn == nil {
		db.Connect()
	}
	err := db.conn.Ping(context.TODO(), nil)
	if err != nil {
		db.Connect()
	}
	return db.conn
}

func NewMongoDatabase() *types.IDatabase {
	var mongoDatabase types.IDatabase
	if err := container.Resolve(&mongoDatabase); err != nil {
		log.Fatal(err)
	}
	return &mongoDatabase
}
