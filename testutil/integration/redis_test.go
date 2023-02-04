package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultHttpPortTcp = "8545/tcp"
	defaultWsPortTcp   = "8546/tcp"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	os.Setenv("DOCKER_BUILDKIT", "1")
	req := testcontainers.ContainerRequest{
		BUI
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    ".",
			Dockerfile: "Dockerfile",
		},
		// ExposedPorts: []string{defaultHttpPortTcp, defaultWsPortTcp},
		// WaitingFor:   wait.ForHTTP("/").WithPort(defaultHttpPortTcp),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 15)
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()
}
