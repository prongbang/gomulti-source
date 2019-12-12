package api

import "github.com/prongbang/gomulti-source/pkg/api/post"

type API interface {
	Register()
}

type api struct {
	PostRoute post.Route
}

func (a *api) Register() {
	a.PostRoute.Initial()
}

func NewAPI(
	postRoute post.Route,
) API {
	return &api{
		PostRoute: postRoute,
	}
}
