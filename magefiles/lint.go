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

const (
	golangCi   = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	golines    = "github.com/segmentio/golines"
	gosec      = "github.com/securego/gosec/v2/cmd/gosec"
	addlicense = "github.com/google/addlicense"
)

func Lint() error {
	cmds := []func() error{GolangCiLint, LicenseCheck, Gosec, Proto{}.Lint, Contracts{}.Fmt}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Run all formatters.
func Format() error {
	cmds := []func() error{Golines, License, GolangCiLintFix, Proto{}.Format, Contracts{}.Fmt}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Run `golangci-lint`.
func GolangCiLint() error {
	LogGreen("Running golangci-lint...")
	for _, dir := range moduleDirs {
		if err := goRun(golangCi,
			"run", "--timeout=10m", "--concurrency", "4", "--config=.golangci.yaml", "-v", "./"+dir+"/"+"...",
		); err != nil {
			return err
		}
	}
	return nil
}

// Run `golangci-lint` with --fix.
func GolangCiLintFix() error {
	LogGreen("Running golangci-lint --fix...")
	for _, dir := range moduleDirs {
		if err := goRun(golangCi,
			"run", "--timeout=10m", "--concurrency", "4", "--config=.golangci.yaml", "-v", "--fix", "./"+dir+"/"+"...",
		); err != nil {
			return err
		}
	}
	return nil
}

// Run `golines`.
func Golines() error {
	LogGreen("Running golines...")
	return goRun(golines,
		"--reformat-tags", "--shorten-comments", "--write-output", "--max-len=99", "-l", "./.",
	)
}

// Run `gosec`.
func Gosec() error {
	LogGreen("Running gosec...")
	return goRun(gosec, "-exclude-generated", "./...")
}

// Run `addlicense`.
func License() error {
	LogGreen("Running addlicense...")
	return ExecuteForAllModules(moduleDirs, func(args ...string) error {
		return goRun(addlicense,
			"-v", "-f", "./LICENSE.header", "./.",
		)
	}, true)
}

// Run `addlicense`.
func LicenseCheck() error {
	LogGreen("Running addlicense -check...")
	return ExecuteForAllModules(moduleDirs, func(args ...string) error {
		return goRun(addlicense,
			"-check", "-v", "-f", "./LICENSE.header", "./.",
		)
	}, true)
}
