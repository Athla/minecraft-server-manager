package docker

import "github.com/docker/docker/client"

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
