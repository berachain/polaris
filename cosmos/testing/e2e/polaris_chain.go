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

package e2e

import (
	"context"

	tc "github.com/testcontainers/testcontainers-go"

	"pkg.berachain.dev/polaris/testing"
)

var _ testing.EthereumLikeChain = &PolarisChain{}

type PolarisChain struct {
	container *testing.Container
}

func (pc *PolarisChain) Start(ctx context.Context) error {
	// Implement your method here
	return pc.container.Start(ctx)
}

func (pc *PolarisChain) Stop() error {
	// Implement your method here
	return pc.container.Stop()
}

func (pc *PolarisChain) Reset(ctx context.Context) error {
	// Implement your method here
	return nil
}

func (pc *PolarisChain) SetGenesis(genesisPath string) {
	// Implement your method here

}

func (pc *PolarisChain) GetGenesis() string {
	return ""
}

func NewPolarisChainWithGenesis(genesisPath string) (*PolarisChain, error) {
	return NewPolarisChainFromDockerfile("../../..", "./cosmos/docker/local/Dockerfile", nil)
}

// NewEthChainFromDockerfile creates a new container from the provided configuration.
func NewPolarisChainFromDockerfile(dockerContext, dockerfilePath string,
	buildArgs map[string]*string) (*PolarisChain, error) {
	if buildArgs == nil {
		buildArgs = make(map[string]*string)
	}
	dockerReq := tc.FromDockerfile{
		Context:       dockerContext,
		Dockerfile:    dockerfilePath,
		BuildArgs:     buildArgs,
		PrintBuildLog: true,
	}

	container, err := testing.NewContainerBindingFromDockerfile(context.Background(), dockerReq)
	if err != nil {
		return nil, err
	}

	return &PolarisChain{container: container}, nil
}
