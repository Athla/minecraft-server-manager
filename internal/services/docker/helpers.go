package docker

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/docker/docker/client"
)

func parseCreateOpts(opts *CreateOptions) {
	if opts.Registry == "" {
		opts.Registry = "docker.io"
	}

	if opts.Version == "" {
		opts.Version = "latest"
	}

	if opts.Image == "" {
		opts.Image = "itzg/minecraft-server"
	}
}

func (s *Service) newDockerClient(opts client.Opt) (*client.Client, error) {
	cli, err := client.NewClientWithOpts(opts)
	if err != nil {
		return nil, err
	}

	return cli, nil
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
		port := strconv.Itoa(conn.LocalAddr().(*net.UDPAddr).Port)

		return port, nil
	}

	return "", fmt.Errorf("Server side error, all ports in use.")
}
