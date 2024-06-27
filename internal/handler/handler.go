package handler

import "github.com/gorilla/mux"

type Handler struct{}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	return r
}
