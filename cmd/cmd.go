package cmd

import (
	"github.com/prongbang/gomulti-source/pkg/api"
	"github.com/prongbang/gomulti-source/pkg/db"
)

func Run() {
	conn := db.NewDbConnection()
	apis := api.CreateAPI(conn)
	apis.Register()
}
