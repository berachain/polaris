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
	"math/big"
	"os"

	storetypes "cosmossdk.io/store/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/precompile/staking"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func NewValidator(operator sdk.ValAddress, pubKey cryptotypes.PubKey) (stakingtypes.Validator, error) {
	return stakingtypes.NewValidator(operator, pubKey, stakingtypes.Description{})
}

var (
	PKs = simtestutil.CreateTestPubKeys(500)
)

var _ = Describe("Processor", func() {
	var (
		k            *keeper.Keeper
		ak           state.AccountKeeper
		bk           state.BankKeeper
		sk           stakingkeeper.Keeper
		ctx          sdk.Context
		sc           precompile.StatefulImpl
		key, _       = crypto.GenerateEthKey()
		signer       = coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
		legacyTxData *coretypes.LegacyTx
		valAddr      = common.Address{0x21}.Bytes()
	)

	BeforeEach(func() {
		err := os.RemoveAll("tmp/berachain")
		Expect(err).ToNot(HaveOccurred())

		legacyTxData = &coretypes.LegacyTx{
			Nonce:    0,
			Gas:      10000000,
			Data:     []byte("abcdef"),
			GasPrice: big.NewInt(1),
		}

		// before chain, init genesis state
		ctx, ak, bk, sk = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("evm"),
			ak, bk,
			"authority",
			simtestutil.NewAppOptionsWithFlagHome("tmp/berachain"),
			evmmempool.NewEthTxPoolFrom(sdkmempool.NewPriorityMempool(
				sdkmempool.DefaultPriorityNonceMempoolConfig()),
			),
		)
		validator, err := NewValidator(sdk.ValAddress(valAddr), PKs[0])
		Expect(err).ToNot(HaveOccurred())
		validator.Status = stakingtypes.Bonded
		sk.SetValidator(ctx, validator)
		sc = staking.NewPrecompileContract(&sk)
		k.Setup(ak, bk, []precompile.Registrable{sc}, nil, "", GinkgoT().TempDir())
		k.ConfigureGethLogger(ctx)
		_ = sk.SetParams(ctx, stakingtypes.DefaultParams())
		for _, plugin := range k.GetHost().GetAllPlugins() {
			plugin.InitGenesis(ctx, types.DefaultGenesis())
		}

		// before every block
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(100000000000000)).
			WithKVGasConfig(storetypes.GasConfig{}).
			WithBlockHeight(1)
		k.BeginBlocker(ctx)
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		AfterEach(func() {
			k.Precommit(ctx)
			err := os.RemoveAll("tmp/berachain")
			Expect(err).ToNot(HaveOccurred())
		})

		It("should panic on nil, empty transaction", func() {
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, nil)
				Expect(err).To(HaveOccurred())
			}).To(Panic())
			Expect(func() {
				_, err := k.ProcessTransaction(ctx, &coretypes.Transaction{})
				Expect(err).To(HaveOccurred())
			}).To(Panic())
		})

		It("should successfully deploy a valid contract and call it", func() {
			legacyTxData.Data = common.FromHex(bindings.SolmateERC20Bin)
			tx := coretypes.MustSignNewTx(key, signer, legacyTxData)
			addr, err := signer.Sender(tx)
			Expect(err).ToNot(HaveOccurred())
			k.GetHost().GetStatePlugin().CreateAccount(addr)
			k.GetHost().GetStatePlugin().AddBalance(addr, big.NewInt(1000000000))
			k.GetHost().GetStatePlugin().Finalize()

			// create the contract
			result, err := k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())
			// call the contract non-view function
			deployAddress := crypto.CreateAddress(crypto.PubkeyToAddress(key.PublicKey), 0)
			legacyTxData.To = &deployAddress
			var solmateABI abi.ABI
			err = solmateABI.UnmarshalJSON([]byte(bindings.SolmateERC20ABI))
			Expect(err).ToNot(HaveOccurred())
			input, err := solmateABI.Pack("mint", common.BytesToAddress([]byte{0x88}), big.NewInt(8888888))
			Expect(err).ToNot(HaveOccurred())
			legacyTxData.Data = input
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			result, err = k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())

			// call the contract view function
			legacyTxData.Data = crypto.Keccak256Hash([]byte("totalSupply()")).Bytes()[:4]
			legacyTxData.Nonce++
			tx = coretypes.MustSignNewTx(key, signer, legacyTxData)
			result, err = k.ProcessTransaction(ctx, tx)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Err).ToNot(HaveOccurred())
		})
	})
})
