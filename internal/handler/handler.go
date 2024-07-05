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

	// user
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/users", h.GetUsersOrderBy).Methods("GET")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/info", h.UserInfo).Methods("GET")

	// task
	r.HandleFunc("/tasks", h.StartTask).Methods("POST")
	r.HandleFunc("/tasks", h.FinishTask).Methods("PUT") // finishing task
	r.HandleFunc("/tasks", h.GetTasks).Methods("GET")

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
