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

package container

import (
	dt "github.com/ory/dockertest"
)

type Client interface {
	Start() error
	Stop() error
	Remove() error
}

type client struct {
	pool     *dt.Pool
	resource *dt.Resource
}

// NewClient creates a new ContainerClient which implements Container.
func NewClient(cfg Config) (Client, error) {
	if err := cfg.ValidateBasic(); err != nil {
		return nil, err
	}

	pool, err := dt.NewPool("")
	if err != nil {
		return nil, err
	}

	runOpts := &dt.RunOptions{
		Name:         cfg.Name,
		Repository:   cfg.Repository,
		Tag:          cfg.Tag,
		ExposedPorts: []string{cfg.HTTPAddress, cfg.WSAddress},
		Env:          cfg.Env,
	}

	resource, err := pool.RunWithOptions(runOpts)
	if err != nil {
		return nil, err
	}

	return &client{
		pool:     pool,
		resource: resource,
	}, nil
}

// Start starts the container.
func (c *client) Start() error {
	return c.pool.Client.StartContainer(c.resource.Container.ID, nil)
}

// Stop stops the container.
func (c *client) Stop() error {
	return c.pool.Client.StopContainer(c.resource.Container.ID, 0)
}

// Remove removes the container.
func (c *client) Remove() error {
	return c.resource.Close()
}
