package handlers

import (
	"encoding/json"
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/pkg/models"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "user crearted successfully",
		Data:    createdUser,
		Code:    http.StatusCreated,
	})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "Login successful",
		Data:    map[string]string{"token": token},
		Code:    http.StatusOK,
	})
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "missing authorization header", http.StatusBadRequest)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	err := h.service.Logout(tokenString)
	if err != nil {
		log.Errorf("Failed to logout due: %v", err)
		http.Error(w, "failed to logout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "logout successful, token invalidated",
		Code:    http.StatusOK,
	})
}
