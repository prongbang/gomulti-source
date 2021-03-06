// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package api

import (
	"github.com/prongbang/gomulti-source/pkg/api/post"
	"github.com/prongbang/gomulti-source/pkg/db"
)

// Injectors from wire.go:

func CreateAPI(dbSource db.DataSource) API {
	dataSource := post.NewDataSource(dbSource)
	repository := post.NewRepository(dataSource)
	useCase := post.NewUseCase(repository)
	handler := post.NewHandler(useCase)
	route := post.NewRoute(handler)
	apiAPI := NewAPI(route)
	return apiAPI
}
