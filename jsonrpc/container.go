// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package jsonrpc

import (
	"context"

	"github.com/berachain/stargazer/lib/errors"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	httpPort  = "8545/tcp"
	wsPort    = "8546/tcp"
	imageName = "jsonrpc-server"
	imageTag  = "dev"
)

var (
	goVersion   = "1.20.1"
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
			WaitingFor: wait.ForListeningPort(httpPort),
			// TODO: switch to this after websockets confirmed to work.
			// WaitingFor: (&wait.MultiStrategy{
			// 	Strategies: []wait.Strategy{
			// 		wait.ForListeningPort(httpPort),
			// 		wait.ForListeningPort(wsPort),
			// 	},
			// }),
		},
		Started: true,
	}

	// Create a container from the request.
	container, err := tc.GenericContainer(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "getting request provider")
	}

	// Retrieve the host mapped to the container.
	host, err := container.Host(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting host")
	}

	// Retrieve the port on the host mapped to the JSON-RPC HTTP port in the container.
	mappedHTTPPort, err := container.MappedPort(ctx, nat.Port(httpPort))
	if err != nil {
		return nil, errors.Wrapf(err, "getting mapped port for (%s)", httpPort)
	}
	config.MappedHTTP = mappedHTTPPort.Port()

	// Retrieve the port on the host mapped to the JSON-RPC WS port in the container.
	mappedWSPort, err := container.MappedPort(ctx, nat.Port(wsPort+"/tcp"))
	if err != nil {
		return nil, errors.Wrapf(err, "getting mapped port for (%s)", wsPort)
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
