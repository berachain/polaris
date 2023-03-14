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
	"fmt"
	"os"
)

var (
	// Buf Commands.
	bufRepo = "github.com/bufbuild/buf/cmd/buf"
	// bufBuild  = RunCmdV("go", "run", bufRepo, "build").
	bufFormat = RunCmdV("go", "run", bufRepo, "format", "-w")
	bufLint   = RunCmdV("go", "run", bufRepo, "lint", "--error-format=json")

	// Docker Args
	// TODO: remove once https://github.com/cosmos/cosmos-sdk/pull/13960 is merged
	protoImageName    = "ghcr.io/cosmos/proto-builder"
	protoImageVersion = "0.12.1"
	protoDir          = "cosmos/proto"
)

func dockerRunProtoImage(pwd string) func(args ...string) error {
	return RunCmdV("docker",
		"run", "--rm", "-v", pwd+":/workspace",
		"--workdir", "/workspace",
		protoImageName+":"+protoImageVersion)
}

// Run all proto related targets.
func Proto() error {
	cmds := []func() error{ProtoFormat, ProtoLint, ProtoGen}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Generate protobuf source files.
func ProtoGen() error {
	PrintMageName()
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return dockerRunProtoImage(dir)(
		"sh", "./cosmos/build/scripts/proto_generate.sh",
	)
}

// Check that the generated protobuf source files are up to date.
func ProtoGenCheck() error {
	PrintMageName()
	if err := ProtoGen(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Format .proto files.
func ProtoFormat() error {
	PrintMageName()
	return bufWrapper(bufFormat)
}

// Lint .proto files.
func ProtoLint() error {
	PrintMageName()
	return bufWrapper(bufLint)
}

// Wraps buf commands with the proper directory change.
func bufWrapper(bufFunc func(args ...string) error) error {
	rootCwd, _ := os.Getwd()
	// Change to the directory where the *.proto's are.
	if err := os.Chdir(protoDir); err != nil {
		return err
	}
	// Run the buf command.
	if err := bufFunc(); err != nil {
		return err
	}
	// Go back to the starting directory.
	if err := os.Chdir(rootCwd); err != nil {
		return err
	}
	return nil
}
