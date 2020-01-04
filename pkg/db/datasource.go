package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/prongbang/gomulti-source/pkg/db/driver"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSource interface {
	GetMariaDB() *sqlx.DB
	GetMongoDB() *mongo.Database
}

type dataSource struct {
	MongoDB *mongo.Database
	MariaDB *sqlx.DB
}

func (d *dataSource) GetMariaDB() *sqlx.DB {
	return d.MariaDB
}

func (d *dataSource) GetMongoDB() *mongo.Database {
	return d.MongoDB
}

func NewDataSource(mongoDB driver.MongoDriver, mariaDB driver.SQLxDriver) DataSource {
	return &dataSource{
		MongoDB: mongoDB.Connect(),
		MariaDB: mariaDB.Connect(),
	}
}
