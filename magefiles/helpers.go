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
	"strings"

	"github.com/magefile/mage/sh"
)

var allPkgs, _ = sh.Output("go", "list", "pkg.berachain.dev/polaris/...")

// RunCmd is a helper function that returns a function that runs the given
// command with the given arguments.
func RunCmdV(cmd string, args ...string) func(args ...string) error {
	return func(args2 ...string) error {
		return sh.RunV(cmd, append(args, args2...)...)
	}
}

// RunOutput is a helper function that returns a function that runs the given
// command with the given arguments and returns the output.
func RunOutput(cmd string, args ...string) func(args ...string) (string, error) {
	return func(args2 ...string) (string, error) {
		return sh.Output(cmd, append(args, args2...)...)
	}
}

// GoListFilter returns a list of packages that match the given filter.
func GoListFilter(include bool, contains ...string) []string {
	return filter(strings.Split(allPkgs, "\n"), func(s string) bool {
		for _, c := range contains {
			if strings.Contains(s, c) {
				return include
			}
		}
		return !include
	})
}

// filter returns a new slice containing only the elements of ss that
// satisfy the predicate test.
func filter[T any](ss []T, test func(T) bool) []T {
	ret := make([]T, 0, len(ss))
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

// Executes a function in a given directory.
func ExecuteInDirectory(dir string, f func(args ...string) error, withArgs bool) error {
	rootCwd, _ := os.Getwd()
	// Change to the directory where the contracts are.
	if err := os.Chdir(dir); err != nil {
		return err
	}
	// Run the command
	if withArgs {
		if err := f(dir); err != nil {
			return err
		}
	} else {
		if err := f(); err != nil {
			return err
		}
	}

	// Go back to the starting directory.
	if err := os.Chdir(rootCwd); err != nil {
		return err
	}
	return nil
}

func ExecuteForAllModules(dirs []string, f func(args ...string) error, withArgs bool) error {
	for _, dir := range dirs {
		if err := ExecuteInDirectory(dir, f, withArgs); err != nil {
			return err
		}
	}
	return nil
}
