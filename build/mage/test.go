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
	"os"

	"github.com/magefile/mage/sh"

	mi "pkg.berachain.dev/polaris/build/mage/internal"
)

var (

	// Arguments.
	junitArgs = []string{"--junit-report", "out.xml"}
	coverArgs = append(junitArgs, []string{"--cover", "--coverprofile",
		"coverage-testunitcover.txt", "--covermode", "atomic"}...)
	raceArgs = append(junitArgs, []string{"-race"}...)

	// Commands.
	goTest     = mi.RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest = mi.RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")

	// Packages.
	packagesUnit = mi.GoListFilter(false, "integration", "e2e", "build")
	// packagesIntegration = mi.GoListFilter(true, "integration").
	packagesEvm = mi.GoListFilter(true, "evm")
)

// Starts a testnet and builds it if necessary.
func Start() error {
	if err := Build(); err != nil {
		return err
	}
	return StartNoBuild()
}

// Starts a testnet without building it.
func StartNoBuild() error {
	return sh.RunV("./runtime/init.sh")
}

// Starts a local docs page.
func Docs() error {
	os.Chdir("docs/web")
	defer os.Chdir("../..")
	if err := sh.RunV("yarn"); err != nil {
		return err
	}
	return sh.RunV("yarn", "dev")
}

// Runs all main tests.
func Test() error {
	tests := []func() error{testUnit /*, testIntegration, testUnitBenchmark*/}

	if err := ForgeBuild(); err != nil {
		return err
	}

	for _, t := range tests {
		if err := t(); err != nil {
			return err
		}
	}
	return nil
}

// Runs the unit tests.
func TestUnit() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	return testUnit()
}

func testUnit() error {
	return ginkgoTest("--skip", ".*integration.*")
}

// Runs the unit tests with coverage.
func TestUnitCover() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	args := []string{
		"--skip", ".*integration.*",
	}
	return ginkgoTest(append(coverArgs, args...)...)
}

// Runs the unit tests with race detection.
func TestUnitRace() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	args := []string{
		"--skip", ".*integration.*",
	}
	return ginkgoTest(append(raceArgs, args...)...)
}

// Runs the unit tests with benchmarking.
func TestUnitBenchmark() error {
	if err := ForgeBuild(); err != nil {
		return err
	}

	return testUnitBenchmark()
}

func testUnitBenchmark() error {
	args := []string{
		"-bench=.",
	}
	return goTest(
		append(args, packagesUnit...)...,
	)
}

// Runs the unit tests with benchmarking.
func TestUnitEvmBenchmark() error {
	if err := ForgeBuild(); err != nil {
		return err
	}

	args := []string{
		"-bench=.",
	}
	return goTest(
		append(args, packagesEvm...)...,
	)
}

// Runs the integration tests.
func TestIntegration() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	return testIntegration()
}

func testIntegration() error {
	args := []string{
		"-timeout", "30m",
		"--focus", ".*integration.*",
	}
	return ginkgoTest(args...)
}

// Runs the integration tests with coverage.
func TestIntegrationCover() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	return testIntegrationCover()
}

func testIntegrationCover() error {
	args := []string{
		"-timeout", "30m",
		"-coverprofile=coverage-testintegrationcover.txt",
		"--focus", ".*integration.*",
	}
	return ginkgoTest(args...)
}
