package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Service struct {
}

// Service to create the server

type CreateOptions struct {
	Name     string   `json:"name"`
	Registry string   `json:"registry"`
	Image    string   `json:"image"`
	Version  string   `json:"version"`
	Commands []string `json:"commands"`
}

const MIN_PORT = 25565
const MAX_PORT = 25595

func NewDockerService() *Service {
	return &Service{}
}

func (s *Service) CreateServer(ctx context.Context, serverType string) (string, error) {
	opts := new(CreateOptions)
	parseCreateOpts(opts)

	cli, err := s.newDockerClient(client.FromEnv)
	if err != nil {
		return "", err
	}

	defer cli.Close()
	imageName := fmt.Sprintf("%s/%s:%s", opts.Registry, opts.Image, opts.Version)
	reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return "", err
	}

	defer reader.Close()
	io.Copy(io.Discard, reader)

	port, err := s.getAvailablePort()
	if err != nil {
		return "", err
	}

	var chosenServerEnv string
	switch strings.ToLower(serverType) {
	case "forge":
		chosenServerEnv = "FORGE"
	case "fabric":
		chosenServerEnv = "FABRIC"
	case "paper":
		chosenServerEnv = "PAPER"
	case "neoforge":
		chosenServerEnv = "NEOFORGE"
	case "vanilla":
		chosenServerEnv = "VANILLA"
	default:
		chosenServerEnv = "VANILLA"
	}

	serverName := fmt.Sprintf("minecraft-server-%s-%v", chosenServerEnv, port)
	log.Infof("SERVICE-DOCKER: Server name: %s", serverName)
	serverType = fmt.Sprintf("TYPE=%s", chosenServerEnv)

	containerConfig := &container.Config{
		Image: imageName,
		Cmd:   opts.Commands,
		ExposedPorts: nat.PortSet{
			nat.Port("25565/tcp"): struct{}{},
		},
		Env: []string{
			"EULA=TRUE",
			serverType,
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"25565/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, serverName)
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	log.Infof("SERVICE-DOCKER: Container started: %s", resp.ID)

	return resp.ID, nil
}

// This should be called only by 2 users
func (s *Service) DeleteServer(ctx context.Context, serverId string) error {
	cli, err := s.newDockerClient(client.FromEnv)
	if err != nil {
		return err
	}

	defer cli.Close()
	rmvOpts := container.RemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}

	if err := cli.ContainerRemove(ctx, serverId, rmvOpts); err != nil {
		return err
	}

	return nil
}

func (s *Service) StopServer(ctx context.Context, serverId int64) error {

	return nil
}
