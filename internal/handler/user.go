package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/OurLuv/time-tracker/internal/model"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	h.log.Debug("creating user")
	var user model.User
	// getting data
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.log.Error("can't decode data", slog.String("err", err.Error()))
		sendError(w, "Data is invalid", http.StatusBadRequest)
	}

	// validating data
	if user.PassportNumber == "" {
		h.log.Error("can't validate user")
		sendError(w, "Passport number is required", http.StatusBadRequest)
		return
	}

	// creating user
	id, err := h.service.UserService.Create(context.Background(), user)
	if err != nil {
		h.log.Error("can't create user", slog.String("err", err.Error()))
		sendError(w, "can't create user", http.StatusInternalServerError)
		return
	}
	user.Id = id
	h.log.Info("user created", slog.Int("id", user.Id))

	// sending response
	resp := Response{
		Msg:  "user is created",
		Data: user,
	}
	json.NewEncoder(w).Encode(&resp)
}
