package db

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSource interface {
	GetMariaDB() *sqlx.DB
	GetMongoDB() *mongo.Client
}

type dataSource struct {
	MongoDB *mongo.Client
	MariaDB *sqlx.DB
}

func (d *dataSource) GetMariaDB() *sqlx.DB {
	return d.MariaDB
}

func (d *dataSource) GetMongoDB() *mongo.Client {
	return d.MongoDB
}

func NewDataSource() DataSource {
	return &dataSource{
		MongoDB: NewMongoDB().Connect(),
		MariaDB: NewMariaDB().Connect(),
	}
}
