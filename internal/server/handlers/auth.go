package handlers

import (
	"encoding/json"
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/pkg/models"
	"net/http"
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

	if !h.service.IsWhitelisted(user.Email) {
		http.Error(w, "access denied, not whitelisted", http.StatusForbidden)
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

	if !h.service.IsWhitelisted(creds.Email) {
		http.Error(w, "access denied, not whitelisted", http.StatusForbidden)
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
