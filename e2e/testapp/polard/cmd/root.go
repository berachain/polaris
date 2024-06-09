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

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"cosmossdk.io/client/v2/autocli"
	clientv2keyring "cosmossdk.io/client/v2/autocli/keyring"
	"cosmossdk.io/core/address"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	evmv1alpha1 "github.com/berachain/polaris/cosmos/api/polaris/evm/v1alpha1"
	polarconfig "github.com/berachain/polaris/cosmos/config"
	ethcryptocodec "github.com/berachain/polaris/cosmos/crypto/codec"
	polarkeyring "github.com/berachain/polaris/cosmos/crypto/keyring"
	signinglib "github.com/berachain/polaris/cosmos/lib/signing"
	testapp "github.com/berachain/polaris/e2e/testapp"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtxconfig "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewRootCmd creates a new root command for simd. It is called once in the main function.
//

func NewRootCmd() *cobra.Command {
	var (
		autoCliOpts        autocli.AppOptions
		moduleBasicManager module.BasicManager
		clientCtx          client.Context
	)

	if err := depinject.Inject(
		depinject.Configs(
			testapp.MakeAppConfig(""),
			depinject.Supply(
				testapp.PolarisConfigFn(polarconfig.DefaultPolarisConfig()),
				testapp.QueryContextFn((&testapp.SimApp{})),
				log.NewNopLogger(),
				simtestutil.NewAppOptionsWithFlagHome(tempDir()),
			),
			depinject.Provide(
				signinglib.ProvideNoopGetSigners[*evmv1alpha1.WrappedEthereumTransaction],
				signinglib.ProvideNoopGetSigners[*evmv1alpha1.WrappedPayloadEnvelope],
				ProvideClientContext,
				ProvideKeyring,
			),
		),
		&autoCliOpts,
		&moduleBasicManager,
		&clientCtx,
	); err != nil {
		panic(err)
	}

	// Register `eth_secp256k1` algo.
	ethcryptocodec.RegisterInterfaces(clientCtx.InterfaceRegistry)

	rootCmd := &cobra.Command{
		Use:           "polard",
		Short:         "node daemon and CLI for interacting with a polaris node",
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			var err error
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			clientCtx = clientCtx.WithCmdContext(cmd.Context())
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err = config.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			customAppTemplate, customAppConfig := initAppConfig()

			return server.InterceptConfigsPreRunHandler(
				cmd, customAppTemplate, customAppConfig, polarconfig.RecommendedCometBFTConfig(),
			)
		},
	}

	initRootCmd(rootCmd, clientCtx.TxConfig,
		clientCtx.InterfaceRegistry, clientCtx.Codec, moduleBasicManager)

	if err := autoCliOpts.EnhanceRootCommand(rootCmd); err != nil {
		panic(err)
	}

	return rootCmd
}

func ProvideClientContext(
	appCodec codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
	txConfigOpts tx.ConfigOptions,
	legacyAmino *codec.LegacyAmino,
) client.Context {
	clientCtx := client.Context{}.
		WithCodec(appCodec).
		WithInterfaceRegistry(interfaceRegistry).
		WithLegacyAmino(legacyAmino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(testapp.DefaultNodeHome).
		WithKeyringOptions(polarkeyring.OnlyEthSecp256k1Option()).
		WithViper("") // In simapp, we don't use any prefix for env variables.

	// Read the config again to overwrite the default values with the values from the config file
	clientCtx, _ = config.ReadFromClientConfig(clientCtx)

	// re-create the tx config grpc instead of bank keeper
	txConfigOpts.TextualCoinMetadataQueryFn = authtxconfig.NewGRPCCoinMetadataQueryFn(clientCtx)
	txConfig, err := tx.NewTxConfigWithOptions(clientCtx.Codec, txConfigOpts)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithTxConfig(txConfig)

	return clientCtx
}

func ProvideKeyring(clientCtx client.Context, _ address.Codec) (clientv2keyring.Keyring, error) {
	kb, err := client.NewKeyringFromBackend(clientCtx, clientCtx.Keyring.Backend())
	if err != nil {
		return nil, err
	}

	return keyring.NewAutoCLIKeyring(kb)
}
