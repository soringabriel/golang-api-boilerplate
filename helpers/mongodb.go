package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupMongoDBDockerContainer() (string, func(), error) {
	ctx := context.Background()

	internalPort := nat.Port("27017/tcp")

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort(internalPort).WithStartupTimeout(120 * time.Second),
	}

	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", nil, err
	}

	// Use the container's internal hostname (Docker network name)
	host, err := mongoContainer.Host(ctx)
	if err != nil {
		return "", nil, err
	}

	// Get mapped port (actual externally accessible port)
	mappedPort, err := mongoContainer.MappedPort(ctx, internalPort)
	if err != nil {
		return "", nil, err
	}

	// MongoDB URI
	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, mappedPort.Port())

	// Cleanup function
	cleanup := func() {
		_ = mongoContainer.Terminate(ctx)
	}

	return mongoURI, cleanup, nil
}
