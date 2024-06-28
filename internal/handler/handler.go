package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/OurLuv/time-tracker/internal/service"
	"github.com/gorilla/mux"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data,omitempty"`
}

type Handler struct {
	service *service.Service
	log     *slog.Logger
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user", h.CreateUser).Methods("POST")

	return r
}

func sendError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	resp := Response{
		Msg: msg,
	}

	json.NewEncoder(w).Encode(&resp)
}

func NewHandler(service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
