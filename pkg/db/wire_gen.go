// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package db

// Injectors from wire.go:

func NewDbConnection() DataSource {
	dbDataSource := NewDataSource()
	return dbDataSource
}
