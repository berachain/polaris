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
	"fmt"
	"strings"

	mi "pkg.berachain.dev/stargazer/build/mage/internal"
)

const (
	golangCi   = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	golines    = "github.com/segmentio/golines"
	gosec      = "github.com/securego/gosec/v2/cmd/gosec"
	addlicense = "github.com/google/addlicense"
	goimports  = "github.com/incu6us/goimports-reviser/v3"
)

func Lint() error {
	cmds := []func() error{GolangCiLint, LicenseCheck, Gosec, ProtoLint}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Run all formatters.
func Format() error {
	cmds := []func() error{Golines, License, GolangCiLintFix, GoImports, ProtoFormat}
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
			"\"write\"", "-project-name", "pkg.berachain.dev/stargazer", dir); err != nil {
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
	return goRun(gosec, "-exclude-generated", "./...")
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
