package block

import (
	"math/big"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/store/offchain"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

var _ = Describe("Genesis", func() {
	var (
		ctx sdk.Context
		p   Plugin
	)

	var (
		header = &coretypes.StargazerHeader{
			Header: &coretypes.Header{
				ParentHash:      common.HexToHash("0x123"),
				UncleHash:       common.HexToHash("0x123"),
				Coinbase:        common.HexToAddress("0x123"),
				Root:            common.HexToHash("0x123"),
				TxHash:          common.HexToHash("0x123"),
				ReceiptHash:     common.HexToHash("0x123"),
				Bloom:           coretypes.BytesToBloom([]byte("0x123")),
				Difficulty:      big.NewInt(1),
				Number:          big.NewInt(1),
				GasLimit:        1,
				GasUsed:         1,
				Time:            1,
				Extra:           []byte("0x123"),
				MixDigest:       common.HexToHash("0x123"),
				Nonce:           coretypes.BlockNonce{0x1},
				BaseFee:         big.NewInt(1),
				WithdrawalsHash: &common.Hash{0x1},
			},
		}
	)

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		sk := testutil.EvmKey // testing key.
		p = NewPlugin(offchain.NewFromDB(dbm.NewMemDB()), sk)
		p.Prepare(ctx)
	})

	It("Init genesis", func() {
		genesis := types.DefaultGenesis()
		genesis.Headers = make([][]byte, 0)
		bz, err := header.MarshalBinary()
		Expect(err).ToNot(HaveOccurred())
		genesis.Headers = append(genesis.Headers, bz)
		p.InitGenesis(ctx, genesis)

		// Check that the header is set correctly.
		p.Prepare(ctx)
		h := p.GetStargazerHeaderByNumber(1)
		Expect(h).ToNot(BeNil())
		Expect(h.BaseFee).To(Equal(header.BaseFee))
		Expect(h.Bloom).To(Equal(header.Bloom))
		Expect(h.Coinbase).To(Equal(header.Coinbase))
	})

	It("Export genesis", func() {
		p.SetStargazerHeader(ctx, header)
		var genesis types.GenesisState
		p.ExportGenesis(ctx, &genesis)
		Expect(genesis).ToNot(BeNil())

		// Check that the headers are exported correctly.
		a := make([]coretypes.StargazerHeader, 0)
		for _, bz := range genesis.Headers {
			var h coretypes.StargazerHeader
			h.UnmarshalBinary(bz)
			a = append(a, h)
		}
		Expect(a).To(HaveLen(1))
		Expect(a[0].BaseFee).To(Equal(header.BaseFee))
		Expect(a[0].Bloom).To(Equal(header.Bloom))
		Expect(a[0].Coinbase).To(Equal(header.Coinbase))
	})
})
