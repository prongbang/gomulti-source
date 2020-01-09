//+build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/prongbang/gomulti-source/pkg/api/post"
	"github.com/prongbang/gomulti-source/pkg/db"
)

func CreateAPI(dbSource db.DataSource) API {
	wire.Build(
		NewAPI,
		post.ProviderSet,
	)
	return nil
}
