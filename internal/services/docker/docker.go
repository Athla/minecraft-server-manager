package docker

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

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

func (s *Service) getAvailablePort() (string, error) {
	for i := MIN_PORT; i < MAX_PORT; i++ {
		currPort := strconv.Itoa(i)
		address := net.JoinHostPort("localhost", currPort)
		conn, err := net.DialTimeout("udp", address, time.Millisecond*500)
		if err != nil {
			continue
		}

		defer conn.Close()
		return "", nil
	}

	return "", fmt.Errorf("Server side error, all ports in use.")
}

func (s *Service) CreateServer(ctx context.Context) (string, error) {
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

	serverName := fmt.Sprintf("minecraft-server-%s", port)

	containerConfig := &container.Config{
		Image: imageName,
		Cmd:   opts.Commands,
		ExposedPorts: nat.PortSet{
			nat.Port("25565/tcp"): struct{}{},
		},
		Env: []string{
			"EULA=TRUE",
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
