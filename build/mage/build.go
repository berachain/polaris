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
	"os"

	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"

	mi "pkg.berachain.dev/polaris/build/mage/internal"
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

	// Dependencies.
	moq = "github.com/matryer/moq"

	// Variables and Helpers.
	production = false
	statically = false
)

// Runs a series of commonly used commands.
func All() error {
	cmds := []func() error{ForgeBuild, Generate, Proto, Format, Lint, BuildPolarisApp, Test, TestIntegration}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

func BuildPolarisApp() error {
	cmd := "polard"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./pkg/cosmos/cmd/" + cmd,
	}
	return goBuild(args...)
}

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

	if err = BuildPolarisApp(); err != nil {
		return err
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
	PrintMageName()
	production = true
	statically = false

	// Verify dependencies.
	if err := goModVerify(); err != nil {
		return err
	}

	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"./pkg/cosmos/cmd/polard",
	}

	return goInstall(args...)
}

// Runs `go generate` on the entire project.
func Generate() error {
	PrintMageName()
	if err := goInstall(moq); err != nil {
		return err
	}
	return goGenerate("-x", "./...")
}

// Runs `go generate` on the entire project and verifies that no files were
// changed.
func GenerateCheck() error {
	PrintMageName()
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
