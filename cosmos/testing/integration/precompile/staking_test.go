package precompile

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCosmosPrecompiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/precompile:integration")
}

var _ = Describe("Staking", func() {
	var net *network.Network
	var client *ethclient.Client
	var validator common.Address
	var stakingPrecompile *bindings.StakingModule
	BeforeEach(func() {
		net, client = StartPolarisNetwork(GinkgoT())
		validator = common.BytesToAddress(net.Validators[0].Address.Bytes())
		stakingPrecompile, _ = bindings.NewStakingModule(
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"), client)
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	It("getActiveValidators()", func() {
		validators, err := stakingPrecompile.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(validators).To(ContainElement(validator))
	})

	It("should be able to delegate tokens", func() {
		delegated, err := stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

		tx, err := stakingPrecompile.Delegate(BuildTransactor(client),
			validator, big.NewInt(100000000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		delegated, err = stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		Expect(err).ToNot(HaveOccurred())
		Expect(delegated).To(Equal(big.NewInt(100000000000)))
	})
})
