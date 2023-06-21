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
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// custom tests for Polaris, struct follows {namespace, files_changed}
// note: this requires an existing instance of the namespace in the Hive repo
type tests struct {
	Name  string
	Files []string
}

const (
	baseHiveDockerPath = "./e2e/hive/"
)

var (
	// Variables.
	hiveClone      = os.Getenv("GOPATH") + "/src/"
	clonePath      = hiveClone + ".hive-e2e/"
	simulatorsPath = clonePath + "simulators/polaris/"
	clientsPath    = clonePath + "clients/polard/"

	simulations = []tests{
		{"rpc", []string{"init/genesis.json", "ethclient.hive"}},
		{"rpc-compat", []string{"Dockerfile", "tests"}},
		{"graphql", []string{"testcases", "init/testGenesis.json"}}}
)

type Hive mg.Namespace

func (h Hive) Setup() error {
	LogGreen("Setting up Hive testing environment...")

	if _, err := os.Stat(hiveClone); os.IsNotExist(err) {
		LogGreen(hiveClone + " does not exist, creating....")
		err = os.Mkdir(hiveClone, 0755) //#nosec
		if err != nil {
			return err
		}
	}

	if err := ExecuteInDirectory(hiveClone, func(...string) error {
		LogGreen("Removing existing files in .hive-e2e...")
		return sh.RunV("rm", "-rf", clonePath)
	}, false); err != nil {
		return err
	}

	if _, err := os.Stat(clonePath); os.IsNotExist(err) {
		LogGreen("Cloning ethereum/hive into " + clonePath + "...")
		err = ExecuteInDirectory(hiveClone, func(...string) error {
			return sh.RunV("git", "clone", "https://github.com/ethereum/hive", ".hive-e2e", "--depth=1")
		}, false)
		if err != nil {
			return err
		}
	}

	if err := h.copyFiles(); err != nil {
		return err
	}

	return ExecuteInDirectory(clonePath, func(...string) error {
		LogGreen("Building Hive...")
		return goBuild(".")
	}, false)
}

func (h Hive) Test(sim, client string) error {
	return ExecuteInDirectory(clonePath, func(...string) error {
		return sh.RunV("./hive", "--sim", sim, "--client", client)
	}, false)
}

func (h Hive) TestV(sim, client string) error {
	return ExecuteInDirectory(clonePath, func(...string) error {
		return sh.RunV("./hive", "--sim", sim, "--client", client, "--docker.output")
	}, false)
}

func (h Hive) View() error {
	if err := ExecuteInDirectory(clonePath, func(...string) error {
		LogGreen("Building HiveView...")
		return sh.RunV("go", "build", "./cmd/hiveview")
	}, false); err != nil {
		return err
	}
	return ExecuteInDirectory(clonePath, func(...string) error {
		LogGreen("Serving HiveView...")
		return sh.RunV("./hiveview", "--serve")
	}, false)
}

func (h Hive) copyFiles() error {
	LogGreen("Copying Polaris Hive setup files...")
	if err := sh.RunV("mkdir", simulatorsPath); err != nil {
		return err
	}
	if err := sh.RunV("cp", "-rf", baseHiveDockerPath+"clients/polard", clientsPath); err != nil {
		return err
	}
	for _, sim := range simulations {
		if err := sh.RunV("cp", "-rf", clonePath+"simulators/ethereum/"+sim.Name, simulatorsPath+sim.Name); err != nil {
			return err
		}
		for _, file := range sim.Files {
			name := file
			if ext := strings.Split(file, "."); len(ext) > 1 && ext[1] == "hive" {
				name = strings.Split(file, ".")[0] + ".go"
			}
			if err := sh.RunV("rm", "-rf", simulatorsPath+sim.Name+"/"+name); err != nil {
				return err
			}
			if err := sh.RunV("cp", "-rf", baseHiveDockerPath+"simulators/"+sim.Name+
				"/"+file, simulatorsPath+sim.Name+"/"+name); err != nil {
				return err
			}
		}
	}
	return nil
}
