package precompile

import (
	"os"
	"testing"

	"pkg.berachain.dev/polaris/eth/common"

	. "github.com/onsi/ginkgo/v2"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"

	"pkg.berachain.dev/polaris/cosmos/testing/integration"

	. "github.com/onsi/gomega"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/precompile:integration")
}

var (
	tf                *integration.TestFixture
	stakingPrecompile *bindings.StakingModule
	validator         common.Address

	// Distr Precompile.
	distributionPrecompile *bindings.DistributionModule
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	validator = common.Address(tf.Network.Validators[0].Address.Bytes())
	stakingPrecompile, _ = bindings.NewStakingModule(
		common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), tf.EthClient)

	distributionPrecompile, _ = bindings.NewDistributionModule(
		common.HexToAddress("0x93354845030274cD4bf1686Abd60AB28EC52e1a7"),
		tf.EthClient,
	)
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})
