package handlers

import (
	"encoding/json"
	"mine-server-manager/internal/internalErrors"
	"mine-server-manager/internal/services/docker"
	"mine-server-manager/pkg/models"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type DockerHandler struct {
	service *docker.Service
}

type ServerCreateRequest struct {
	ServerName string `json:"serverName"`
}

func NewDockerHandler(service *docker.Service) *DockerHandler {
	return &DockerHandler{
		service: service,
	}
}

type MinecraftServerType string

const (
	Forge    MinecraftServerType = "forge"
	Fabric   MinecraftServerType = "fabric"
	Paper    MinecraftServerType = "paper"
	Vanilla  MinecraftServerType = "vanilla"
	Java     MinecraftServerType = "java"
	NeoForge MinecraftServerType = "neoforge"
)

var validServerTypes = []MinecraftServerType{
	Forge,
	Fabric,
	Paper,
	Vanilla,
	Java,
	NeoForge,
}

const MIN_VERSION = 1_17_000
const LATEST_VERSION_SUPPORTED = 1_21_100

func (h *DockerHandler) validateVars(vars map[string]string) error {
	if vars["serverType"] == "" {
		return internalErrors.ErrServerTypeNeeded
	}

	serverType := MinecraftServerType(vars["serverType"])
	if slices.Contains(validServerTypes, serverType) {
		return internalErrors.ErrInvalidServerType
	}

	rawVersion := vars["version"]
	// change to int, check if between valid versions
	if rawVersion == "" {
		rawVersion = ""
	}

	version, err := strconv.Atoi(strings.Join(strings.Split(rawVersion, "."), ""))
	if err != nil {
		return err
	}

	if version < MIN_VERSION || version > LATEST_VERSION_SUPPORTED {
		return internalErrors.ErrInvalidVersion
	}

	return nil
}

func (h *DockerHandler) CreateServerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := h.validateVars(vars); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp, err := h.service.CreateServer(r.Context(), vars["serverType"], vars["version"])
	if err != nil {
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
