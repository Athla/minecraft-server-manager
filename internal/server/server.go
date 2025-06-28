package server

import (
	"fmt"
	"mine-server-manager/internal/config"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	cfg  *config.Config
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	NewServer := &Server{
		port: port,
		cfg:  cfg,
	}

	router := NewServer.RegisterRoutes()
	handler := handlers.LoggingHandler(os.Stdout, router)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
