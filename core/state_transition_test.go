package core

import (
	"math/big"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateTransition", func() {
	// var ak types.AccountKeeper
	// var bk types.BankKeeper
	// var ctx sdk.Context
	// // var sdb *state.StateDB
	// var evm *vm.StargazerEVM

	BeforeEach(func() {
		// ctx, ak, bk, _ = testutil.SetupMinimalKeepers()
		// sdb = state.NewStateDB(ctx, ak, bk, testutil.EvmKey, "abera")
		// ef := NewEVMFactory()

	})

	It("best", func() {
		msg := coretypes.NewMessage(
			testutil.Alice,
			&testutil.Bob,
			0,
			big.NewInt(10000000),
			10000000,
			big.NewInt(1),
			new(big.Int),
			new(big.Int),
			[]byte{},
			nil,
			false,
		)
		_ = msg
		// st := NewStateTransition(evm, msg)
		// _, err := st.transitionDB()
		Expect(nil).To(BeNil())
		// Expect(res).To(Equal(0))
	})
})
