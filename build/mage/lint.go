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
	"strings"

	mi "github.com/berachain/stargazer/build/mage/internal"
)

const (
	golangCi   = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	golines    = "github.com/segmentio/golines"
	gosec      = "github.com/securego/gosec/v2/cmd/gosec"
	addlicense = "github.com/google/addlicense"
	goimports  = "github.com/incu6us/goimports-reviser/v3"
)

func Lint() error {
	cmds := []func() error{GolangCiLint, Gosec, LicenseCheck, ProtoLint}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Run all formatters.
func Format() error {
	cmds := []func() error{Golines, GolangCiLintFix, License, GoImports, ProtoFormat}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Run `golangci-lint`.
func GolangCiLint() error {
	PrintMageName()
	return goRun(golangCi,
		"run", "--timeout=10m", "--concurrency", "4", "--config=build/.golangci.yaml", "-v", "./...",
	)
}

// Run `golangci-lint` with --fix.
func GolangCiLintFix() error {
	PrintMageName()
	return goRun(golangCi,
		"run", "--timeout=10m", "--concurrency", "4", "--config=build/.golangci.yaml", "-v", "--fix", "./...",
	)
}

// Run `golines`.
func Golines() error {
	PrintMageName()
	return goRun(golines,
		"--reformat-tags", "--shorten-comments", "--write-output", "--max-len=99", "-l", "./.",
	)
}

func GoImports() error {
	PrintMageName()
	// everything but ignore the tools folder
	var x = make([]string, 0)
	for _, dir := range mi.GoListFilter(true, "build/tools") {
		stripped := strings.ReplaceAll(dir, "github.com/berachain", "")
		x = append(x, stripped)
	}

	for _, dir := range x {
		if err := goRun(goimports,
			"-recursive", "-rm-unused",
			"-use-cache", "-output",
			"-company-prefixes", "github.com/berachain",
			"\"write\"", "-project-name", "github.com/berachain/stargazer", dir); err != nil {
			return err
		}
	}
	return nil
}

// Check that golang imports are formatted correctly.
func GoImportsLint() error {
	PrintMageName()
	if err := GoImports(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("please run `mage goimports`: %w", err)
	}
	return nil
}

// Run `gosec`.
func Gosec() error {
	PrintMageName()
	return goRun(gosec, "./...")
}

// Run `addlicense`.
func License() error {
	PrintMageName()
	return goRun(addlicense,
		"-v", "-f", "./build/LICENSE.header", "./.")
}

// Run `addlicense` with -check.
func LicenseCheck() error {
	PrintMageName()
	return goRun(addlicense,
		"-v", "-check", "-f", "./build/LICENSE.header", "./.")
}
