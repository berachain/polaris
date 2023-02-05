// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package mage

import (
	"github.com/magefile/mage/sh"

	mi "github.com/berachain/stargazer/build/mage/internal"
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
	packagesUnit        = mi.GoListFilter(false, "integration", "cli", "e2e", "build")
	packagesIntegration = mi.GoListFilter(true, "integration", "cli")
	packagesEvm         = mi.GoListFilter(true, "evm")
)

// Starts a node and builds it if necessary.
func Start() error {
	if err := Build(); err != nil {
		return err
	}
	return StartNoBuild()
}

// Starts a node without building it.
func StartNoBuild() error {
	return sh.RunV("./build/scripts/run-local-dev.sh")
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
	return ginkgoTest(
		append(args, packagesIntegration...)...,
	)
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
		"-coverprofile=coverage-testIntegrationCover.txt",
		"--focus", ".*integration.*",
	}
	return ginkgoTest(
		append(args, packagesIntegration...)...,
	)
}
