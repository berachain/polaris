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

	// Docker.
	dockerBuild  = RunCmdV("docker", "build", "--rm=false")
	dockerBuildX = RunCmdV("docker", "buildx", "build", "--rm=false")
)

/* -------------------------------------------------------------------------- */
/*                             Packages & Modules                             */
/* --------------------------------------------------------------------------. */
var (
	repoModuleDirs = readGoModulesFromGoWork("go.work")
)

/* -------------------------------------------------------------------------- */
/*                                   Docker                                   */
/* -------------------------------------------------------------------------- */

var (
	baseImageVersion = "polard/base:v0.0.0"
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
