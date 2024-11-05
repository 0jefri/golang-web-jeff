package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-web/common"
	"github.com/golang-web/model"
	"github.com/golang-web/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) RegisterHandlerTask(w http.ResponseWriter, r *http.Request) {
	var payload model.Task

	payload.ID = common.GenerateUUID()

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.taskService.RegisterNewTask(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Task registered successfully",
		Data:       payload,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to retrieve taks", http.StatusInternalServerError)
		log.Printf("error retrieving tasks: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		log.Printf("error encoding tasks: %v", err)
	}
}
