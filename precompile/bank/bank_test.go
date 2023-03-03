package bank

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/x/evm/plugins/precompile/log"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBankPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "precompile/bank")
}

var _ = Describe("Bank Precompile Test", func() {
	var (
		contract *Contract
		addr     sdk.AccAddress
		factory  *log.Factory
	)

	BeforeEach(func() {
		contract = utils.MustGetAs[*Contract](NewPrecompileContract())
		addr = sdk.AccAddress([]byte("bank"))

		// Register the events.
		factory = log.NewFactory()
		factory.RegisterAllEvents([]vm.RegistrablePrecompile{contract})
	})

	It("should register the send event", func() {
		event := sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).To(BeNil())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the transfer event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeTransfer,
			sdk.NewAttribute(banktypes.AttributeKeyRecipient, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).To(BeNil())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin spent event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinSpent,
			sdk.NewAttribute(banktypes.AttributeKeySpender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).To(BeNil())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin received event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinReceived,
			sdk.NewAttribute(banktypes.AttributeKeyReceiver, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).To(BeNil())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the burn event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinBurn,
			sdk.NewAttribute(banktypes.AttributeKeyBurner, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdk.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).To(BeNil())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})
})
