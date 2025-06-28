package services

import (
	"mine-server-manager/internal/services/auth"
	"mine-server-manager/internal/services/docker"
	"mine-server-manager/internal/services/modpack"
	"mine-server-manager/internal/services/monitor"
	"mine-server-manager/internal/services/servers"
)

type Service struct {
	AuthService    *auth.AuthService
	DockerService  *docker.DockerService
	ModpackService *modpack.ModpackService
	MonitorService *monitor.MonitoringService
	ServerService  *servers.ServerManagementService
}
