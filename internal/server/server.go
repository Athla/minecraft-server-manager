package server

import (
	"fmt"
	"log"
	"mine-server-manager/internal/config"
	"mine-server-manager/internal/repository"
	"mine-server-manager/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port     int
	cfg      *config.Config
	services *services.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db := repository.NewRepository(cfg)
	services := services.NewServiceWrapper(cfg, db)
	NewServer := &Server{
		port:     port,
		cfg:      cfg,
		services: services,
	}

	router := NewServer.RegisterRoutes()
	handler := handlers.LoggingHandler(os.Stdout, router)

	// Declare Server config

	log.Printf("Server started on port :%d\n", NewServer.port)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
