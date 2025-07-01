package docker

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Service struct {
}

// Service to create the server
// Allocate the ports

type CreateOptions struct {
	Name     string   `json:"name"`
	Registry string   `json:"registry"`
	Image    string   `json:"image"`
	Version  string   `json:"version"`
	Commands []string `json:"commands"`
}

func NewDockerService() *Service {
	return &Service{}
}

func (s *Service) CreateServer(ctx context.Context, r *http.Request) (string, error) {
	opts := new(CreateOptions)
	parseCreateOpts(opts)

	cli, err := newDockerClient(client.FromEnv)
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
	config := &container.Config{
		Image: imageName,
		Cmd:   opts.Commands,
		ExposedPorts: nat.PortSet{
			nat.Port("25565/tcp"): struct{}{},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, nil, nil, nil, opts.Name)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (s *Service) DeleteServer(ctx context.Context, serverId int64) error {

	return nil
}

func (s *Service) StopServer(ctx context.Context, serverId int64) error {

	return nil
}

// func (s *ContainerHandler) CreateContainerHandler(e echo.Context) error {
//
// 	resp, err := createDockerContainer(context.Background(), cli, reader, opts, imageName)
// 	if err != nil {
// 		e.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "internal server error.",
// 		})
//
// 		return err
// 	}
//
// 	log.Infof("CONTAINER: Container ID: %s", resp.ID)
// 	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
// 		log.Warnf("CONTAINER: Unable to start container due: %s", err)
// 		e.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "internal server error.",
// 		})
// 		return err
// 	}
//
// 	log.Info("CONTAINER: Container created sucessfully!")
// 	e.JSON(http.StatusCreated, map[string]string{
// 		"success": fmt.Sprintf("created container with ID: %s successfully!", resp.ID),
// 	})
//
// 	return nil
// }
