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
	"errors"
	"math/big"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth/core"
	"github.com/berachain/polaris/eth/core/types"
	"github.com/berachain/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Header", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext(
			log.NewTestLogger(GinkgoT())).
			WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		p = utils.MustGetAs[*plugin](NewPlugin(testutil.EvmKey,
			func() func(height int64, prove bool) (sdk.Context, error) { return mockQueryContext }))
		p.Prepare(ctx) // on block 0 (genesis)
	})

	It("should handle genesis header", func() {
		header := &ethtypes.Header{
			Number: big.NewInt(0),
		}
		genHash := header.Hash()
		Expect(p.StoreHeader(header)).ToNot(HaveOccurred())

		genHeadByNum, err := p.getGenesisHeader()
		Expect(err).NotTo(HaveOccurred())
		Expect(genHeadByNum.Hash()).To(Equal(genHash))

		genHeadByHash, err := p.GetHeaderByHash(genHash)
		Expect(err).NotTo(HaveOccurred())
		Expect(genHeadByHash.Hash()).To(Equal(genHash))
	})

	It("set and get header", func() {
		// we are on block 10
		p.Prepare(ctx.WithBlockHeight(10))

		// just finished processing block 10
		header := &ethtypes.Header{
			ParentHash:  common.Hash{0x01},
			UncleHash:   common.Hash{0x02},
			Coinbase:    common.Address{0x03},
			Root:        common.Hash{0x04},
			TxHash:      common.Hash{0x05},
			ReceiptHash: common.Hash{0x06},
			Number:      big.NewInt(10),
		}
		err := p.StoreHeader(header)
		Expect(err).ToNot(HaveOccurred())

		// now querying
		header2, err := p.GetHeaderByNumber(10)
		Expect(err).ToNot(HaveOccurred())
		Expect(header2.Hash()).To(Equal(header.Hash()))

		// get unknown header should return the latest header (10)
		header3, err := p.GetHeaderByNumber(11)
		Expect(err).ToNot(HaveOccurred())
		Expect(header3.Hash()).To(Equal(header.Hash()))
	})

	It("should be able to prune headers", func() {
		toAdd := int64(prevHeaderHashes + 5) // the first 5 hashes will eventually get deleted
		var deletedHashes []common.Hash
		for i := int64(1); i <= toAdd; i++ {
			ctx = ctx.WithBlockHeight(i)
			p.Prepare(ctx)
			header := generateHeaderAtHeight(i)
			if i <= 5 { // these hashes will be deleted
				deletedHashes = append(deletedHashes, header.Hash())
			}
			err := p.StoreHeader(header)
			Expect(err).ToNot(HaveOccurred())
		}

		// Get the header with height 1.
		_, err := p.GetHeaderByNumber(1)
		Expect(err).ToNot(HaveOccurred())

		// Get the header with height 10.
		_, err = p.GetHeaderByNumber(uint64(toAdd))
		Expect(err).ToNot(HaveOccurred())

		// check to see that the hashes are actually pruned
		// these 5 hashes will not be found because we only store last prevHeaderHashes (256)
		for _, deletedHash := range deletedHashes {
			var deletedHeader *ethtypes.Header
			deletedHeader, err = p.GetHeaderByHash(deletedHash)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(core.ErrHeaderNotFound))
			Expect(deletedHeader).To(BeNil())
		}
	})
})

func mockQueryContext(height int64, _ bool) (sdk.Context, error) {
	if height <= 0 {
		return sdk.Context{}, errors.New("cannot query context at this height")
	}
	ctx := testutil.NewContext(log.NewTestLogger(GinkgoT())).WithBlockHeight(height)
	header := generateHeaderAtHeight(height)
	headerBz, err := types.MarshalHeader(header)
	if err != nil {
		return sdk.Context{}, err
	}
	ctx.MultiStore().GetKVStore(testutil.EvmKey).Set([]byte{evmtypes.HeaderKey}, headerBz)
	ctx.MultiStore().GetKVStore(testutil.EvmKey).Set(header.Hash().Bytes(), header.Number.Bytes())
	return ctx, nil
}

func generateHeaderAtHeight(height int64) *ethtypes.Header {
	return &ethtypes.Header{
		ParentHash:  common.Hash{0x01},
		UncleHash:   common.Hash{0x02},
		Coinbase:    common.Address{0x03},
		Root:        common.Hash{0x04},
		TxHash:      common.Hash{0x05},
		ReceiptHash: common.Hash{0x06},
		Number:      big.NewInt(height),
	}
}
