package services

import (
	"log/slog"
	"mine-server-manager/internal/config"
	"mine-server-manager/internal/repository"
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/internal/services/docker"
	"mine-server-manager/internal/services/modpack"
	"mine-server-manager/internal/services/monitor"
	"mine-server-manager/internal/services/servers"
)

type Service struct {
	AuthService    *auth.AuthService
	DockerService  *docker.Service
	ModpackService *modpack.ModpackService
	MonitorService *monitor.MonitoringService
	ServerService  *servers.ServerManagementService
}

func NewServiceWrapper(cfg *config.Config, db *repository.Repository) *Service {
	return &Service{
		AuthService: auth.NewAuthService(cfg.AuthConfig, slog.Default(), db),
	}
}
