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

package main

var (
	// Commands.
	dockerBuild = RunCmdV("docker", "build", "--rm=false")

	// Variables.
	baseDockerPath         = "./build/docker/"
	beradDockerPath        = baseDockerPath + "berad.Dockerfile"
	jsonrpcDockerPath      = "./jsonrpc/Dockerfile"
	imageName              = "berachain-node"
	testImageVersion       = "e2e-test-dev"
	goVersion              = "1.20.2"
	debianStaticImage      = "gcr.io/distroless/static-debian11"
	golangAlpine           = "golang:1.20-alpine3.17"
	precompileContractsDir = "./cosmos/precompile/contracts/solidity"
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
