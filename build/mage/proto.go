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

	mi "pkg.berachain.dev/polaris/build/mage/internal"
)

var (
	// Buf Commands.
	bufRepo = "github.com/bufbuild/buf/cmd/buf"
	// bufBuild  = mi.RunCmdV("go", "run", bufRepo, "build").
	bufFormat = mi.RunCmdV("go", "run", bufRepo, "format", "-w")
	bufLint   = mi.RunCmdV("go", "run", bufRepo, "lint", "--error-format=json")

	// Docker Args
	// TODO: remove once https://github.com/cosmos/cosmos-sdk/pull/13960 is merged
	protoImageName    = "ghcr.io/cosmos/proto-builder"
	protoImageVersion = "0.12.0"
	protoDir          = "proto"
)

func dockerRunProtoImage(pwd string) func(args ...string) error {
	return mi.RunCmdV("docker",
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
		"sh", "./build/scripts/proto/proto_generate.sh",
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
