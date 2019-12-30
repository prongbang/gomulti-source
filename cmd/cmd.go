package cmd

import (
	"github.com/prongbang/gomulti-source/pkg/api"
	"github.com/prongbang/gomulti-source/pkg/db"
)

func Run() {
	api.CreateAPI(db.NewDbConnection()).Register()
}
