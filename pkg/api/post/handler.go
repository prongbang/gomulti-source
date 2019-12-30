package post

import (
	"fmt"
	"net/http"
	"time"
)

type Handler interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Uc UseCase
}

func (h *handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}
