//+build wireinject

package post

import (
	"github.com/google/wire"
	"github.com/prongbang/gomulti-source/pkg/db"
)

func NewPostAPI(dbSource db.DataSource) Route {
	wire.Build(
		NewRoute,
		NewHandler,
		NewUseCase,
		NewRepository,
		NewDataSource,
	)
	return nil
}
