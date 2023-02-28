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
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/utils"
	offchain "pkg.berachain.dev/stargazer/store/offchain"
	"pkg.berachain.dev/stargazer/testutil"
)

var _ = Describe("Block Plugin", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		sk := testutil.EvmKey // testing key.
		p = utils.MustGetAs[*plugin](NewPlugin(offchain.NewFromDB(dbm.NewMemDB()), sk))
		p.Prepare(ctx)
	})

	It("should give the constant base fee", func() {
		Expect(p.BaseFee()).To(Equal(bf))
	})

	It("should get the header at current height", func() {
		header, _ := p.GetStargazerHeaderByNumber(ctx.BlockHeight())
		Expect(header.Hash()).To(Equal(header.Header.Hash()))
		Expect(header.TxHash).To(Equal(common.BytesToHash(ctx.BlockHeader().DataHash)))
	})

	It("should return empty header for non-existent height", func() {
		header, _ := p.GetStargazerHeaderByNumber(100000)
		Expect(*header).To(Equal(types.StargazerHeader{}))
	})

	It("should return header hash from context", func() {
		ctx = ctx.WithHeaderHash([]byte("test"))
		a := blockHashFromCosmosContext(ctx)
		Expect(a).To(Equal(common.BytesToHash([]byte("test"))))
	})
})
