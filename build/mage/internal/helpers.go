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
func ExecuteInDirectory(dir string, f func(args ...string) error) error {
	rootCwd, _ := os.Getwd()
	// Change to the directory where the contracts are.
	if err := os.Chdir(dir); err != nil {
		return err
	}
	// Run the forge command.
	if err := f(); err != nil {
		return err
	}

	// Go back to the starting directory.
	if err := os.Chdir(rootCwd); err != nil {
		return err
	}
	return nil
}
