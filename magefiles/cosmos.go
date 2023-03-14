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

	"github.com/TwiN/go-color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// Variables and Helpers.
	production = false
	statically = false

	// Commands.
	dockerBuild = RunCmdV("docker", "build", "--rm=false")

	// Variables.
	baseDockerPath  = "./build/docker/"
	beradDockerPath = baseDockerPath + "berad.Dockerfile"
	imageName       = "polaris-cosmos"
	// testImageVersion       = "e2e-test-dev".
	goVersion              = "1.20.2"
	debianStaticImage      = "gcr.io/distroless/static-debian11"
	golangAlpine           = "golang:1.20-alpine3.17"
	precompileContractsDir = "./cosmos/precompile/contracts/solidity"
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
	return sh.RunV("./cosmos/runtime/init.sh")
}

// Builds the Cosmos SDK chain.
func (Cosmos) Build() error {
	cmd := "polard"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./cosmos/cmd/" + cmd,
	}
	fmt.Println(color.Ize(color.Yellow, "Building Cosmos app..."))
	return goBuild(args...)
}

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) BuildRelease() error {
	PrintMageName()
	production = true
	statically = false
	return c.Build()
}

// ===========================================================================
// Docker
// ===========================================================================

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) BuildDocker() error {
	PrintMageName()
	return c.dockerBuildBeradWith(goVersion, debianStaticImage, version)
}

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) BuildDockerDebug() error {
	PrintMageName()
	return c.dockerBuildBeradWith(goVersion, golangAlpine, version)
}

// Build a docker image for berad with the supplied arguments.
func (c Cosmos) dockerBuildBeradWith(goVersion, runnerImage, imageVersion string) error {
	return dockerBuild(
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "RUNNER_IMAGE="+runnerImage,
		"--build-arg", "PRECOMPILE_CONTRACTS_DIR="+precompileContractsDir,
		"-f", beradDockerPath,
		"-t", imageName+":"+imageVersion,
		".",
	)
}

// ===========================================================================
// Install
// ===========================================================================

// Installs a release version of the Cosmos SDK chain.
func (Cosmos) Install() error {
	PrintMageName()
	production = true
	statically = false

	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"./cmd/polard",
	}

	return goInstall(args...)
}

// ===========================================================================
// Test
// ===========================================================================

// Runs all main tests.
func (c Cosmos) Test() error {
	PrintMageName()
	return testUnit(c.directory())
}

// Runs all unit tests for the Cosmos SDK chain.
func (c Cosmos) TestUnit() error {
	PrintMageName()
	return testUnit(c.directory())
}

// Runs all integration for the Cosmos SDK chain.
func (c Cosmos) TestIntegration() error {
	PrintMageName()
	return testIntegration(c.directory())
}
