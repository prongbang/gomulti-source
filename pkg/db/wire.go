//+build wireinject

package db

import "github.com/google/wire"

func NewDbConnection() DataSource {
	wire.Build(NewDataSource, NewMariaDB, NewMongoDB)
	return nil
}
