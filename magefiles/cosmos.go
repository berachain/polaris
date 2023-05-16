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

import (
	"fmt"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// Variables and Helpers.
	production = false
	statically = false

	// Commands.
	dockerBuild = RunCmdV("docker", "build", "--rm=false")

	dockerBuildX = RunCmdV("docker", "buildx", "build", "--rm=false")
	dockerRun    = RunCmdV("docker", "run")

	// Variables.
	baseDockerPath         = "./cosmos/docker/"
	execDockerPath         = baseDockerPath + "base.Dockerfile"
	testImageVersion       = "e2e-test-dev"
	goVersion              = "1.20.4"
	precompileContractsDir = "./contracts"
)

// Compile-time assertion that we implement the interface correctly.
var _ MageModule = (*Cosmos)(nil)

// Cosmos is a namespace for Cosmos SDK related commands.
type Cosmos mg.Namespace

// directory returns the directory name for the Cosmos SDK chain.
func (Cosmos) directory() string {
	return "cosmos"
}

// ===========================================================================
// Build
// ===========================================================================

// Starts a local development net and builds it if necessary.
func Start() error {
	return sh.RunV("./cosmos/init.sh")
}

// Builds the Cosmos SDK chain.
func (Cosmos) Build() error {
	LogGreen("Building the Cosmos SDK chain...")
	cmd := "polard"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./cosmos/cmd/" + cmd,
	}
	fmt.Println(strings.Join(args, " "))
	return goBuild(args...)
}

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) BuildRelease() error {
	LogGreen("Building release version of the Cosmos SDK chain...")
	production = true
	statically = false
	return c.Build()
}

// ===========================================================================
// Docker
// ===========================================================================

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) Docker(node string) error {
	LogGreen("Build a release docker image for the Cosmos SDK chain...")
	var path string
	if node == "base" {
		path = execDockerPath
	} else {
		path = baseDockerPath + node + "/Dockerfile"
	}
	return c.dockerBuildNode("polard-"+node, path, goVersion, version)
}

func (c Cosmos) RunDockerLocal() error {
	return dockerRun("-p", "8545:8545", "polard-local:e2e-test-dev")
}

func (c Cosmos) DockerX() error {
	LogGreen("Build a release docker image for the Cosmos SDK chain...")
	return c.dockerBuildBeradWithX(goVersion)
}

func (c Cosmos) dockerBuildBeradWithX(goVersion string) error {
	return dockerBuildFn(true)(
		"--build-arg", "GO_VERSION="+goVersion,
		"--platform", "linux/amd64", // TODO: do not hard code, have ability to pass as arg
		"--build-arg", "FOUNDRY_DIR="+precompileContractsDir,
		"--build-arg", "GOOS=linux",
		"--build-arg", "GOARCH=amd64",
		"-f", execDockerPath,
		"-t", "polaris:devnet-0.1", //TODO: do not hardcode, have ability to pass as arg
		".",
	)
}

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) DockerDebug() error {
	LogGreen("Build a debug docker image for the Cosmos SDK chain...")
	return c.dockerBuildNode("debug", execDockerPath, goVersion, version)
}

// Build a docker image for berad with the supplied arguments.
func (c Cosmos) dockerBuildNode(name, dockerFilePath, goVersion, imageVersion string) error {
	return dockerBuildFn(false)(
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "FOUNDRY_DIR="+precompileContractsDir,
		"--build-arg", "GOOS=darwin",
		"--build-arg", "GOARCH=arm64",
		"-f", dockerFilePath,
		"-t", name+":"+testImageVersion, //TODO: do not hardcode, have ability to pass as arg
		".",
	)
}

// ===========================================================================
// Install
// ===========================================================================

// Installs a release version of the Cosmos SDK chain.
func (Cosmos) Install() error {
	LogGreen("Installing the Cosmos SDK chain...")
	production = true
	statically = false

	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"./cosmos/cmd/polard",
	}

	return goInstall(args...)
}

// ===========================================================================
// Test
// ===========================================================================

// Runs all main tests.
func (c Cosmos) Test() error {
	if err := TestUnit(); err != nil {
		return err
	}

	if err := TestIntegration(); err != nil {
		return err
	}
	return nil
}

// Runs all unit tests for the Cosmos SDK chain.
func (c Cosmos) TestUnit() error {
	LogGreen("Running unit tests for the Cosmos SDK chain.")
	return testUnit(c.directory())
}

// Runs all integration for the Cosmos SDK chain.
func (c Cosmos) TestIntegration() error {
	LogGreen("Running integration tests for the Cosmos SDK chain.")
	return testIntegration(c.directory() + "/testing/integration")
}

func dockerBuildFn(useX bool) func(args ...string) error {
	if useX {
		return dockerBuildX
	}
	return dockerBuild
}
