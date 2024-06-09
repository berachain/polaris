// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cmd_test

import (
	"fmt"
	"os"
	"testing"

	testapp "github.com/berachain/polaris/e2e/testapp"
	"github.com/berachain/polaris/e2e/testapp/polard/cmd"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e/testapp/polard/cmd")
}

var _ = Describe("Init command", func() {
	It("should initialize the app with given options", func() {
		stdout := os.Stdout
		defer func() { os.Stdout = stdout }()
		os.Stdout = os.NewFile(0, os.DevNull)
		rootCmd := cmd.NewRootCmd()
		rootCmd.SetArgs([]string{
			"init",        // Test the init cmd
			"simapp-test", // Moniker
			fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json
		})

		err := svrcmd.Execute(rootCmd, "", testapp.DefaultNodeHome)
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("Home flag registration", func() {
	It("should set home directory correctly", func() {
		// Redirect standard out to null
		stdout := os.Stdout
		defer func() { os.Stdout = stdout }()
		os.Stdout = os.NewFile(0, os.DevNull)
		homeDir := os.TempDir()

		rootCmd := cmd.NewRootCmd()
		rootCmd.SetArgs([]string{
			"query",
			fmt.Sprintf("--%s", flags.FlagHome),
			homeDir,
		})

		err := svrcmd.Execute(rootCmd, "", testapp.DefaultNodeHome)
		Expect(err).ToNot(HaveOccurred())

		result, err := rootCmd.Flags().GetString(flags.FlagHome)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(homeDir))
	})
})
