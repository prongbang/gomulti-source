package post

import (
	"fmt"
	"net/http"
	"time"
)

type Route interface {
	Initial()
}

type route struct {
	Handle Handler
}

func (r *route) Initial() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World! %s", time.Now())
	})
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}
