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
	"fmt"
)

type Localnet interface {
	Build() error
	Start() error
	Stop() error
	Reset() error
	SetGenesis(string) error
	GetGenesis() string
	GetHTTPAddress() string
	GetWSAddress() string
}

type dockerizedNetwork struct {
	Container
	genesis     string
	httpAddress string
	wsAddress   string

	imageConfig ImageBuildConfig
}

// NewDockerizedNetwork creates a new localnet client.
func NewDockerizedNetwork(name, imageName, genesis, httpAddress, wsAddress string) (*dockerizedNetwork, error) {
	// Check for a genesis file.
	if genesis == "" {
		return nil, fmt.Errorf("genesis cannot be empty")
	}

	// Create the container config using the given input args.
	config := ContainerConfig{
		Name:        name,
		ImageName:   imageName,
		HTTPAddress: httpAddress,
		WSAddress:   wsAddress,
	}

	// Create the container client.
	imageConfig := ImageBuildConfig{
		ImageName:  localnetImageName,
		Context:    localnetContext,
		Dockerfile: localnetDockerfile,
		BuildArgs: map[string]string{
			"BASE_IMAGE": baseImageName,
		},
	}
	container, err := NewDefaultContainerClient(config, imageConfig)
	if err != nil {
		return nil, err
	}

	return &dockerizedNetwork{
		genesis:     genesis,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
		imageConfig: imageConfig,
		Container:   container,
	}, nil
}

func (c *dockerizedNetwork) Build() error {
	return c.Container.Build(c.imageConfig)
}

func (c *dockerizedNetwork) Reset() error {
	if err := c.Stop(); err != nil {
		return err
	}
	if err := c.Build(); err != nil {
		return err
	}
	return c.Start()
}

func (c *dockerizedNetwork) SetGenesis(genesis string) error {
	// override a config file/set one
	return nil
}

func (c *dockerizedNetwork) GetGenesis() string {
	return c.genesis
}

func (c *dockerizedNetwork) GetHTTPAddress() string {
	return c.httpAddress
}

func (c *dockerizedNetwork) GetWSAddress() string {
	return c.wsAddress
}
