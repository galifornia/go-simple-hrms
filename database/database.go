package database

import (
	"context"
	"time"

	"github.com/galifornia/go-simple-hrms/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "simple-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

var DB types.MongoInstance

func OpenDB() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db := client.Database(dbName)

	DB.Client = client
	DB.Database = db

	return nil
}
