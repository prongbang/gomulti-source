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

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func (r *route) Initial() {

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}
