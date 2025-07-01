package handlers

import (
	"encoding/json"
	"mine-server-manager/internal/services/docker"
	"mine-server-manager/pkg/models"
	"net/http"

	"github.com/charmbracelet/log"
)

type DockerHandler struct {
	service *docker.Service
}

func NewDockerHandler(service *docker.Service) *DockerHandler {
	return &DockerHandler{
		service: service,
	}
}

func (h *DockerHandler) CreateServerHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.CreateServer(r.Context(), r)
	if err != nil {
		log.Errorf("Failed to create server due: %s", err)
		http.Error(w, "failed to create server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "server created successfully",
		Data:    resp,
		Code:    http.StatusCreated,
	})
}
