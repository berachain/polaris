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
	"github.com/magefile/mage/mg"
)

// Compile-time assertion that we implement the interface correctly.
var _ MageModule = (*Playground)(nil)

// Playground is a namespace for Cosmos SDK related commands.
type Playground mg.Namespace

// directory returns the directory name for the Playground chain.
func (Playground) directory() string {
	return "playground"
}

// Build builds the Playground app.
func (Playground) Build() error {
	LogGreen("Building the Playground chain...")
	cmd := "playground"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./playground/cmd/",
	}
	return goBuild(args...)
}

// ===========================================================================
// Install
// ===========================================================================

// Installs a release version of the Playground chain.
func (Playground) Install() error {
	LogGreen("Installing the Playground chain...")
	production = true
	statically = false

	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"./playground/cmd/",
	}

	return goInstall(args...)
}

// ===========================================================================
// Test
// ===========================================================================

// Runs all main tests.
func (p Playground) Test() error {
	// return testUnit(p.directory())
	return nil
}

// Runs all unit tests for the Cosmos SDK chain.
func (p Playground) TestUnit() error {
	// return testUnit(p.directory())
	return nil
}

// Runs all integration for the Cosmos SDK chain.
func (p Playground) TestIntegration() error {
	// return testIntegration(p.directory())
	return nil
}
