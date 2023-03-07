package governance

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	anothertestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	testutil "pkg.berachain.dev/polaris/testing/utils"
)

func setupTest(ctrl *gomock.Controller) (ctx sdk.Context, bk bankkeeper.Keeper, gk *governancekeeper.Keeper) {
	ctx, ak, bk, sk := testutil.SetupMinimalKeepers()

	// Create the distribution keeper.
	dk := govtestutil.NewMockDistributionKeeper(ctrl)

	// Create the codec.
	encCfg := anothertestutil.MakeTestEncodingConfig(
		gov.AppModuleBasic{},
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
	)

	BeforeEach(func() {
		t := GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		ctx, bk, gk = setupTest(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should compile", func() {
		fmt.Println(ctx)
		fmt.Println(bk)
		fmt.Println(gk)
	})
})
