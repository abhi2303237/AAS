package mongodb

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/structs"
	"github.com/gofiber/fiber/v2/log"

	"github.com/abhi2303237/AAS/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/maps"
)

type MongoCrudRepo[T any] struct {
	EntityName   string
	DatabaseName string
	Db           *mongo.Client
	Data         map[string]T
}

func (sdb *MongoCrudRepo[T]) Get(ctx context.Context) ([]T, error) {
	return maps.Values(sdb.Data), nil
}

func (sdb *MongoCrudRepo[T]) GetById(ctx context.Context, id string) (*T, error) {

	collection := sdb.Db.Database(sdb.DatabaseName).Collection(sdb.EntityName)

	// Create a filter to check if the document already exists
	var result T
	filter := bson.M{"_id": id}
	collection.FindOne(ctx, filter).Decode(&result)
	return &result, nil
}

func (sdb *MongoCrudRepo[T]) Put(ctx context.Context, e *T, id string) (*T, error) {

	collection := sdb.Db.Database(sdb.DatabaseName).Collection(sdb.EntityName)
	document, err := structs.ToMap(e)
	if err != nil {
		return nil, err
	}

	// Set the document ID
	document["_id"] = id

	// Insert the document
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	_, err = collection.InsertOne(ctx, document)
	if err != nil {
		log.Errorf("failed to insert document: %v", err)
		return nil, err
	}
	return e, nil
}

func (sdb *MongoCrudRepo[T]) Patch(ctx context.Context, e *T, id string) (*T, error) {
	// Access the collection
	collection := sdb.Db.Database(sdb.DatabaseName).Collection(sdb.EntityName)

	// Create a filter to check if the document already exists
	filter := bson.M{"_id": id}

	// Upsert the record, meaning it will be inserted if it doesn't exist, or updated if it does
	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(ctx, filter, e, opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Record with ID '%s' saved successfully", id)
	return e, nil
}

func (sdb *MongoCrudRepo[T]) Delete(ctx context.Context, id string) (*T, error) {
	if _, ok := sdb.Data[id]; ok {
		delete(sdb.Data, id)
	}
	return nil, utils.GetError(http.StatusNotFound, "Not Found")

}

func (sdb *MongoCrudRepo[T]) Init() bool {
	var exists bool
	// Checking if the collection exists
	database := sdb.Db.Database(sdb.DatabaseName)
	collections, err := database.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Info(err)
		return createDatabase(sdb.DatabaseName, sdb.EntityName, database)
	}

	// Iterating over existing collections to check if the target collection exists
	for _, coll := range collections {
		if coll == sdb.EntityName {
			log.Infof("Collection '%s' already exists in database '%s'", sdb.EntityName, sdb.DatabaseName)
			return true
		} else {
			// If the collection doesn't exist, create it
			return createDatabase(sdb.DatabaseName, sdb.EntityName, database)
		}
	}
	return exists
}

func createDatabase(databaseName string, collectionName string, database *mongo.Database) bool {
	log.Infof("Collection '%s' does not exist in database '%s'. Creating...", collectionName, databaseName)
	err := database.CreateCollection(context.Background(), collectionName)
	if err != nil {
		log.Fatal("Failed to create collection ", err)
	}
	log.Infof("Collection '%s' created successfully in database '%s'", collectionName, databaseName)
	return true
}
