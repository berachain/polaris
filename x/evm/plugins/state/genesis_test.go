package state

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/testutil"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

var (
	alice = testutil.Alice
)

var _ = Describe("Genesis", func() {
	var (
		ctx sdk.Context
		sp  Plugin
	)

	BeforeEach(func() {
		var ak AccountKeeper
		var bk BankKeeper
		ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		sp = NewPlugin(ak, bk, testutil.EvmKey, "abera", nil)
		sp.InitGenesis(ctx, &types.GenesisState{
			CodeRecords: []types.CodeRecord{
				{
					Address: alice.Hex(),
					Code:    []byte("code"),
				},
			},
			StateRecords: []types.StateRecord{
				{
					Address: alice.Hex(),
					Slot:    []byte("slot"),
					Value:   []byte("value"),
				},
			},
		})
	})

	It("should export current state", func() {
		var gs types.GenesisState
		sp.ExportGenesis(ctx, &gs)

		Expect(gs.CodeRecords).To(HaveLen(1))
		Expect(gs.CodeRecords[0].Address).To(Equal(alice.Hex()))
	})
})
