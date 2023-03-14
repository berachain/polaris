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

	"github.com/magefile/mage/sh"
)

var (

	// Arguments.
	junitArgs = []string{"--junit-report", "out.xml"}
	coverArgs = append(junitArgs, []string{"--cover", "--coverprofile",
		"coverage-testunitcover.txt", "--covermode", "atomic"}...)
	raceArgs = append(junitArgs, []string{"-race"}...)

	// Commands.
	goTest     = RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest = RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")
)

// Starts a local docs page.
func Docs() error {
	_ = os.Chdir("docs/web")
	defer func() { _ = os.Chdir("../..") }()
	if err := sh.RunV("yarn"); err != nil {
		return err
	}
	return sh.RunV("yarn", "dev")
}

// Runs the unit tests.
func TestUnit() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	PrintMageName()
	return testUnit(".")
}

func testUnit(path string) error {
	return ginkgoTest("--skip", ".*integration.*", path+"/...")
}

// Runs the unit tests with coverage.
func TestUnitCover() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	args := []string{
		"--skip", ".*integration.*",
	}
	return ginkgoTest(append(coverArgs, args...)...)
}

// Runs the unit tests with race detection.
func TestUnitRace() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	args := []string{
		"--skip", ".*integration.*",
	}
	PrintMageName()
	return ginkgoTest(append(raceArgs, args...)...)
}

// Runs the unit tests with benchmarking.
func TestUnitBenchmark() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	PrintMageName()
	return testUnitBenchmark()
}

func testUnitBenchmark() error {
	args := []string{
		"-bench=.",
	}
	return goTest(args...)
}

// Runs the unit tests with benchmarking.
func TestUnitEvmBenchmark() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}

	args := []string{
		"-bench=.",
	}
	return goTest(args...)
}

// Runs the integration tests.
func TestIntegration() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	PrintMageName()
	return testIntegration(".")
}

func testIntegration(path string) error {
	args := []string{
		"-timeout", "30m",
		"--focus", ".*integration.*", path + "/...",
	}
	return ginkgoTest(args...)
}

// Runs the integration tests with coverage.
func TestIntegrationCover() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	PrintMageName()
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
