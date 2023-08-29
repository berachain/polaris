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

import "github.com/magefile/mage/mg"

type Eth mg.Namespace

func (Eth) directory() string {
	return "eth"
}

// ===========================================================================
// Test
// ===========================================================================

func (e Eth) Test() error {
	return testUnit(e.directory())
}

// Runs all unit tests for the Cosmos SDK chain.
func (e Eth) TestUnit() error {
	LogGreen("Running all Polaris Ethereum unit tests...")
	return testUnit(e.directory())
}

// Runs all e2e tests for the Cosmos SDK chain.
func (e Eth) TestE2E() error {
	LogGreen("Running all Polaris Ethereum e2e tests...")
	return testE2E(e.directory())
}
