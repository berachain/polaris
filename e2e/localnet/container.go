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

import (
	"os"

	dt "github.com/ory/dockertest"
	dc "github.com/ory/dockertest/docker"
)

type Container interface {
	Start() error
	Stop() error
	Build(ImageBuildConfig) error
}

type ContainerClient struct {
	pool      *dt.Pool
	container *dc.Container
}

// NewContainerClient creates a new ContainerClient which implements Container.
func NewContainerClient(config ContainerConfig, imageConfig ImageBuildConfig) (*ContainerClient, error) {
	pool, err := dt.NewPool("")
	if err != nil {
		return nil, err
	}

	if err = BuildImage(pool, imageConfig); err != nil {
		return nil, err
	}

	container, err := pool.Client.CreateContainer(dc.CreateContainerOptions{
		Name: config.Name,
		Config: &dc.Config{
			Image: config.ImageName,
			ExposedPorts: map[dc.Port]struct{}{
				dc.Port(config.HTTPAddress): {},
				dc.Port(config.WSAddress):   {},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &ContainerClient{
		pool:      pool,
		container: container,
	}, nil
}

// Start starts the container.
func (c *ContainerClient) Start() error {
	return c.pool.Client.StartContainer(c.container.ID, nil)
}

// Stop stops the container.
func (c *ContainerClient) Stop() error {
	return c.pool.Client.StopContainer(c.container.ID, 0)
}

// Build builds the container image.
func (c *ContainerClient) Build(config ImageBuildConfig) error {
	return BuildImage(c.pool, config)
}

// BuildImage is a helper function for building a container image.
func BuildImage(pool *dt.Pool, config ImageBuildConfig) error {
	containerBuildArgs := make([]dc.BuildArg, len(config.BuildArgs))
	i := 0
	for k, v := range config.BuildArgs {
		containerBuildArgs[i] = dc.BuildArg{
			Name:  k,
			Value: v,
		}
		i++
	}

	baseBuildOpts := dc.BuildImageOptions{
		Name:         config.ImageName,
		ContextDir:   config.Context,
		Dockerfile:   config.Dockerfile,
		BuildArgs:    containerBuildArgs,
		OutputStream: os.Stdout,
	}
	return pool.Client.BuildImage(baseBuildOpts)
}
