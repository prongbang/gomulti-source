package post

import (
	"net/http"
)

type Route interface {
	Initial()
}

type route struct {
	Handle Handler
}

func (r *route) Initial() {
	http.HandleFunc("/", r.Handle.HelloWorld)
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}
