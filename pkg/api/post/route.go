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
	http.HandleFunc("/v1/posts", r.Handle.GetPostList)
	http.HandleFunc("/v1/post", r.Handle.GetPost)
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}
