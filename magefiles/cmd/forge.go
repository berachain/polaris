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

//

package cmd

import (
	"pkg.berachain.dev/polaris/magefiles/cmd/utils"
)

var (
	// Commands.
	forgeBuild = utils.RunCmdV("forge", "build", "--extra-output-files", "bin", "--extra-output-files", "abi", "--silent")
	forgeClean = utils.RunCmdV("forge", "clean")
	forgeTest  = utils.RunCmdV("forge", "test")
	forgeFmt   = utils.RunCmdV("forge", "fmt")
)

// ===========================================================================
// Build
// ===========================================================================

// Runs `forge build` in all smart contract directories.
func ForgeBuild() error {
	LogGreen("Building solidity contracts with foundry...")
	return forgeWrapper(forgeBuild)
}

// Check that the generated forge build source files are up to date.
func ForgeBuildCheck() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	// if err := gitDiff(); err != nil {
	// 	return fmt.Errorf("generated files are out of date: %w", err)
	// }
	return nil
}

// Run `forge clean` in all smart contract directories.
func ForgeClean() error {
	return forgeWrapper(forgeClean)
}

// ===========================================================================
// Test
// ===========================================================================

// Run `forge test` in all smart contract directories.
func ForgeTest() error {
	if err := ForgeTestUnit(); err != nil {
		return err
	}
	return nil
}

// Run `forge test` in all smart contract directories.
func ForgeTestUnit() error {
	LogGreen("Running foundry unit tests...")
	return forgeWrapper(forgeTest)
}

// Run `forge fmt` in all smart contract directories.
func ForgeFmt() error {
	LogGreen("Running forge fmt...")
	return forgeWrapper(forgeFmt)
}

func ForgeTestIntegration() error {
	LogGreen("Running foundry integration tests...")
	return forgeWrapper(forgeTest)
}

// ===========================================================================
// Helper
// ===========================================================================

// Wraps forge commands with the proper directory change.
func forgeWrapper(forgeFunc func(args ...string) error) error {
	if err := utils.ExecuteInDirectory("./contracts", forgeFunc, false); err != nil {
		return err
	}
	return nil
}
