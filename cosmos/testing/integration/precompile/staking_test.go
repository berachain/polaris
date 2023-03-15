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
	tbindings "pkg.berachain.dev/polaris/contracts/bindings/testing"

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
		_ = stakingPrecompile
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	// It("getActiveValidators()", func() {
	// 	validators, err := stakingPrecompile.GetActiveValidators(nil)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(validators).To(ContainElement(validator))
	// })

	// It("should be able to delegate tokens", func() {
	// 	delegated, err := stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(delegated.Cmp(big.NewInt(0))).To(Equal(0))

	// 	tx, err := stakingPrecompile.Delegate(BuildTransactor(client),
	// 		validator, big.NewInt(100000000000))
	// 	Expect(err).ToNot(HaveOccurred())
	// 	ExpectMined(client, tx)
	// 	ExpectSuccessReceipt(client, tx)

	// 	delegated, err = stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
	// 	Expect(err).ToNot(HaveOccurred())
	// 	Expect(delegated).To(Equal(big.NewInt(100000000000)))
	// })

	It("should be able to call a precompile from a smart contract", func() {
		// Deploy a contract
		_, tx, contract, err := tbindings.DeployLiquidStaking(
			BuildTransactor(client),
			client,
			"myToken",
			"MTK",
			common.HexToAddress("0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF"),
			common.BytesToAddress(net.Validators[0].ValAddress.Bytes()),
		)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		value, err := contract.TotalDelegated(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Cmp(big.NewInt(0))).To(Equal(0))

		addresses, err := contract.GetActiveValidators(nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(addresses).To(HaveLen(1))
		Expect(addresses[0]).To(Equal(validator))

		// Send tokens to the contract
		txr := BuildTransactor(client)
		txr.Value = big.NewInt(100000000000)
		tx, err = contract.Receive(txr)
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		tx, err = contract.Delegate(BuildTransactor(client), big.NewInt(100000000000))
		Expect(err).ToNot(HaveOccurred())
		ExpectMined(client, tx)
		ExpectSuccessReceipt(client, tx)

		// delegated, err = stakingPrecompile.GetDelegation(nil, network.TestAddress, validator)
		// Expect(err).ToNot(HaveOccurred())
		// Expect(delegated).To(Equal(big.NewInt(100000000000)))
	})
})
