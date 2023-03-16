// SPDX-License-Identifier: MIT
//
// # Copyright (c) 2023 Berachain Foundation
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

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	protoImageName    = "ghcr.io/cosmos/proto-builder"
	protoImageVersion = "0.12.1"
	protoDir          = "cosmos/proto"

	bufCommand = sh.RunCmd("buf")
)

type Proto mg.Namespace

// Run all proto related targets.
func (Proto) All() {
	mg.SerialDeps(Proto.Format, Proto.Lint, Proto.Gen)
}

// Generate protobuf source files.
func (Proto) Gen() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dockerArgs := []string{
		"run", "--rm", "-v", dir + ":/workspace",
		"--workdir", "/workspace",
		protoImageName + ":" + protoImageVersion,
		"sh", "./cosmos/build/scripts/proto_generate.sh",
	}

	return sh.Run("docker", dockerArgs...)
}

// Check that the generated protobuf source files are up to date.
func (Proto) GenCheck() error {
	mg.Deps(Proto.Gen)
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Format .proto files.
func (Proto) Format() error {
	return bufCommand("format", "-w", protoDir)
}

// Lint .proto files.
func (Proto) Lint() error {
	return bufCommand("lint", "--error-format=json", protoDir)
}
