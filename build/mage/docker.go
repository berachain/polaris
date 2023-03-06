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

package mage

import mi "pkg.berachain.dev/polaris/build/mage/internal"

var (
	// Commands.
	dockerBuild = mi.RunCmdV("docker", "build", "--rm=false")

	// Variables.
	baseDockerPath         = "./build/docker/"
	beradDockerPath        = baseDockerPath + "berad.Dockerfile"
	jsonrpcDockerPath      = "./jsonrpc/Dockerfile"
	imageName              = "berachain-node"
	testImageVersion       = "e2e-test-dev"
	goVersion              = "1.20.1"
	debianStaticImage      = "gcr.io/distroless/static-debian11"
	golangAlpine           = "golang:1.20-alpine3.17"
	precompileContractsDir = "./precompile/contracts/solidity"
)

// Build a lightweight docker image for berad.
func DockerGen() error {
	return dockerBuildBeradWith(goVersion, debianStaticImage, version)
}

// Build a debuggable docker image for berad.
func DockerDebug() error {
	return dockerBuildBeradWith(goVersion, golangAlpine, version)
}

// Build a docker image for berad with e2e test dependencies.
func DockerE2eTest() error {
	return dockerBuildBeradWith(goVersion, golangAlpine, testImageVersion)
}

func DockerBuildJSONRPCServer() error {
	return dockerBuild(
		"-f", jsonrpcDockerPath,
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "RUNNER_IMAGE="+debianStaticImage,
		"-t", "jsonrpc-server:dev",
		".",
	)
}

// Build a docker image for berad with the supplied arguments.
func dockerBuildBeradWith(goVersion, runnerImage, imageVersion string) error {
	return dockerBuild(
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "RUNNER_IMAGE="+runnerImage,
		"--build-arg", "PRECOMPILE_CONTRACTS_DIR="+precompileContractsDir,
		"-f", beradDockerPath,
		"-t", imageName+":"+imageVersion,
		".",
	)
}
