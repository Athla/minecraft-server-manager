package docker

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	svc    = &Service{}
	ctx    = context.Background()
	cancel = func() {}
)

var (
	server = flag.String("server", "vanilla", "minecraft server type")
)

func TestMain(m *testing.M) {
	svc = NewDockerService()
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	flag.Parse()

	m.Run()
}

func TestCreateServerVanilla(t *testing.T) {
	id, err := svc.CreateServer(ctx, *server)
	assert.NoError(t, err, fmt.Sprintf("Unable to create server due: %v", err))

	t.Logf("Current id: %s", id)
}

// DONE
func TestDeleteServer(t *testing.T) {
	id, err := svc.CreateServer(ctx, "vanilla")
	assert.NoError(t, err)

	err = svc.DeleteServer(ctx, id)
	assert.NoError(t, err)
}
