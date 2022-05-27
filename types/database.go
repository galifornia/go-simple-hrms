package types

import "go.mongodb.org/mongo-driver/mongo"

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}
