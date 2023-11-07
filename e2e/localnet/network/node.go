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
	"context"
	"errors"
	"time"

	"github.com/berachain/polaris/e2e/localnet/container"

	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
)

const (
	defaultTimeout = 30 * time.Second
	// TODO: this is so hood.
	nodeStartTime = 15 * time.Second
)

// ContainerizedNode is an interface for a containerized network.
type ContainerizedNode interface {
	Start() error
	Stop() error
	Reset() error
	Remove() error
	GetHTTPEndpoint() string
	GetWSEndpoint() string
	EthClient() *ethclient.Client
	EthWsClient() *ethclient.Client
	WaitForBlock(number uint64) error
	WaitForNextBlock() error
	DumpLogs() (string, error)
}

// containerizedNode implements ContainerizedNode.
type containerizedNode struct {
	containerClient container.Client
	httpEndpoint    string
	wsEndpoint      string
	ethClient       *ethclient.Client
	ethWsClient     *ethclient.Client
}

// NewContainerizedNode creates an implementation of Localnet using Docker.
//
//nolint:nonamedreturns // deferred error handling.
func NewContainerizedNode(
	repository string,
	tag string,
	name string,
	httpAddress string,
	wsAddress string,
	env []string,
) (c ContainerizedNode, err error) {
	// Create the container using the given input args for config.
	var containerClient container.Client
	containerClient, err = container.NewClient(
		container.Config{
			Repository:  repository,
			Tag:         tag,
			Name:        name,
			HTTPAddress: httpAddress,
			WSAddress:   wsAddress,
			Env:         env,
		},
	)
	if err != nil {
		return nil, err
	}

	// If we error out, make sure to stop and remove the container.
	defer func() {
		if err != nil {
			_ = containerClient.Stop()
			_ = containerClient.Remove()
		}
	}()

	// Create the containerized node object.
	node := &containerizedNode{
		containerClient: containerClient,
		httpEndpoint:    "http://" + containerClient.GetEndpoint(httpAddress),
		wsEndpoint:      "ws://" + containerClient.GetEndpoint(wsAddress),
	}

	// Set up the http eth client.
	node.ethClient, err = ethclient.Dial(node.httpEndpoint)
	if err != nil {
		return nil, err
	}

	// Wait for the chain to start.
	time.Sleep(nodeStartTime)
	if err = node.WaitForNextBlock(); err != nil {
		return nil, err
	}

	// Set up the websocket eth client.
	ws, err := gethrpc.DialWebsocket(
		context.Background(), node.wsEndpoint, "*",
	)
	if err != nil {
		return nil, err
	}
	node.ethWsClient = ethclient.NewClient(ws)

	return node, nil
}

// Start starts the node.
func (c *containerizedNode) Start() error {
	return c.containerClient.Start()
}

// Stop stops the node.
func (c *containerizedNode) Stop() error {
	return c.containerClient.Stop()
}

// Reset stops the node, clears the database, and restarts the node.
func (c *containerizedNode) Reset() error {
	if err := c.containerClient.Stop(); err != nil {
		return err
	}

	// TODO: clear genesis / reset genesis state.
	return c.containerClient.Start()
}

// Remove removes the node.
func (c *containerizedNode) Remove() error {
	return c.containerClient.Remove()
}

func (c *containerizedNode) DumpLogs() (string, error) {
	logsBz, err := c.containerClient.GetContainerLogs()
	return string(logsBz), err
}

// GetHTTPEndpoint returns the HTTP endpoint of the node.
func (c *containerizedNode) GetHTTPEndpoint() string {
	return c.httpEndpoint
}

// GetWSEndpoint returns the WS endpoint of the node.
func (c *containerizedNode) GetWSEndpoint() string {
	return c.wsEndpoint
}

// EthClient returns an Ethereum client for the node.
func (c *containerizedNode) EthClient() *ethclient.Client {
	return c.ethClient
}

// EthWsClient returns an Ethereum client for the node.
func (c *containerizedNode) EthWsClient() *ethclient.Client {
	return c.ethWsClient
}

// WaitForBlock waits for the chain to reach the given block height.
func (c *containerizedNode) WaitForBlock(number uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			currHeight, err := c.ethClient.BlockNumber(ctx)
			if err != nil {
				return err
			}

			if currHeight > number {
				return errors.New("block height already passed")
			}

			if currHeight == number {
				return nil
			}
		}
	}
}

// WaitForNextBlock waits for the chain to reach the next block.
func (c *containerizedNode) WaitForNextBlock() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var currHeight uint64
	var currDone bool
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if !currDone {
				var err error
				currHeight, err = c.ethClient.BlockNumber(ctx)
				if err != nil {
					return err
				}
				currDone = true
				continue
			}

			newHeight, err := c.ethClient.BlockNumber(ctx)
			if err != nil {
				return err
			}

			if newHeight == currHeight+1 {
				return nil
			}
		}
	}
}
