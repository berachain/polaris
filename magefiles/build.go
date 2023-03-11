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
//
//nolint:forbidigo // its okay.
package main

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

var (
	// Commands.
	goInstall   = RunCmdV("go", "install", "-mod=readonly")
	goBuild     = RunCmdV("go", "build", "-mod=readonly")
	goRun       = RunCmdV("go", "run")
	goGenerate  = RunCmdV("go", "generate")
	goModVerify = RunCmdV("go", "mod", "verify")
	goModTidy   = RunCmdV("go", "mod", "tidy")
	goWorkSync  = RunCmdV("go", "work", "sync")

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

	moduleDirs = []string{"eth", "host/cosmos", "host/playground", "magefiles", "lib"}
)

// Runs a series of commonly used commands.
func All() error {
	cmds := []func() error{ForgeBuild, Generate, Proto, Format, Lint,
		BuildPolarisCosmosApp, BuildPolarisPlaygroundApp, TestUnit, TestIntegration}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Runs `go build` on the cosmos app.
func BuildPolarisCosmosApp() error {
	cmd := "polard"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./host/cosmos/cmd/" + cmd,
	}
	fmt.Println(color.Ize(color.Yellow, "Building Cosmos app..."))
	return goBuild(args...)
}

// Runs `go build` on the playground app.
func BuildPolarisPlaygroundApp() error {
	cmd := "playground"
	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"-o", generateOutDirectory(cmd),
		"./host/playground/cmd/",
	}
	fmt.Println(color.Ize(color.Yellow, "Building Playground app..."))
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

	// Build the cosmos app
	if err = BuildPolarisCosmosApp(); err != nil {
		return err
	}

	// Build the playground app
	if err = BuildPolarisPlaygroundApp(); err != nil {
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
		"./host/cosmos/cmd/polard",
	}

	return goInstall(args...)
}

// Runs `go generate` on the entire project.
func Generate() error {
	PrintMageName()
	if err := goInstall(moq); err != nil {
		return err
	}
	if err := ExecuteForAllModules(moduleDirs, goGenerate, true); err != nil {
		return err
	}
	return nil
}

// Runs `go generate` on the entire project and verifies that no files were
// changed.
func GenerateCheck() error {
	PrintMageName()
	if err := ExecuteForAllModules(moduleDirs, goGenerate, true); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Runs 'go tidy' on the entire project.
func Tidy() error {
	return ExecuteForAllModules(moduleDirs, goModTidy, false)
}

// Runs 'go work sync' on the entire project.
func Sync() error {
	return goWorkSync()
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
