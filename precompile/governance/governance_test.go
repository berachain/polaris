package governance

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	anothertestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/lib/utils"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	testutil "pkg.berachain.dev/polaris/testing/utils"
)

func setupTest(ctrl *gomock.Controller) (ctx sdk.Context, bk bankkeeper.Keeper, gk *governancekeeper.Keeper) {
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()

	// Create the distribution keeper.
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Create the codec.
	encCfg := anothertestutil.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
		bank.AppModuleBasic{},
	)

	// Register the governance module account.
	ak.SetModuleAccount(
		ctx,
		authtypes.NewEmptyModuleAccount(governancetypes.ModuleName, authtypes.Minter),
	)

	// Create the governance keeper.
	msr := baseapp.NewMsgServiceRouter()
	gk = governancekeeper.NewKeeper(
		encCfg.Codec,
		testutil.EvmKey, // test key.
		ak,
		bk,
		sk,
		dk,
		msr,
		governancetypes.DefaultConfig(),
		authtypes.NewModuleAddress(governancetypes.ModuleName).String(),
	)

	// Register all the handlers for the MsgServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v1.RegisterMsgServer(msr, governancekeeper.NewMsgServerImpl(gk))
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	banktypes.RegisterMsgServer(msr, bankkeeper.NewMsgServerImpl(bk))

	// Set the Params.
	params := v1.DefaultParams()
	gk.SetParams(ctx, params)

	// Set the first proposal ID.
	gk.SetProposalID(ctx, 1)

	// Set the bank balance of the governance account to hella "stake".
	FundGovernanceAccount(ctx, *gk, bk, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000000000000))))

	return ctx, bk, gk
}

type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "precompile/governance")
}

var _ = Describe("Governance precompile", func() {
	var (
		ctx      sdk.Context
		gk       *governancekeeper.Keeper
		bk       bankkeeper.Keeper
		mockCtrl *gomock.Controller
		caller   common.Address

		contract *Contract
	)

	BeforeEach(func() {
		t := GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		ctx, bk, gk = setupTest(mockCtrl)
		contract = utils.MustGetAs[*Contract](NewContract(&gk))
		caller = common.HexToAddress("0x1000000001231231")
		FundAccount(ctx, bk, sdk.AccAddress(caller.Bytes()), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000000000000))))
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("submitProposalHelper", func() {
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			bankMsg := &banktypes.MsgSend{
				FromAddress: govAcct.String(),
				ToAddress:   sdk.AccAddress(caller.Bytes()).String(),
				Amount:      initDeposit,
			}
			msg, err := codectypes.NewAnyWithValue(bankMsg)
			Expect(err).To(BeNil())
			res, err := contract.submitProposalHelper(
				ctx,
				[]*codectypes.Any{msg},
				[]generated.IGovernanceModuleCoin{
					{
						Amount: big.NewInt(100),
						Denom:  "stake",
					},
				},
				sdk.AccAddress(caller.Bytes()),
				"metadata",
				"title",
				"summary",
				false,
			)
			Expect(err).To(BeNil())
			fmt.Println(res, bk)
		})
	})
})

func FundAccount(ctx sdk.Context, bk bankkeeper.Keeper, account sdk.AccAddress, coins sdk.Coins) error {
	if err := bk.MintCoins(ctx, governancetypes.ModuleName, coins); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToAccount(ctx, governancetypes.ModuleName, account, coins)
}

func FundGovernanceAccount(ctx sdk.Context, gk governancekeeper.Keeper, bk bankkeeper.Keeper, coins sdk.Coins) error {
	return FundAccount(ctx, bk, gk.GetGovernanceAccount(ctx).GetAddress(), coins)
}
