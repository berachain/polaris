// // SPDX-License-Identifier: MIT
// //
// // Copyright (c) 2023 Berachain Foundation
// //
// // Permission is hereby granted, free of charge, to any person
// // obtaining a copy of this software and associated documentation
// // files (the "Software"), to deal in the Software without
// // restriction, including without limitation the rights to use,
// // copy, modify, merge, publish, distribute, sublicense, and/or sell
// // copies of the Software, and to permit persons to whom the
// // Software is furnished to do so, subject to the following
// // conditions:
// //
// // The above copyright notice and this permission notice shall be
// // included in all copies or substantial portions of the Software.
// //
// // THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// // EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// // OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// // NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// // HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// // WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// // FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// // OTHER DEALINGS IN THE SOFTWARE.

package cmd

// import (
// 	"strings"

// 	"github.com/magefile/mage/sh"
// 	"pkg.berachain.dev/polaris/magefiles/cmd/utils"
// )

// var (
// 	// Variables and Helpers.
// 	production = false
// 	statically = false

// 	// Commands.
// 	dockerBuild = utils.RunCmdV("docker", "build", "--rm=false")

// 	// Variables.
// 	baseDockerPath  = "./cosmos/"
// 	beradDockerPath = baseDockerPath + "Dockerfile"
// 	imageName       = "polaris-cosmos"
// 	// testImageVersion       = "e2e-test-dev".
// 	goVersion              = "1.20.2"
// 	debianStaticImage      = "gcr.io/distroless/static-debian11"
// 	golangAlpine           = "golang:1.20-alpine3.17"
// 	precompileContractsDir = "./cosmos/precompile/contracts/solidity"
// )

// // ===========================================================================
// // Build
// // ===========================================================================

// // Starts a local development net and builds it if necessary.
// func CosmosStart() error {
// 	return sh.RunV("./cosmos/runtime/init.sh")
// }

// // Builds the Cosmos SDK chain.
// func CosmosBuild() error {
// 	LogGreen("Building the Cosmos SDK chain...")
// 	cmd := "polard"
// 	args := []string{
// 		generateBuildTags(),
// 		generateLinkerFlags(production, statically),
// 		"-o", generateOutDirectory(cmd),
// 		"./cosmos/cmd/" + cmd,
// 	}
// 	return goBuild(args...)
// }

// // Builds a release version of the Cosmos SDK chain.
// func CosmosBuildRelease() error {
// 	LogGreen("Building release version of the Cosmos SDK chain...")
// 	production = true
// 	statically = false
// 	return CosmosBuild()
// }

// // ===========================================================================
// // Docker
// // ===========================================================================

// // Builds a release version of the Cosmos SDK chain.
// func CosmosBuildDocker() error {
// 	LogGreen("Build a release docker image for the Cosmos SDK chain...")
// 	return dockerBuildBeradWith(goVersion, debianStaticImage, version)
// }

// // Builds a release version of the Cosmos SDK chain.
// func CosmosBuildDockerDebug() error {
// 	LogGreen("Build a debug docker image for the Cosmos SDK chain...")
// 	return dockerBuildBeradWith(goVersion, golangAlpine, version)
// }

// // Build a docker image for berad with the supplied arguments.
// func dockerBuildBeradWith(goVersion, runnerImage, imageVersion string) error {
// 	return dockerBuild(
// 		"--build-arg", "GO_VERSION="+goVersion,
// 		"--build-arg", "RUNNER_IMAGE="+runnerImage,
// 		"--build-arg", "PRECOMPILE_CONTRACTS_DIR="+precompileContractsDir,
// 		"-f", beradDockerPath,
// 		"-t", imageName+":"+imageVersion,
// 		".",
// 	)
// }

// // ===========================================================================
// // Install
// // ===========================================================================

// // Installs a release version of the Cosmos SDK chain.
// func CosmosInstall() error {
// 	LogGreen("Installing the Cosmos SDK chain...")
// 	production = true
// 	statically = false

// 	args := []string{
// 		generateBuildTags(),
// 		generateLinkerFlags(production, statically),
// 		"./cmd/polard",
// 	}

// 	return goInstall(args...)
// }

// // ===========================================================================
// // Test
// // ===========================================================================

// // // Runs all main tests.
// // func CosmosTest() error {
// // 	if err := TestUnit(); err != nil {
// // 		return err
// // 	}

// // 	if err := CosmosTestIntegration(); err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// // var (
// // 	sdkRepo        = "github.com/cosmos/cosmos-sdk"
// // 	version        = "0.0.0"
// // 	commit, _      = sh.Output("git", "log", "-1", "--format='%H'")
// // 	defaultDB      = "pebbledb"
// // 	ledgerEnabled  = true
// // 	appName        = "berachain"
// // 	executableName = "berad"
// // )

// // // generateOutDirectory returns the output directory for a given command.
// // func generateOutDirectory(cmd string) string {
// // 	return outdir + "/" + cmd
// // }

// // generateBuildTags returns the build tags to be used when building the binary.
// func generateBuildTags() string {
// 	tags := []string{defaultDB}
// 	if ledgerEnabled {
// 		tags = append(tags, "ledger")
// 	}
// 	return "-tags='" + strings.Join(tags, " ") + "'"
// }

// // generateLinkerFlags returns the linker flags to be used when building the binary.
// func generateLinkerFlags(production, statically bool) string {
// 	baseFlags := []string{
// 		"-X ", sdkRepo + "/version.Name=" + executableName,
// 		" -X ", sdkRepo + "/version.AppName=" + appName,
// 		" -X ", sdkRepo + "/version.Version=" + version,
// 		" -X ", sdkRepo + "/version.Commit=" + commit,
// 		// TODO: Refactor versioning more broadly.
// 		// " \"-X " + sdkRepo + "/version.BuildTags=" + strings.Join(generateBuildTags(), ",") +
// 		" -X ", sdkRepo + "/version.DBBackend=" + defaultDB,
// 	}

// 	if production {
// 		baseFlags = append(baseFlags, "-w", "-s")
// 	}

// 	if statically {
// 		baseFlags = append(
// 			baseFlags,
// 			"-linkmode=external",
// 			"-extldflags",
// 			"\"-Wl,-z,muldefs -static\"",
// 		)
// 	}

// 	return "-ldflags=" + strings.Join(baseFlags, " ")
// }
