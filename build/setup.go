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

//nolint:forbidigo // its okay.
package main

import (
	"errors"
	"fmt"

	"github.com/carolynvs/magex/pkg"
	"github.com/magefile/mage/sh"
)

var (
	// tools.
	buf              = "github.com/bufbuild/buf/cmd/buf"
	gosec            = "github.com/cosmos/gosec/v2/cmd/gosec"
	golangcilint     = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	addlicense       = "github.com/google/addlicense"
	goimportsReviser = "github.com/incu6us/goimports-reviser/v3"
	moq              = "github.com/matryer/moq"
	ginkgo           = "github.com/onsi/ginkgo/v2/ginkgo"
	golines          = "github.com/segmentio/golines"
	goimports        = "golang.org/x/tools/cmd/goimports"
	rlpgen           = "github.com/ethereum/go-ethereum/rlp/rlpgen"

	allTools = []string{buf, gosec, golangcilint, addlicense,
		goimportsReviser, moq, ginkgo, golines, goimports, rlpgen}
)

// Setup runs the setup script for the current OS.
func main() {
	var err error

	// Ensure Mage is installed and available on the $PATH.
	if err = pkg.EnsureMage(""); err != nil {
		panic(err)
	}

	// Coming soon
	// // Run the setup script for the current OS.
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	err = setupMac()
	// case "linux":
	// 	err = setupLinux()
	// default:
	// 	err = fmt.Errorf("ngmi unsupported OS")
	// }

	if err = setupGoDeps(); err != nil {
		panic(err)
	}
}

func setupGoDeps() error {
	for _, tool := range allTools {
		fmt.Println("Installing ", tool)
		if err := sh.RunCmd("go", "install", "-mod=readonly", tool); err() != nil {
			return errors.New("failed to install " + tool + ": " + err().Error())
		}
	}
	fmt.Println("\n==============================================================")
	fmt.Println("Tools installed successful! Ensure $GOPATH/bin is on your $PATH!")
	fmt.Println("==============================================================")
	return nil
}

// // setupMac runs the setup script for macOS.
// func setupMac() error {
// 	return fmt.Errorf("mac setup coming soon")
// }

// // setupLinux runs the setup script for Linux.
// func setupLinux() error {
// 	return fmt.Errorf("linux setup coming soon")
// }
