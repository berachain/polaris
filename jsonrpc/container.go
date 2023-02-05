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

package jsonrpc

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
)

const (
	httpPort  = "8545/tcp"
	wsPort    = "8546/tcp"
	imageName = "jsonrpc-server"
	imageTag  = "dev"
)

var (
	goVersion   = "1.19.5"
	runnerImage = "golang:" + goVersion + "-alpine"
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
	// `Name` is the name of the image.
	Name string
	// `ImageTag` is the tag of the image.
	ImageTag string
	// `Host` is the host mapped to the container.
	Host string
	// `MappedHTTP` is the port on the host mapped to the JSON-RPC HTTP port in the container.
	MappedHTTP string
	//	`MappedWS` is the port on the host mapped to the JSON-RPC WS port in the container.
	MappedWS string
}

// `DefaultContainerConfig` returns a default container configuration.
func DefaultContainerConfig() ContainerConfig {
	return ContainerConfig{
		Name:     imageName,
		ImageTag: imageTag,
	}
}

// `NewContainer` creates a new container from the provided configuration.
func NewContainer(ctx context.Context, config ContainerConfig) (*Container, error) {
	// Create a request to the container.
	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    "../",
				Dockerfile: "jsonrpc/Dockerfile",
				BuildArgs: map[string]*string{
					"GO_VERSION":   &goVersion,
					"RUNNER_IMAGE": &runnerImage},
				PrintBuildLog: true,
			},
			ExposedPorts: []string{
				httpPort, wsPort,
			},
			// Image:      fmt.Sprintf("%s:%s", config.Name, config.ImageTag),
			// WaitingFor: wait.ForLog("Starting"),
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

	// Retrieve the port on the host mapped to the JSON-RPC HTTP port in the container.
	mappedHTTPPort, err := container.MappedPort(ctx, nat.Port(httpPort))
	if err != nil {
		return nil, fmt.Errorf("getting mapped port for (%s): %w", httpPort, err)
	}
	config.MappedHTTP = mappedHTTPPort.Port()

	// Retrieve the port on the host mapped to the JSON-RPC WS port in the container.
	mappedWSPort, err := container.MappedPort(ctx, nat.Port(wsPort+"/tcp"))
	if err != nil {
		return nil, fmt.Errorf("getting mapped port for (%s): %w", wsPort, err)
	}
	config.MappedWS = mappedWSPort.Port()

	// Set the host to the host mapped to the container.
	config.Host = host

	// Return the container.
	return &Container{
		Container: container,
		config:    config,
	}, nil
}
