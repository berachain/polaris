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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keeper", func() {
	var (
		keeper       *keeper.Keeper
		ctx          sdk.Context
		genesisState types.GenesisState
	)

	JustBeforeEach(func() {
		ctx = sdk.Context{} // will have to update actual values here
		genesisState = types.GenesisState{}
	})

	DescribeTable("InitGenesis",
		func(ctx sdk.Context, genesisState types.GenesisState, expectedErr error) {
			err := keeper.InitGenesis(ctx, genesisState)
			Expect(err).To(Equal(expectedErr))
		},

		Entry("the genesis is valid", ctx, *types.DefaultGenesis(), nil),
		Entry("the GasLimit is invalid", ctx, genesisState, fmt.Errorf("gas limit mismatch: expected %d, got %d", ethGenesis.GasLimit, ctx.GasMeter().Limit())),
		Entry("the ChainID is invalid", ctx, genesisState, fmt.Errorf("invalid ChainID: 0")),
		Entry("the coinbase is invalid", ctx, genesisState, fmt.Errorf("invalid coinbase: ")),
		Entry("the timestamp is invalid", ctx, genesisState, fmt.Errorf("invalid timestamp: 0")),
		Entry("the balance is invalid", ctx, genesisState, fmt.Errorf("invalid balance: []")),
	)
})
