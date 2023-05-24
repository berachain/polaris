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

package keeper_test

import (
	"fmt"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"

	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keeper", func() {
	var (
		k            *keeper.Keeper
		ak           state.AccountKeeper
		bk           state.BankKeeper
		sk           stakingkeeper.Keeper
		sc           ethprecompile.StatefulImpl
		ctx          sdk.Context
		genesisState *types.GenesisState
		err          error
	)

	BeforeEach(func() {
		// setup keepers for genesis
		ctx, ak, bk, sk = testutil.SetupMinimalKeepers()
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(30000000))

		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("evm"),
			ak, bk, sk,
			"authority",
			simtestutil.NewAppOptionsWithFlagHome("tmp/berachain"),
			evmmempool.NewEthTxPoolFrom(evmmempool.DefaultPriorityMempool()),
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
		)

		lib.MintCoinsToAddress(ctx, bk, types.ModuleName, testutil.Alice, "abera", big.NewInt(69000))

		genesisState = types.DefaultGenesis()
	})

	Context("InitGenesis is called", func() {
		JustBeforeEach(func() {
			err = k.InitGenesis(ctx, *genesisState)
		})

		When("the genesis is valid", func() {
			It("should execute without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
		When("the coinbase is invalid", func() {
			var invalidCoinbase string
			BeforeEach(func() {
				// TODO: find a way to change the coinbase programmatically
				// this is so bad but it works so....
				invalidCoinbase = "0x0000000000000000000000000000000000000001"
				genesisState.Params.EthGenesis = fmt.Sprintf("{\"config\":{\"chainId\":69420,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"berlinBlock\":0,\"londonBlock\":0,\"arrowGlacierBlock\":0,\"grayGlacierBlock\":0,\"mergeNetsplitBlock\":0,\"shanghaiTime\":0,\"terminalTotalDifficulty\":0,\"terminalTotalDifficultyPassed\":true},\"nonce\":\"0x45\",\"timestamp\":\"0x0\",\"extraData\":\"0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa\",\"gasLimit\":\"0x1c9c380\",\"difficulty\":\"0x45\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"%s\",\"alloc\":{},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"baseFeePerGas\":null}", invalidCoinbase)
			})
			It("should report a coinbase mismatch error", func() {
				Expect(err).To(Equal(fmt.Errorf("coinbase of the genesis block must be the null address, not: %s", invalidCoinbase)))
			})
		})
		When("the balance is invalid", func() {
			var (
				invalidAddress string
				invalidBalance string
			)
			BeforeEach(func() {
				// TODO: find a way to change the balance programmatically
				// ethGenesis := enclib.MustUnmarshalJSON[core.Genesis]([]byte(genesisState.Params.EthGenesis))
				// ethGenesis.Alloc[testutil.Bob] = core.GenesisAccount{
				// 	Balance: big.NewInt(100),
				// }
				// genesisState.Params.EthGenesis =
				invalidAddress = "0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"
				invalidBalance = "0x09184e72a000"
				genesisState.Params.EthGenesis = fmt.Sprintf("{\"config\":{\"chainId\":69420,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"berlinBlock\":0,\"londonBlock\":0,\"arrowGlacierBlock\":0,\"grayGlacierBlock\":0,\"mergeNetsplitBlock\":0,\"shanghaiTime\":0,\"terminalTotalDifficulty\":0,\"terminalTotalDifficultyPassed\":true},\"nonce\":\"0x45\",\"timestamp\":\"0x0\",\"extraData\":\"0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa\",\"gasLimit\":\"0x1c9c380\",\"difficulty\":\"0x45\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"%s\": {\"balance\":\"%s\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"baseFeePerGas\":null}", invalidAddress, invalidBalance)
			})
			It("should report a balance mismatch error", func() {
				balance, _ := new(big.Int).SetString(invalidBalance, 0) // convert invalidBalance to big.Int
				Expect(err).To(Equal(fmt.Errorf("account %s balance mismatch: expected 0, got %v", invalidAddress, balance)))
			})
		})
	})
})
