package auth

import (
	"encoding/json"
	"net/http"
)

func (s *AuthService) WhitelistMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Email string `json:"email"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if !s.IsWhitelisted(body.Email) {
			http.Error(w, "email not whitelisted", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
