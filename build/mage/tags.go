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

import (
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	sdkRepo        = "github.com/cosmos/cosmos-sdk"
	version        = "0.0.0"
	commit, _      = sh.Output("git", "log", "-1", "--format='%H'")
	defaultDB      = "pebbledb"
	ledgerEnabled  = true
	appName        = "berachain"
	executableName = "berad"
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
