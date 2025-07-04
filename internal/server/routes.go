package server

import (
	"encoding/json"
	"log"
	"mine-server-manager/internal/server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	authHandler := handlers.NewAuthHandler(s.services.AuthService)

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.LoginHandler).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", authHandler.LogoutHandler).Methods("POST", "OPTIONS")
	authRouter.Handle(
		"/register",
		s.services.AuthService.WhitelistMiddleware(
			http.HandlerFunc(authHandler.RegisterHandler))).Methods("POST", "OPTIONS")

	dockerHandler := handlers.NewDockerHandler(s.services.DockerService)
	r.HandleFunc("/create/{serverType}/{version}", dockerHandler.CreateServerHandler).Queries("version", "{version}").Methods("POST", "OPTIONS")

	r.HandleFunc("/", s.HelloWorldHandler)

	return r
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SecurityHeadersMiddlewware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
