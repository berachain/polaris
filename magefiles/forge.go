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
//
//nolint:forbidigo // its okay.
package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

var (
	// Commands.
	forgeBuild = RunCmdV("forge", "build", "--extra-output-files", "bin", "--extra-output-files", "abi", "--silent")
	forgeClean = RunCmdV("forge", "clean")
	forgeTest  = RunCmdV("forge", "test")
	forgeFmt   = RunCmdV("forge", "fmt")

	// Directories.
	testContractsDir = "./eth/testutil/contracts/solidity"
	allForgeDirs     = []string{testContractsDir, precompileContractsDir}
)

// Runs `forge build` in all smart contract directories.
func ForgeBuild() error {
	fmt.Println(color.Ize(color.Yellow, "Building Solidity contracts..."))
	return forgeWrapper(forgeBuild)
}

// Check that the generated forge build source files are up to date.
func ForgeBuildCheck() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Run `forge clean` in all smart contract directories.
func ForgeClean() error {
	return forgeWrapper(forgeClean)
}

// Run `forge test` in all smart contract directories.
func ForgeTest() error {
	return forgeWrapper(forgeTest)
}

// Run `forge fmt` in all smart contract directories.
func ForgeFmt() error {
	return forgeWrapper(forgeFmt)
}

// Wraps forge commands with the proper directory change.
func forgeWrapper(forgeFunc func(args ...string) error) error {
	for _, dir := range allForgeDirs {
		if err := ExecuteInDirectory(dir, forgeFunc, false); err != nil {
			return err
		}
	}
	return nil
}
