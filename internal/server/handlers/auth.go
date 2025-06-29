package handlers

import (
	"context"
	"encoding/json"
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/pkg/models"
	"net/http"
)

type AuthHandler struct {
	service *auth.AuthService
	ctx     context.Context
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pwd := r.FormValue("password")

	if !h.service.IsWhitelisted(email) {
		http.Error(w, "Access denied", http.StatusForbidden)

		return
	}

	if !h.service.ValidatePwd(h.ctx, email, pwd) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Message: "Success",
		Data:    "Login successful",
		Code:    http.StatusOK,
	})
}
