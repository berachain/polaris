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
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
	"pkg.berachain.dev/polaris/e2e/localnet/container"
)

const defaultTimeout = 30 * time.Second

// ContainerizedNode is an interface for a containerized network.
type ContainerizedNode interface {
	Start() error
	Stop() error
	Reset() error
	Remove() error
	GetHTTPAddress() string
	GetWSAddress() string
	EthClient() *ethclient.Client
	EthWsClient() *ethclient.Client
	WaitForBlock(number uint64) error
}

// containerizedNode implements ContainerizedNode.
type containerizedNode struct {
	containerClient container.Client
	httpAddress     string
	wsAddress       string
	ethClient       *ethclient.Client
	ethWsClient     *ethclient.Client
}

// NewContainerizedNode creates an implementation of Localnet using Docker. The node will be past
// block 2 by the time node is available.
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
			containerClient.Stop()
			containerClient.Remove()
		}
	}()

	// Set up the http eth client.
	var ethClient *ethclient.Client
	ethClient, err = ethclient.Dial("http://" + containerClient.GetEndpoint(httpAddress))
	if err != nil {
		return nil, err
	}

	// Create the containerized node object and wait for the chain to be past block 2.
	node := &containerizedNode{
		containerClient: containerClient,
		httpAddress:     httpAddress,
		wsAddress:       wsAddress,
		ethClient:       ethClient,
		// ethWsClient:     ethclient.NewClient(ws),
	}
	time.Sleep(10 * time.Second)
	if err = node.WaitForBlock(2); err != nil {
		return nil, err
	}

	// Set up the websocket eth client.
	ws, err := gethrpc.DialWebsocket(
		context.Background(), "ws://"+containerClient.GetEndpoint(wsAddress), "*",
	)
	if err != nil {
		fmt.Println("ERR DIALING WS", err)
		return nil, err
	}
	node.ethWsClient = ethclient.NewClient(ws)

	return node, nil
}

// Start starts the network.
func (c *containerizedNode) Start() error {
	return c.containerClient.Start()
}

// Stop stops the network.
func (c *containerizedNode) Stop() error {
	return c.containerClient.Stop()
}

// Reset stops the network, clears the database, and restarts the network.
func (c *containerizedNode) Reset() error {
	if err := c.containerClient.Stop(); err != nil {
		return err
	}

	// TODO: clear genesis / reset genesis state.
	return c.containerClient.Start()
}

// Remove removes the network.
func (c *containerizedNode) Remove() error {
	return c.containerClient.Remove()
}

// GetHTTPAddress returns the HTTP address of the network.
func (c *containerizedNode) GetHTTPAddress() string {
	return c.httpAddress
}

// GetWSAddress returns the WS address of the network.
func (c *containerizedNode) GetWSAddress() string {
	return c.wsAddress
}

// EthClient returns an Ethereum client for the network.
func (c *containerizedNode) EthClient() *ethclient.Client {
	return c.ethClient
}

// EthWsClient returns an Ethereum client for the network.
func (c *containerizedNode) EthWsClient() *ethclient.Client {
	return c.ethWsClient
}

// WaitForBlock waits for the chain to reach the given block height
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
