// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package localnet

import "pkg.berachain.dev/polaris/e2e/localnet/container"

// TODO: Move this into new test fixture, when we have one.
const (
	baseImageName  = "polard/base:v0.0.0"
	baseContext    = "../../"
	baseDockerfile = "./cosmos/docker/base.Dockerfile"

	localnetImageName  = "polard/localnet:v0.0.0"
	localnetContext    = "./"
	localnetDockerfile = "Dockerfile"
)

// ContainerizedNetwork is an interface for a containerized network.
type ContainerizedNetwork interface {
	Start() error
	Stop() error
	Reset() error
	GetHTTPAddress() string
	GetWSAddress() string
}

// containerizedNetwork implements ContainerizedNetwork.
type containerizedNetwork struct {
	containerClient container.Client
	httpAddress     string
	wsAddress       string
	imageConfig     container.ImageBuildConfig
}

// NewDockerizedNetwork creates an implementation of Localnet using Docker.
func NewContainerizedNetwork(
	name string,
	imageName string,
	context string,
	dockerfile string,
	httpAddress string,
	wsAddress string,
	buildArgs map[string]string,
) (ContainerizedNetwork, error) {
	// Create the container config using the given input args.
	config := container.Config{
		Name:        name,
		ImageName:   imageName,
		HTTPAddress: httpAddress,
		WSAddress:   wsAddress,
	}

	// Create the image config using the given input args.
	imageConfig := container.ImageBuildConfig{
		ImageName:  imageName,
		Context:    context,
		Dockerfile: dockerfile,
		BuildArgs:  buildArgs,
	}

	containerClient, err := container.NewClient(config, imageConfig)
	if err != nil {
		return nil, err
	}

	if err = containerClient.Build(imageConfig); err != nil {
		return nil, err
	}

	return &containerizedNetwork{
		containerClient: containerClient,
		httpAddress:     httpAddress,
		wsAddress:       wsAddress,
		imageConfig:     imageConfig,
	}, nil
}

// Start starts the network.
func (c *containerizedNetwork) Start() error {
	return c.containerClient.Start()
}

// Stop stops the network.
func (c *containerizedNetwork) Stop() error {
	return c.containerClient.Stop()
}

// Reset stops the network, clears the database, and restarts the network.
func (c *containerizedNetwork) Reset() error {
	if err := c.containerClient.Stop(); err != nil {
		return err
	}

	// TODO: clear genesis / reset genesis state.
	return c.containerClient.Start()
}

// GetHTTPAddress returns the HTTP address of the network.
func (c *containerizedNetwork) GetHTTPAddress() string {
	return c.httpAddress
}

// GetWSAddress returns the WS address of the network.
func (c *containerizedNetwork) GetWSAddress() string {
	return c.wsAddress
}
