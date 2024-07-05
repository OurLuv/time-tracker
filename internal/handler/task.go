package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/OurLuv/time-tracker/internal/model"
)

// * Starting task
func (h *Handler) StartTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var task model.Task

	// getting data
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.log.Error("error parsing task json", slog.String("err", err.Error()))
		sendError(w, "data is not valid", http.StatusBadRequest)
		return
	}

	// validating
	if task.Title == "" || task.UserId == 0 {
		h.log.Error("error validating task")
		sendError(w, "task.title and tasl.user_id are required", http.StatusBadRequest)
		return
	}

	// starting task
	id, err := h.service.TaskService.StartTask(context.Background(), task)
	if err != nil {
		h.log.Error("error from service. Can't start a task", slog.String("err", err.Error()))
		sendError(w, "server error", http.StatusInternalServerError)
		return
	}

	// sending response
	msg := fmt.Sprintf("Task is started. Task's id = [%d]", id)
	resp := Response{
		Msg: msg,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Finishing task
func (h *Handler) FinishTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var data struct {
		Id int `json:"id"`
	}

	// getting data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error("error parsing json", slog.String("err", err.Error()))
		sendError(w, "data is not valid", http.StatusBadRequest)
		return
	}

	// validating
	if data.Id == 0 {
		h.log.Error("error validating task")
		sendError(w, "id is required and it can't be 0", http.StatusBadRequest)
		return
	}

	// finishing task
	dur, err := h.service.TaskService.FinishTask(context.Background(), data.Id)
	if err != nil {
		h.log.Error("error from service. Can't finish a task", slog.String("err", err.Error()))
		sendError(w, "server error", http.StatusInternalServerError)
		return
	}

	// sending response
	msg := fmt.Sprintf("Task[id][%d] is finished. Task's duration = [%s]", data.Id, dur)
	resp := Response{
		Msg: msg,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Getting tasks of user for certain period
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var data struct {
		From   string `json:"from"`
		To     string `json:"to"`
		UserId int    `json:"user_id"`
	}

	// getting data
	var period model.TaskPeriod
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error("error parsing json", slog.String("err", err.Error()))
		sendError(w, "data is not valid", http.StatusBadRequest)
		return
	}
	if err := period.SetTime(data.From, data.To); err != nil {
		h.log.Error("error with time parsing", slog.String("err", err.Error()))
		sendError(w, "data is not valid", http.StatusBadRequest)
		return
	}
	period.UserId = data.UserId

	// validating
	if period.From.After(period.To) || period.From.IsZero() || period.To.IsZero() || period.UserId == 0 {
		h.log.Error("error validating task", slog.Any("TaskPeriod", period))
		sendError(w, "error validating task", http.StatusBadRequest)
		return
	}

	// getting tasks
	tasks, err := h.service.TaskService.GetTasks(context.Background(), period)
	if err != nil {
		h.log.Error("can't get tasks", slog.String("err", err.Error()))
		sendError(w, "server error", http.StatusInternalServerError)
		return
	}

	// sending response
	msg := fmt.Sprintf("Got all tasks for user[id][%d] for period [%s] - [%s]", period.UserId, period.From, period.To)
	resp := Response{
		Msg:  msg,
		Data: tasks,
	}
	json.NewEncoder(w).Encode(resp)
}
