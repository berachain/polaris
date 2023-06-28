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

type Localnet interface {
	Build() error
	Start() error
	Stop() error
	Reset() error
	GetHTTPAddress() string
	GetWSAddress() string
}

type dockerizedNetwork struct {
	Container
	httpAddress string
	wsAddress   string

	imageConfig ImageBuildConfig
}

// NewDockerizedNetwork creates an implementation of Localnet using Docker.
func NewDockerizedNetwork(
	name string,
	imageName string,
	context string,
	dockerfile string,
	httpAddress string,
	wsAddress string,
	buildArgs map[string]string,
) (Localnet, error) {
	if context == "" {
		return nil, ErrEmptyContext
	}
	if dockerfile == "" {
		return nil, ErrEmptyDockerfile
	}

	// Create the container config using the given input args.
	config := ContainerConfig{
		Name:        name,
		ImageName:   imageName,
		HTTPAddress: httpAddress,
		WSAddress:   wsAddress,
	}

	// Create the image config using the given input args.
	imageConfig := ImageBuildConfig{
		ImageName:  imageName,
		Context:    context,
		Dockerfile: dockerfile,
		BuildArgs:  buildArgs,
	}

	container, err := NewContainerClient(config, imageConfig)
	if err != nil {
		return nil, err
	}

	return &dockerizedNetwork{
		Container:   container,
		httpAddress: httpAddress,
		wsAddress:   wsAddress,
		imageConfig: imageConfig,
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

func (c *dockerizedNetwork) GetHTTPAddress() string {
	return c.httpAddress
}

func (c *dockerizedNetwork) GetWSAddress() string {
	return c.wsAddress
}
