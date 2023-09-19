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
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"pkg.berachain.dev/polaris/magefiles/utils"
)

var (
	// Variables and Helpers.
	production = false
	statically = false

	// Variables.
	imageName              = "polard"
	imageVersion           = "v0.0.0"
	baseDockerPath         = "./e2e/testapp/docker/"
	execDockerPath         = baseDockerPath + "base.Dockerfile"
	localDockerPath        = baseDockerPath + "local/Dockerfile"
	seedDockerPath         = baseDockerPath + "seed/Dockerfile"
	valDockerPath          = baseDockerPath + "validator/Dockerfile"
	goVersion              = "1.21.1"
	precompileContractsDir = "./contracts"

	// Localnet.
	baseImage          = "polard/base:v0.0.0"
	localnetClientPath = "./e2e/precompile/polard"
	localnetDockerPath = localnetClientPath + "/Dockerfile"
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
	return sh.RunV("./e2e/testapp/entrypoint.sh")
}

// Builds the Cosmos SDK chain.
func (Cosmos) Build() error {
	utils.LogGreen("Building the Cosmos SDK chain...")
	cmd := "polard"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./e2e/testapp/" + cmd,
	}
	return goBuild(args...)
}

// ===========================================================================
// Docker
// ===========================================================================

// Builds a release version of the Cosmos SDK chain.
func (c Cosmos) Docker(dockerType, arch string) error {
	utils.LogGreen("Build a release docker image for the Cosmos SDK chain...")
	return c.dockerBuildBeradWith(dockerType, goVersion, arch, false)
}

func (c Cosmos) dockerBuildBeradWith(dockerType, goVersion, arch string, withX bool) error {
	var dockerFilePath string
	opts := []string{
		"--build-arg", "GO_VERSION=" + goVersion,
		"--platform", "linux/" + arch,
		"--build-arg", "PRECOMPILE_CONTRACTS_DIR=" + precompileContractsDir,
		"--build-arg", "GOOS=linux",
		"--build-arg", "GOARCH=" + arch,
		"--build-arg", "GO_WORK=" + strings.Join(repoModuleDirs, " "),
	}
	buildContext := "."
	switch dockerType {
	case "local":
		dockerFilePath = localDockerPath
	case "seed":
		dockerFilePath = seedDockerPath
	case "validator":
		dockerFilePath = valDockerPath
	case "localnet":
		buildContext = localnetClientPath
		dockerFilePath = localnetDockerPath
		opts = append(opts, "--build-arg", "BASE_IMAGE="+baseImage)
	default:
		dockerFilePath = execDockerPath
	}
	tag := imageName + "/" + dockerType + ":" + imageVersion
	utils.LogGreen(
		"Building a "+dockerType+" polard docker image",
		"platform", "linux"+"/"+arch,
		"tag", tag,
	)
	opts = append(opts, "-f", dockerFilePath, "-t", tag, buildContext)
	return dockerBuildFn(withX)(
		opts...,
	)
}

// ===========================================================================
// Test
// ===========================================================================

// Runs all main tests.
func (c Cosmos) Test() error {
	if err := TestUnit(); err != nil {
		return err
	}

	return TestE2E()
}

// Runs all unit tests for the Cosmos SDK chain.
func (c Cosmos) TestUnit() error {
	utils.LogGreen("Running unit tests for the Cosmos SDK chain.")
	return testUnit(c.directory())
}

// Runs all unit tests for the Cosmos SDK chain.
func (c Cosmos) TestUnitRace() error {
	utils.LogGreen("Running unit tests for the Cosmos SDK chain.")
	return testUnitRace(c.directory())
}

// Runs all e2e tests for the Cosmos SDK chain.
func (c Cosmos) TestE2E() error {
	utils.LogGreen("Running e2e tests for the Cosmos SDK chain.")
	return testE2E(c.directory() + "/testing/e2e")
}

func (c Cosmos) TestHive(sim string) error {
	if out, _ := sh.Output("docker", "images", "-q", baseImageVersion); out == "" {
		utils.LogGreen("No existing base docker image found, building...")
		if err := c.Docker("base", runtime.GOARCH); err != nil {
			return err
		}
	}

	if err := (Hive{}).Setup(); err != nil {
		return err
	}

	return Hive{}.TestV(sim, "polard")
}

func dockerBuildFn(useX bool) func(args ...string) error {
	if useX {
		return dockerBuildX
	}
	return dockerBuild
}
