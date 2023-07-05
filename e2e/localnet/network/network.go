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

// ContainerizedNetwork is an interface for a containerized network.
type ContainerizedNetwork interface {
	Start() error
	Stop() error
	Reset() error
	Remove() error
	GetHTTPAddress() string
	GetWSAddress() string
}

// containerizedNetwork implements ContainerizedNetwork.
type containerizedNetwork struct {
	containerClient container.Client
	httpAddress     string
	wsAddress       string
}

// NewDockerizedNetwork creates an implementation of Localnet using Docker.
func NewContainerizedNetwork(
	repository string,
	tag string,
	name string,
	httpAddress string,
	wsAddress string,
	env []string,
) (ContainerizedNetwork, error) {
	// Create the container config using the given input args.
	config := container.Config{
		Repository:  repository,
		Tag:         tag,
		Name:        name,
		HTTPAddress: httpAddress,
		WSAddress:   wsAddress,
		Env:         env,
	}

	containerClient, err := container.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &containerizedNetwork{
		containerClient: containerClient,
		httpAddress:     httpAddress,
		wsAddress:       wsAddress,
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

// Remove removes the network.
func (c *containerizedNetwork) Remove() error {
	return c.containerClient.Remove()
}

// GetHTTPAddress returns the HTTP address of the network.
func (c *containerizedNetwork) GetHTTPAddress() string {
	return c.httpAddress
}

// GetWSAddress returns the WS address of the network.
func (c *containerizedNetwork) GetWSAddress() string {
	return c.wsAddress
}
