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
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// Variables.
	baseHiveDockerPath = "./e2e/hive/"
	hiveClone          = os.Getenv("GOPATH") + "/src/"
	clonePath          = hiveClone + ".hive-e2e"
	simulatorsPath     = clonePath + "/simulators/polaris"
	clientsPath        = clonePath + "/clients/polard"
)

type Hive mg.Namespace

func (Hive) directory() string {
	return "Hive"
}

func (h Hive) Setup() error {
	LogGreen("Executing Hive tests on polard client...")

	if _, err := os.Stat(hiveClone); os.IsNotExist(err) {
		LogGreen(hiveClone + " does not exist, creating....")
		os.Mkdir(hiveClone, 0755)
	}

	if _, err := os.Stat(clonePath); os.IsNotExist(err) {
		LogGreen("Cloning ethereum/hive into " + clonePath + "...")
		return ExecuteInDirectory(hiveClone, func(...string) error {
			return sh.RunV("git", "clone", "https://github.com/ethereum/hive", ".hive-e2e", "--depth=1")
		}, false)
	}

	if err := ExecuteInDirectory(clonePath, func(...string) error {
		LogGreen("Building Hive...")
		return goBuild(".")
	}, false); err != nil {
		return err
	}

	LogGreen("Copying Polaris Hive setup files...")
	if err := sh.RunV("cp", "-rf", baseHiveDockerPath+"clients/polard", clientsPath); err != nil {
		return err
	}
	if err := sh.RunV("cp", "-rf", "./e2e/hive/simulators", simulatorsPath); err != nil {
		return err
	}

	return ExecuteInDirectory(clonePath, func(...string) error {
		LogGreen("Building HiveView...")
		return goBuild(".")
	}, false)
}

func (h Hive) Test(sim, client string) error {
	return ExecuteInDirectory(clonePath, func(...string) error {
		return sh.RunV("./hive", "--sim", sim, "--client", client)
	}, false)
}
