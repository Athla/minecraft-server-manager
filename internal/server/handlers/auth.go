package handlers

import (
	"mine-server-manager/internal/services/auth"
	"net/http"
)

type AuthHandler struct {
	service *auth.AuthService
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pwd := r.FormValue("password")

	if !h.service.IsWhitelisted(email) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	if !h.service.ValidatePwd()

}
