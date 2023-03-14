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
	goInstall  = RunCmdV("go", "install", "-mod=readonly")
	goBuild    = RunCmdV("go", "build", "-mod=readonly")
	goRun      = RunCmdV("go", "run")
	goGenerate = RunCmdV("go", "generate")
	goModTidy  = RunCmdV("go", "mod", "tidy")
	goWorkSync = RunCmdV("go", "work", "sync")

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

	moduleDirs = []string{"contracts", "eth", "cosmos", "playground", "magefiles", "lib"}
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
		"./cosmos/cmd/" + cmd,
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
		"./playground/cmd/",
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
	return Build()
}

// Runs `go install` on the entire project.
func Install() error {
	PrintMageName()
	production = true
	statically = false

	args := []string{
		generateBuildTags(),
		generateLinkerFlags(production, statically),
		"./cmd/polard",
	}

	return goInstall(args...)
}

// Runs `go generate` on the entire project.
func Generate() error {
	PrintMageName()
	if err := goInstall(moq); err != nil {
		return err
	}
	if err := ExecuteForAllModules(moduleDirs, func(...string) error { return goGenerate("./...") }, false); err != nil {
		return err
	}
	return nil
}

// Runs `go generate` on the entire project and verifies that no files were
// changed.
func GenerateCheck() error {
	PrintMageName()
	if err := ExecuteForAllModules(moduleDirs, func(...string) error { return goGenerate("./...") }, false); err != nil {
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
