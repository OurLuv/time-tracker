package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/gorilla/mux"
)

// * Creating user
// @Summary Create new user
// @Description Creates a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.User true "User data"
// @Success 201 {object} Response
// @Failure 400 {object} Response "Invalid data"
// @Failure 500 {object} Response "Internal server error"
// @Router /users [post]
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

// * Getting all users ordered by param
// @Summary Get all users ordered by param
// @Description Get all users, ordered by "order" param
// @Tags users
// @Produce  json
// @Param order query string false "Order by field (passport_number, name, surname, patronymic, address)"
// @Success 200 {object} Response
// @Failure 500 {object} Response "Internal server error"
// @Router /users [get]
func (h *Handler) GetUsersOrderBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	h.log.Debug("getting all users")

	// getting param
	ordered := r.URL.Query().Get("order")

	// getting users
	users, err := h.service.UserService.GetOrderBy(context.Background(), ordered)
	if err != nil {
		h.log.Error("can't get users", slog.String("err", err.Error()))
		sendError(w, "can't get users", http.StatusInternalServerError)
		return
	}

	// sending response
	msg := fmt.Sprintf("got all users ordered by %s", ordered)
	resp := Response{
		Msg:  msg,
		Data: users,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Deleting user
// @Summary Delete user
// @Description Deletes a user by ID
// @Tags users
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 {object} Response
// @Failure 400 {object} Response "Invalid ID"
// @Failure 500 {object} Response "Internal server error"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("can't validate param", slog.String("id", idStr))
		sendError(w, "can't validate param", http.StatusBadRequest)
		return
	}

	if err := h.service.UserService.DeleteUser(context.Background(), id); err != nil {
		h.log.Error("can't delete user", slog.String("err", err.Error()))
		sendError(w, "can't delete users", http.StatusInternalServerError)
		return
	}
}

// * Updating user
// @Summary Update user
// @Description Updates a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body model.User true "User data"
// @Success 204 {object} Response
// @Failure 400 {object} Response "Invalid data"
// @Failure 500 {object} Response "Internal server error"
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// getting data
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("can't validate param", slog.String("id", idStr))
		sendError(w, "can't validate param", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.log.Error("can't decode data", slog.String("err", err.Error()))
		sendError(w, "Data is invalid", http.StatusBadRequest)
	}
	user.Id = id

	// updating the user
	if err := h.service.UserService.UpdateUser(context.Background(), user); err != nil {
		h.log.Error("can't delete user", slog.String("err", err.Error()))
		sendError(w, "can't delete users", http.StatusInternalServerError)
		return
	}
}

// * Getting user by passport
// @Summary Get user by passport number
// @Description Get user by passport number
// @Tags users
// @Produce  json
// @Param passportSerie path string true "Passport serie"
// @Param passportNumber path string true "Passport number"
// @Success 200 {object} model.User
// @Failure 400 {object} Response "Invalid data"
// @Failure 500 {object} Response "Internal server error"
// @Router /users/{passportSerie}/{passportNumber} [get]
func (h *Handler) UserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	passportSerie := mux.Vars(r)["passportSerie"]
	passportNumber := mux.Vars(r)["passportNumber"]

	pass := fmt.Sprintf("%s %s", passportSerie, passportNumber)

	user, err := h.service.UserService.GetUserByPassport(context.Background(), pass)
	if err != nil {
		h.log.Error("can't get user", slog.String("err", err.Error()))
		sendError(w, "can't get users", http.StatusInternalServerError)
		return
	}

	// sending response
	json.NewEncoder(w).Encode(&user)
}
