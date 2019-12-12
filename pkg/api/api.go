package api

import (
	"github.com/prongbang/gomulti-source/pkg/api/post"
	"net/http"
)

type API interface {
	Register()
}

type api struct {
	PostRoute post.Route
}

func (a *api) Register() {
	a.PostRoute.Initial()

	_ = http.ListenAndServe(":8080", nil)
}

func NewAPI(
	postRoute post.Route,
) API {
	return &api{
		PostRoute: postRoute,
	}
}
