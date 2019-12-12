package driver

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDriver is the interface
type MongoDriver interface {
	Connect() *mongo.Client
}
