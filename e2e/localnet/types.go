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
	"errors"
)

// ContainerConfig is a configuration struct for a container.
type ContainerConfig struct {
	Name        string
	ImageName   string
	HTTPAddress string
	WSAddress   string
}

// ImageBuildConfig is a configuration struct for an image build.
type ImageBuildConfig struct {
	ImageName  string
	Context    string
	Dockerfile string
	BuildArgs  map[string]string
}

// Errors returned by the localnet package.
var (
	EmptyGenesisError    = errors.New("genesis cannot be empty")
	EmptyContextError    = errors.New("context cannot be empty")
	EmptyDockerfileError = errors.New("dockerfile cannot be empty")
)

// TODO: Move this into new test fixture, when we have one.
const (
	baseImageName  = "polard/base:v0.0.0"
	baseContext    = "../../"
	baseDockerfile = "./cosmos/docker/base.Dockerfile"

	localnetImageName  = "polard/localnet:v0.0.0"
	localnetContext    = "./"
	localnetDockerfile = "Dockerfile"
)
