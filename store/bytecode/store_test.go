package bytecode

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/sims"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/crypto"
)

func TestByteCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "store/bytecode")
}

var _ = Describe("bytecodeStore", func() {
	var (
		addr  = common.BytesToAddress([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
		code1 = []byte{1}
		dbDir = sims.NewAppOptionsWithFlagHome("/tmp/berachain")
		store = NewByteCodeStore(dbDir)
	)

	It("should set and ge byte code", func() {
		store.StoreByteCode(addr, code1)
		codeHash := crypto.Keccak256Hash(code1[:])
		code, err := store.GetByteCode(addr, codeHash)
		Expect(err).To(BeNil())
		Expect(code).To(Equal(code1))
	})

	It("should fail to get byte code if the code hash does not match", func() {
		store.StoreByteCode(addr, code1)
		codeHash := crypto.Keccak256Hash([]byte{2})
		code, err := store.GetByteCode(addr, codeHash)
		Expect(err).To(Equal(ErrByteCodeDoesNotMatch))
		Expect(code).To(BeNil())
	})

	It("should iterate over byte code", func() {
		log := make([]byte, 0)

		store.StoreByteCode(addr, code1)
		store.IterateByteCode(nil, nil, func(addr common.Address, code []byte) bool {
			log = append(log, code...)
			return true // break the iteration
		})

		Expect(log).To(Equal(code1))
	})

	It("should set and get version", func() {
		store.SetVersion(1)
		Expect(store.GetVersion()).To(Equal(int64(1)))
	})
})
