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

//nolint:mnd // from sdk.
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/spf13/viper"

	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testnet command", func() {
	var home string
	var dir *os.File
	BeforeEach(func() {
		home = os.TempDir()
		dir = os.NewFile(0, os.DevNull)
		v, err := dir.ReadDir(-1)
		if err == nil {
			Expect(dir.Seek(0, 0)).To(BeNil())
		}
		if len(v) == 0 {
			os.RemoveAll(home)
		}
	})

	It("should generate a testnet with given options", func() {
		encodingConfig := moduletestutil.MakeTestEncodingConfig(staking.AppModuleBasic{},
			auth.AppModuleBasic{})
		logger := log.NewNopLogger()
		cfg, err := genutiltest.CreateDefaultCometConfig(home)
		Expect(err).To(BeNil())
		err = genutiltest.ExecInitCmd(simapp.ModuleBasics, home, encodingConfig.Codec)
		Expect(err).To(BeNil())

		serverCtx := server.NewContext(viper.New(), cfg, logger)
		clientCtx := client.Context{}.
			WithCodec(encodingConfig.Codec).
			WithHomeDir(home).
			WithTxConfig(encodingConfig.TxConfig)

		ctx := context.Background()
		ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
		ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
		cmd := testnetInitFilesCmd(simapp.ModuleBasics, banktypes.GenesisBalancesIterator{})
		cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend),
			fmt.Sprintf("--output-dir=%s", home)})
		err = cmd.ExecuteContext(ctx)
		Expect(err).To(BeNil())

		genFile := cfg.GenesisFile()
		appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
		Expect(err).To(BeNil())

		bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
		Expect(bankGenState.Supply.IsZero()).To(BeFalse())
	})
})
