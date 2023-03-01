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
package block

import (
	"math/big"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/store/offchain"
	"pkg.berachain.dev/stargazer/testutil"
)

var _ = Describe("Header", func() {
	ctx := testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
	sk := testutil.EvmKey // testing key.
	p := utils.MustGetAs[*plugin](NewPlugin(offchain.NewFromDB(dbm.NewMemDB()), sk))
	p.Prepare(ctx)
	qc := func(height int64, prove bool) (sdk.Context, error) {
		return ctx, nil
	}
	p.SetQueryContextFn(qc)

	It("set and get header", func() {
		ctx = ctx.WithBlockHeight(1).WithProposer(sdk.ConsAddress([]byte("test")))
		header := types.NewStargazerHeader(
			&types.Header{
				ParentHash:  common.Hash{0x01},
				UncleHash:   common.Hash{0x02},
				Coinbase:    common.Address{0x03},
				Root:        common.Hash{0x04},
				TxHash:      common.Hash{0x05},
				ReceiptHash: common.Hash{0x06},
				Number:      big.NewInt(1),
			},
			blockHashFromCosmosContext(ctx),
		)
		err := p.ProcessHeader(ctx, header)
		Expect(err).To(BeNil())

		header, err = p.GetStargazerHeaderByNumber(1)
		header.Extra = nil // Processed as nil.
		Expect(err).To(BeNil())

		// Expected header after all the values are set from the context.
		expectedHeader := p.fillHeader(ctx, header)
		Expect(*expectedHeader).To(Equal(*header))
	})
})
