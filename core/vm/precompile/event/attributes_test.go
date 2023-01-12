package event_test

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/vm/precompile/event"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestAttributes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Attributes Test Suite")
}

var _ = Describe("Attributes Test Suite", func() {
	var gethValue any
	var err error

	Describe("Test Default Attribute Value Decoder Functions", func() {
		It("should correctly convert sdk coin strings to big.Int", func() {
			denom10 := sdk.NewCoin("denom", sdk.NewInt(10))
			gethValue, err = event.ConvertSdkCoin(denom10.String())
			Expect(err).To(BeNil())
			bigVal, ok := gethValue.(*big.Int)
			Expect(ok).To(BeTrue())
			Expect(bigVal).To(Equal(big.NewInt(10)))
		})

		It("should correctly convert creation height to int64", func() {
			creationHeightStr := strconv.FormatInt(55, 10)
			gethValue, err = event.ConvertCreationHeight(creationHeightStr)
			Expect(err).To(BeNil())
			int64Val, ok := gethValue.(int64)
			Expect(ok).To(BeTrue())
			Expect(int64Val).To(Equal(int64(55)))
		})

		It("should correctly convert ValAddress to common.Address", func() {
			valAddr := sdk.ValAddress([]byte("alice"))
			gethValue, err = event.ConvertValAddressFromBech32(valAddr.String())
			Expect(err).To(BeNil())
			valAddrVal, ok := gethValue.(common.Address)
			Expect(ok).To(BeTrue())
			Expect(valAddrVal).To(Equal(common.ValAddressToEthAddress(valAddr)))
		})

		It("should correctly convert AccAddress to common.Address", func() {
			accAddr := sdk.AccAddress([]byte("alice"))
			gethValue, err = event.ConvertAccAddressFromBech32(accAddr.String())
			Expect(err).To(BeNil())
			accAddrVal, ok := gethValue.(common.Address)
			Expect(ok).To(BeTrue())
			Expect(accAddrVal).To(Equal(common.AccAddressToEthAddress(accAddr)))
		})
	})
})
