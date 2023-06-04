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

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/testing"
	"pkg.berachain.dev/polaris/cosmos/precompile/staking"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmmempool "pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/utils"

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
		sk           stakingkeeper.Keeper
		ctx          sdk.Context
		sc           ethprecompile.StatefulImpl
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
			Gas:      10000000000,
			Data:     []byte("abcdef"),
			GasPrice: big.NewInt(2 ^ 63), // overpaying so test doesn't fail due to EIP-1559 math.
		}

		// before chain, init genesis state
		ctx, ak, _, sk = testutil.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			ak, sk,
			storetypes.NewKVStoreKey("evm"),
			"authority",
			evmmempool.NewWrappedGethTxPool(),
			func() *ethprecompile.Injector {
				return ethprecompile.NewPrecompiles([]ethprecompile.Registrable{sc}...)
			},
		)
		for _, plugin := range k.GetHost().GetAllPlugins() {
			plugin, hasInitGenesis := utils.GetAs[plugins.HasGenesis](plugin)
			if hasInitGenesis {
				plugin.InitGenesis(ctx, core.DefaultGenesis)
			}
		}
		validator, err := NewValidator(sdk.ValAddress(valAddr), PKs[0])
		Expect(err).ToNot(HaveOccurred())
		validator.Status = stakingtypes.Bonded
		sk.SetValidator(ctx, validator)
		sc = staking.NewPrecompileContract(&sk)
		k.Setup(storetypes.NewKVStoreKey("offchain-evm"), nil, "", GinkgoT().TempDir(), log.NewNopLogger())
		_ = sk.SetParams(ctx, stakingtypes.DefaultParams())

		// Set validator with consensus address.
		consAddr, err := validator.GetConsAddr()
		Expect(err).ToNot(HaveOccurred())
		err = sk.SetValidatorByConsAddr(ctx, validator)
		Expect(err).ToNot(HaveOccurred())

		// Set header's consensus address to match the validator's.
		header := ctx.BlockHeader()
		header.ProposerAddress = consAddr.Bytes()
		ctx = ctx.WithBlockHeader(header)

		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(100000000000000)).
			WithKVGasConfig(storetypes.GasConfig{}).
			WithBlockHeight(1)
		err = k.BeginBlocker(ctx)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("New Block", func() {
		BeforeEach(func() {
			// before every tx
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
		})

		AfterEach(func() {
			err := k.EndBlock(ctx)
			Expect(err).ToNot(HaveOccurred())
			err = os.RemoveAll("tmp/berachain")
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
			legacyTxData.GasPrice = big.NewInt(10000000000)
			tx := coretypes.MustSignNewTx(key, signer, legacyTxData)
			addr, err := signer.Sender(tx)
			Expect(err).ToNot(HaveOccurred())
			k.GetHost().GetStatePlugin().Reset(ctx)
			k.GetHost().GetStatePlugin().CreateAccount(addr)
			k.GetHost().GetStatePlugin().AddBalance(addr, (&big.Int{}).Mul(big.NewInt(9000000000000000000), big.NewInt(999)))
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
