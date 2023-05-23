// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package testing

import (
	"context"
	"fmt"
	"runtime"

	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
)

const (
	defaultHttpPort = "8545/tcp"
	defaultWsPort   = "8546/tcp"
)

// `Container` is a container for the JSON-RPC server.
type Container struct {
	// `Container` is a container for the JSON-RPC server.
	tc.Container
	// `config` is a configuration for a container.
	config ContainerConfig
}

// `ContainerConfig` is a configuration for a container.
type ContainerConfig struct {
	// `Host` is the host mapped to the container.
	Host string
	// `MappedHTTP` is the port on the host mapped to the JSON-RPC HTTP port in the container.
	MappedHTTP string
	//	`MappedWS` is the port on the host mapped to the JSON-RPC WS port in the container.
	MappedWS string
}

func NewContainerBinding(ctx context.Context, imageName string) (*Container, error) {
	// Create a request to the container.
	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			Image: imageName,
			// We want to expose the standard ethereum json-rpc ports.
			ExposedPorts: []string{
				defaultHttpPort, defaultWsPort,
			},
		},
	}

	// Create a container from the request.
	container, err := tc.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting request provider: %w", err)
	}

	// Retrieve the host mapped to the container.
	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting host for: %w", err)
	}

	// Return the container.
	return &Container{
		Container: container,
		config: ContainerConfig{
			Host: host,
		},
	}, nil
}

// `NewContainer` creates a new container from the provided configuration.
func NewContainerBindingFromDockerfile(ctx context.Context, fromDockerfileArgs tc.FromDockerfile) (*Container, error) {
	// Update the architecture to match the machine that it is running on.
	arch := runtime.GOARCH
	fromDockerfileArgs.BuildArgs["GOARCH"] = &arch

	// Create a request to the container.
	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: fromDockerfileArgs,
			// We want to expose the standard ethereum json-rpc ports.
			ExposedPorts: []string{
				defaultHttpPort, defaultWsPort,
			},
		},
		Started: true,
	}

	// Create a container from the request.
	container, err := tc.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting request provider: %w", err)
	}

	// Retrieve the host mapped to the container.
	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting host for: %w", err)
	}

	// Return the container.
	return &Container{
		Container: container,
		config: ContainerConfig{
			Host: host,
		},
	}, nil
}

// `Start` starts the container.
func (c *Container) Start(ctx context.Context) error {
	// Retrieve the port on the host mapped to the JSON-RPC HTTP port in the container.
	// We map them to a random port on the host.
	mappedHTTPPort, err := c.Container.MappedPort(ctx, nat.Port(defaultHttpPort))
	if err != nil {
		return fmt.Errorf("getting mapped port for (%s): %w", defaultHttpPort, err)
	}

	// Retrieve the port on the host mapped to the JSON-RPC WS port in the container.
	// We map them to a random port on the host.
	mappedWSPort, err := c.Container.MappedPort(ctx, nat.Port(defaultWsPort+"/tcp"))
	if err != nil {
		return fmt.Errorf("getting mapped port for (%s): %w", defaultWsPort, err)
	}

	c.config.MappedHTTP = mappedHTTPPort.Port()
	c.config.MappedWS = mappedWSPort.Port()

	return c.Container.Start(ctx)
}

// `Stop` stops the container from running.
func (c *Container) Stop() error {
	return c.Container.Terminate(context.Background())
}
