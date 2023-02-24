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
	offchain "pkg.berachain.dev/stargazer/store/offchain"
	"pkg.berachain.dev/stargazer/testutil"
)

var _ = Describe("Header", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		sk := testutil.EvmKey // testing key.
		p = utils.MustGetAs[*plugin](NewPlugin(offchain.NewFromDB(dbm.NewMemDB()), sk))
		p.Prepare(ctx)
	})

	It("set and get header", func() {
		header := types.NewStargazerHeader(
			&types.Header{
				ParentHash:  common.Hash{0x01},
				UncleHash:   common.Hash{0x02},
				Coinbase:    common.Address{0x03},
				Root:        common.Hash{0x04},
				TxHash:      common.Hash{0x05},
				ReceiptHash: common.Hash{0x06},
				Number:      big.NewInt(10),
			},
			common.Hash{0x01},
		)
		p.SetStargazerHeader(ctx, header)

		header2, found := p.GetStargazerHeader(ctx, 10)
		Expect(found).To(BeTrue())
		Expect(header2.Hash()).To(Equal(header.Hash()))

		// get unknown header
		header3, found := p.GetStargazerHeader(ctx, 11)
		Expect(found).To(BeFalse())
		Expect(header3).To(BeNil())
	})

	It("should be able to prune headers", func() {
		header := types.NewStargazerHeader(
			&types.Header{
				ParentHash:  common.Hash{0x01},
				UncleHash:   common.Hash{0x02},
				Coinbase:    common.Address{0x03},
				Root:        common.Hash{0x04},
				TxHash:      common.Hash{0x05},
				ReceiptHash: common.Hash{0x06},
				Number:      big.NewInt(10),
			},
			common.Hash{0x01},
		)
		p.SetStargazerHeader(ctx, header)

		// Prune header.
		p.PruneStargazerHeader(ctx, header)

		// Get header.
		header2, found := p.GetStargazerHeader(ctx, 10)
		Expect(found).To(BeFalse())
		Expect(header2).To(BeNil())
	})

	It("should be able to track the headers", func() {
		for i := 1; i <= 260; i++ {
			ctx := ctx.WithBlockHeight(int64(i))
			header := types.NewStargazerHeader(
				&types.Header{Number: big.NewInt(int64(i))}, common.Hash{0x01})
			p.SetStargazerHeader(ctx, header)
		}

		// Run TrackHistoricalStargazerHeader on the header with height 260.
		ctx := ctx.WithBlockHeight(260)
		p.TrackHistoricalStargazerHeader(ctx, types.NewStargazerHeader(
			&types.Header{Number: big.NewInt(260)}, common.Hash{0x01}))

		// Get the header with height 1.
		_, found := p.GetStargazerHeader(ctx, 1)
		Expect(found).To(BeFalse())

		// Get the header with height 10.
		_, found = p.GetStargazerHeader(ctx, 10)
		Expect(found).To(BeTrue())
	})
})
