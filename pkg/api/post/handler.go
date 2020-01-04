package post

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler interface {
	GetPostList(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Uc UseCase
}

func (h *handler) GetPostList(w http.ResponseWriter, r *http.Request) {
	list := h.Uc.GetPostList()
	jsonList, _ := json.Marshal(list)
	_, _ = fmt.Fprintf(w, "%s", fmt.Sprintf("%s", jsonList))
}

func (h *handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id := "1"
	obj := h.Uc.GetPostById(id)
	jsonObj, _ := json.Marshal(obj)
	_, _ = fmt.Fprintf(w, "%s", fmt.Sprintf("%s", jsonObj))
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}
