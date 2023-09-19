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

import "os"

/* -------------------------------------------------------------------------- */
/*                                  Commands                                  */
/* -------------------------------------------------------------------------- */

var (
	// Forge.
	forgeBuild = RunCmdV("forge", "build", "--extra-output-files", "bin", "--extra-output-files", "abi", "--silent")
	forgeClean = RunCmdV("forge", "clean")
	forgeTest  = RunCmdV("forge", "test")
	forgeFmt   = RunCmdV("forge", "fmt")

	// Docker.
	dockerBuild  = RunCmdV("docker", "build", "--rm=false")
	dockerBuildX = RunCmdV("docker", "buildx", "build", "--rm=false")

	// Buf.
	bufCommand = RunCmdV("buf")

	// Testing.
	goTest     = RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest = RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")

	// Toolchain.
	goInstall  = RunCmdV("go", "install", "-mod=readonly")
	goBuild    = RunCmdV("go", "build", "-mod=readonly")
	goRun      = RunCmdV("go", "run")
	goGenerate = RunCmdV("go", "generate")
	goModTidy  = RunCmdV("go", "mod", "tidy")
	goWorkSync = RunCmdV("go", "work", "sync")
	gitDiff    = RunCmdV("git", "diff", "--stat", "--exit-code", ".",
		"':(exclude)*.mod' ':(exclude)*.sum'")
)

/* -------------------------------------------------------------------------- */
/*                             Packages & Modules                             */
/* --------------------------------------------------------------------------. */
var (
	repoModuleDirs = readGoModulesFromGoWork("go.work")
)

const (
	cosmosSDK  = "github.com/cosmos/cosmos-sdk"
	moq        = "github.com/matryer/moq"
	golangCi   = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	golines    = "github.com/segmentio/golines"
	gosec      = "github.com/securego/gosec/v2/cmd/gosec"
	addlicense = "github.com/google/addlicense"
)

/* -------------------------------------------------------------------------- */
/*                                   Docker                                   */
/* -------------------------------------------------------------------------- */

var (
	baseImageVersion  = "polard/base:v0.0.0"
	protoImageName    = "ghcr.io/cosmos/proto-builder"
	protoImageVersion = "0.14.0"
	protoDir          = "cosmos/proto"
)

/* -------------------------------------------------------------------------- */
/*                                 Directories                                */
/* -------------------------------------------------------------------------- */

const (
	outdir             = "./bin"
	baseHiveDockerPath = "./e2e/hive/"
)

var (
	hiveClone      = os.Getenv("GOPATH") + "/src/"
	clonePath      = hiveClone + ".hive-e2e/"
	simulatorsPath = clonePath + "simulators/polaris/"
	clientsPath    = clonePath + "clients/polard/"
)
