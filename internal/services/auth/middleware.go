package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (s *AuthService) WhitelistMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		var body struct {
			Email string `json:"email"`
		}

		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if !s.IsWhitelisted(body.Email) {
			http.Error(w, "email not whitelisted", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
