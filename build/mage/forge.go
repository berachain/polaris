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
	"fmt"
	"os"

	mi "github.com/berachain/stargazer/build/mage/internal"
)

var (
	// Commands.
	forgeBuild = mi.RunCmdV("forge", "build")
	forgeClean = mi.RunCmdV("forge", "clean")
	forgeTest  = mi.RunCmdV("forge", "test")
	forgeFmt   = mi.RunCmdV("forge", "fmt")

	// Directories.
	testContractsDir = "./testutil/contracts/solidity"
	// precompileContractsDir = "./pkg/dahlia/pkg/core/vm/precompile/contracts".
	allForgeDirs = []string{testContractsDir /*, precompileContractsDir*/}
)

// Runs `forge build` in all smart contract directories.
func ForgeBuild() error {
	return forgeWrapper(forgeBuild)
}

// Check that the generated forge build source files are up to date.
func ForgeBuildCheck() error {
	if err := ForgeBuild(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Run `forge clean` in all smart contract directories.
func ForgeClean() error {
	return forgeWrapper(forgeClean)
}

// Run `forge test` in all smart contract directories.
func ForgeTest() error {
	return forgeWrapper(forgeTest)
}

// Run `forge fmt` in all smart contract directories.
func ForgeFmt() error {
	return forgeWrapper(forgeFmt)
}

// Wraps forge commands with the proper directory change.
func forgeWrapper(forgeFunc func(args ...string) error) error {
	rootCwd, _ := os.Getwd()
	for _, dir := range allForgeDirs {
		// Change to the directory where the contracts are.
		if err := os.Chdir(dir); err != nil {
			return err
		}
		// Run the forge command.
		if err := forgeFunc(); err != nil {
			return err
		}

		// Go back to the starting directory.
		if err := os.Chdir(rootCwd); err != nil {
			return err
		}
	}
	return nil
}
