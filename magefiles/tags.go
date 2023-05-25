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
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	sdkRepo        = "github.com/cosmos/cosmos-sdk"
	version        = "v0.0.0"
	commit, _      = sh.Output("git", "log", "-1", "--format='%H'")
	defaultDB      = "pebbledb"
	ledgerEnabled  = true
	appName        = "polaris-cosmos"
	executableName = "polard"
)

// generateOutDirectory returns the output directory for a given command.
func generateOutDirectory(cmd string) string {
	return outdir + "/" + cmd
}

// generateBuildTags returns the build tags to be used when building the binary.
func generateBuildTags() string {
	tags := []string{defaultDB}
	if ledgerEnabled {
		tags = append(tags, "ledger")
	}
	return "-tags='" + strings.Join(tags, " ") + "'"
}

// generateLinkerFlags returns the linker flags to be used when building the binary.
func generateLinkerFlags(production, statically bool) string {
	baseFlags := []string{
		"-X ", sdkRepo + "/version.Name=" + executableName,
		" -X ", sdkRepo + "/version.AppName=" + appName,
		" -X ", sdkRepo + "/version.Version=" + version,
		" -X ", sdkRepo + "/version.Commit=" + commit,
		// TODO: Refactor versioning more broadly.
		// " \"-X " + sdkRepo + "/version.BuildTags=" + strings.Join(generateBuildTags(), ",") +
		" -X ", sdkRepo + "/version.DBBackend=" + defaultDB,
	}

	if production {
		baseFlags = append(baseFlags, "-w", "-s")
	}

	if statically {
		baseFlags = append(
			baseFlags,
			"-linkmode=external",
			"-extldflags",
			"\"-Wl,-z,muldefs -static\"",
		)
	}

	return "-ldflags=" + strings.Join(baseFlags, " ")
}
