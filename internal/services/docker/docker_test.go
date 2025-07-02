package docker

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	svc    = &Service{}
	ctx    = context.Background()
	cancel = func() {}
)

func TestMain(m *testing.M) {
	svc = NewDockerService()
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	m.Run()
}

func TestCreateServer(t *testing.T) {
	id, err := svc.CreateServer(ctx)
	assert.NoError(t, err)

	t.Logf("Current id: %s", id)
}

func TestDeleteServer(t *testing.T) {
	id, err := svc.CreateServer(ctx)
	assert.NoError(t, err)

	err = svc.DeleteServer(ctx, id)
	assert.NoError(t, err)
}
