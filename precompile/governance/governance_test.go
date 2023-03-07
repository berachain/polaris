package governance

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/polaris/lib/utils"
	polarisprecompile "pkg.berachain.dev/polaris/precompile"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	testutil "pkg.berachain.dev/polaris/testing/utils"
	evmutils "pkg.berachain.dev/polaris/x/evm/utils"
)

// Test Reporter to use governance module tests with Ginkgo.
type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func fundAccount(ctx sdk.Context, bk bankkeeper.Keeper, acc sdk.AccAddress, coins sdk.Coins) {
	if err := bk.MintCoins(ctx, governancetypes.ModuleName, coins); err != nil {
		panic(err)
	}
	if err := bk.SendCoinsFromModuleToAccount(ctx, governancetypes.ModuleName, acc, coins); err != nil {
		panic(err)
	}
}

// Helper functions for setting up the tests.
func setup(ctrl *gomock.Controller, caller sdk.AccAddress) (
	sdk.Context,
	bankkeeper.Keeper,
	*governancekeeper.Keeper,
) {
	// Setup the keepers and context.
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Register the governance module account.
	ak.SetModuleAccount(
		ctx,
		authtypes.NewEmptyModuleAccount(governancetypes.ModuleName, authtypes.Minter),
	)

	// Create the codec.
	encCfg := cosmostestutil.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
		bank.AppModuleBasic{},
	)

	// Create the base app msgRouter.
	msr := baseapp.NewMsgServiceRouter()

	// Create the governance keeper.
	gk := governancekeeper.NewKeeper(
		encCfg.Codec,
		testutil.EvmKey,
		ak,
		bk,
		sk,
		dk,
		msr,
		governancetypes.DefaultConfig(),
		authtypes.NewModuleAddress(governancetypes.ModuleName).String(),
	)

	// Register the msg Service Handlers.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v1.RegisterMsgServer(msr, governancekeeper.NewMsgServerImpl(gk))
	banktypes.RegisterMsgServer(msr, bankkeeper.NewMsgServerImpl(bk))

	// Set the Params and first proposal ID.
	params := v1.DefaultParams()
	gk.SetParams(ctx, params)
	gk.SetProposalID(ctx, 1)

	// Fund the caller with some coins.
	fundAccount(ctx, bk, caller, sdk.NewCoins(sdk.NewInt64Coin("usdc", 100000000)))

	return ctx, bk, gk
}

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "precompile/governance")
}

var _ = Describe("Governance Precompile", func() {
	var (
		ctx      sdk.Context
		bk       bankkeeper.Keeper
		gk       *governancekeeper.Keeper
		caller   sdk.AccAddress
		mockCtrl *gomock.Controller
		contract *Contract
	)

	BeforeEach(func() {
		t := GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		caller = evmutils.AddressToAccAddress(testutil.Alice)
		ctx, bk, gk = setup(mockCtrl, caller)
		contract = utils.MustGetAs[*Contract](NewContract(&gk))
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("Submitting a proposal", func() {
		It("should fail if the message is not of type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidAny))
			Expect(res).To(BeNil())
		})
		It("should fail if the initial deposit is wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				"invalid",
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidCoin))
			Expect(res).To(BeNil())
		})
		It("should fail if metadata is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if title is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if summary is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if expadited is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				"summary",
				123,
			)
			Expect(err).To(MatchError(polarisprecompile.ErrInvalidBool))
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("usdc", 100))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			fundAccount(ctx, bk, govAcct, initDeposit)
			message := &banktypes.MsgSend{
				FromAddress: govAcct.String(),
				ToAddress:   caller.String(),
				Amount:      initDeposit,
			}

			metadata := "metadata"
			title := "title"
			summary := "summary"

			msg, err := codectypes.NewAnyWithValue(message)
			Expect(err).ToNot(HaveOccurred())

			res, err := contract.SubmitProposal(
				ctx,
				evmutils.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{msg},
				[]generated.IGovernanceModuleCoin{
					{
						Amount: big.NewInt(100),
						Denom:  "usdc",
					},
				},
				metadata,
				title,
				summary,
				false,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})
})
