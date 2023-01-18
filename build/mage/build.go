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
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

var (
	// Commands.
	goInstall   = mi.RunCmdV("go", "install", "-mod=readonly")
	goBuild     = mi.RunCmdV("go", "build", "-mod=readonly")
	goRun       = mi.RunCmdV("go", "run")
	goGenerate  = mi.RunCmdV("go", "generate")
	goModVerify = mi.RunCmdV("go", "mod", "verify")
	goModTidy   = mi.RunCmdV("go", "mod", "tidy")

	// Directories.
	outdir = "./bin"

	// Tools.
	gitDiff = sh.RunCmd("git", "diff", "--stat", "--exit-code", ".",
		"':(exclude)*.mod' ':(exclude)*.sum'")

	// Variables and Helpers.
	cmds       = []string{""}
	production = false
	statically = false
)

// Runs `go build` on the entire project.
func Build() error {
	PrintMageName()
	// If outdir doesn't exist, create.
	_, err := target.Dir(outdir)
	if os.IsNotExist(err) {
		if err = sh.Run("mkdir", "-p", outdir); err != nil {
			return err
		}
	}

	// Build all solidity contracts.
	if err = ForgeBuild(); err != nil {
		return err
	}

	for _, cmd := range cmds {
		args := []string{
			generateBuildTags(),
			generateLinkerFlags(production, statically),
			"-o", generateOutDirectory(cmd),
			generateCmdToBuild(cmd),
		}
		if err = goBuild(args...); err != nil {
			return err
		}
	}
	return nil
}

// Runs `go build` on the entire project with the release flags.
func BuildRelease() error {
	PrintMageName()
	production = true
	statically = false

	// Verify dependencies.
	if err := goModVerify(); err != nil {
		return err
	}

	return Build()
}

// Runs `go install` on the entire project.
func Install() error {
	production = true
	statically = false

	// Verify dependencies.
	if err := goModVerify(); err != nil {
		return err
	}

	// Build all commands.
	for _, cmd := range cmds {
		err := goInstall(generateCmdToBuild(cmd))
		if err != nil {
			return err
		}
	}
	return nil
}

// Runs `go generate` on the entire project.
func Generate() error {
	return goGenerate("-x", "./...")
}

// Runs `go generate` on the entire project and verifies that no files were
// changed.
func GenerateCheck() error {
	if err := Generate(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Runs 'go tidy' on the entire project.
func Tidy() error {
	return goModTidy()
}

// Cleans the project.
func Clean() error {
	// Remove golang build artifacts.
	if err := sh.Rm("bin"); err != nil {
		return err
	}

	// Remove solidity build artifacts.
	if err := ForgeClean(); err != nil {
		return err
	}

	// Remove test cache.
	if err := sh.RunV("go", "clean", "-testcache"); err != nil {
		return err
	}

	return nil
}
