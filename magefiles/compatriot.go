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
	"github.com/magefile/mage/sh"
)

const (
	compatriotPath = "./e2e/compatriot/"
)

type Compatriot mg.Namespace

func (c Compatriot) Setup() error {
	LogGreen("Building local docker base image...")

	// if err := (Cosmos{}).Docker("base", "arm64"); err != nil {
	// 	LogRed("Failed to build local docker base image...")
	// 	return err
	// }

	// (Cosmos{}).RunDockerLocal()

	LogGreen("Building compatriot...")

	return ExecuteInDirectory(compatriotPath, func(...string) error {
		return sh.RunV("go", "build", "-v", ".")
	}, false)
}

func (c Compatriot) Test() error {
	LogGreen("Running compatriot...")

	return ExecuteInDirectory(compatriotPath, func(...string) error {
		return sh.RunV("./compatriot")
	}, false)
}

func (c Compatriot) TestV() error {
	LogGreen("Running compatriot...")

	return ExecuteInDirectory(compatriotPath, func(...string) error {
		return sh.RunV("./compatriot", "-v")
	}, false)
}

func (c Compatriot) Run() error {
	LogGreen("Building and running compatriot...")

	if err := c.Setup(); err != nil {
		LogRed("Failed to build compatriot...")
		return err
	}

	return c.TestV()
}
