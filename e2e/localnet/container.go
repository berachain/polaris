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

type ContainerConfig struct {
	Name        string
	ImageName   string
	HTTPAddress string
	WSAddress   string
}

type ImageBuildConfig struct {
	ImageName  string
	Context    string
	Dockerfile string
	BuildArgs  map[string]string
}

// NewDefaultContainerClient builds a container with the base image.
func NewDefaultContainerClient(config ContainerConfig, imageConfig ImageBuildConfig) (*ContainerClient, error) {
	pool, err := dt.NewPool("")
	if err != nil {
		return nil, err
	}

	baseBuildArgs := map[string]string{
		"GO_VERSION":               "1.20.4",
		"PRECOMPILE_CONTRACTS_DIR": "./contracts",
		"GOOS":                     "linux",
		"GOARCH":                   "arm64",
	}

	baseImageConfig := ImageBuildConfig{
		ImageName:  baseImageName,
		Context:    baseContext,
		Dockerfile: baseDockerfile,
		BuildArgs:  baseBuildArgs,
	}

	if err = BuildImage(pool, baseImageConfig); err != nil {
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

func (c *ContainerClient) Start() error {
	return c.pool.Client.StartContainer(c.container.ID, nil)
}

func (c *ContainerClient) Stop() error {
	return c.pool.Client.StopContainer(c.container.ID, 0)
}

func (c *ContainerClient) Build(config ImageBuildConfig) error {
	return BuildImage(c.pool, config)
}

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
