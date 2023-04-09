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

package main

import (
	"github.com/magefile/mage/mg"
	"pkg.berachain.dev/polaris/magefiles/cmd"
)

// Compile-time assertion that we implement the interface correctly.
var _ MageModule = (*Contracts)(nil)

// Contracts is a namespace for smart contract related commands.
type Contracts mg.Namespace

func (Contracts) directory() string {
	return "contracts"
}

// Runs `forge build` in all smart contract directories.
func (Contracts) Build() error {
	return cmd.ForgeBuild()
}

// Check that the generated forge build source files are up to date.
func (c Contracts) BuildCheck() error {
	return cmd.ForgeBuildCheck()
}

// Run `forge clean` in all smart contract directories.
func (Contracts) Clean() error {
	return cmd.ForgeClean()
}

// Run `forge test` in all smart contract directories.
func (c Contracts) Test() error {
	return cmd.ForgeTest()
}

// Run `forge test` in all smart contract directories.
func (Contracts) TestUnit() error {
	return cmd.ForgeTestUnit()
}

// Run `forge fmt` in all smart contract directories.
func (Contracts) Fmt() error {
	return cmd.ForgeFmt()
}
