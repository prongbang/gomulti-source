package db

import (
	"context"
	"fmt"
	"github.com/prongbang/gomulti-source/pkg/db/driver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoDB struct {
	Cfg Config
}

func (m *mongoDB) Connect() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", m.Cfg.Host)))
	if err != nil {
		log.Fatalf("MongoDb Connection fail : %s", err)
	}
	log.Print("MongoDb Connected.")
	return client.Database(m.Cfg.DatabaseName)
}

func NewMongoDB() driver.MongoDriver {
	return &mongoDB{
		Cfg: Config{
			User:         "root",
			Pass:         "pass",
			Host:         "localhost",
			DatabaseName: "nameDB",
		},
	}
}
