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
