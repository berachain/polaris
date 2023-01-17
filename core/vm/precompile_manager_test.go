// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package vm_test

import (
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/testutil"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Precompile Manager", func() {
	var pr *vm.PrecompileRegistry
	var pm *vm.PrecompileManager
	var addr1 common.Address
	var addr2 common.Address

	BeforeEach(func() {
		addr1 = common.BytesToAddress([]byte{1})
		addr2 = common.BytesToAddress(authtypes.NewModuleAddress("test").Bytes())

		pr = vm.NewPrecompileRegistry(testutil.EvmKey)
		err := pr.RegisterStatelessContract(addr1, &mockStateless{})
		Expect(err).To(BeNil())
		err = pr.RegisterModule("test", &mockStatefulContract{})
		Expect(err).To(BeNil())
		pm = vm.NewPrecompileManager(pr)
	})

	Describe("Test Exists", func() {
		It("should correctly find and return stateless precompiles", func() {
			pc, found := pm.Exists(addr1)
			Expect(found).To(BeTrue())
			Expect(pc).ToNot(BeNil())

			pc, found = pm.Exists(common.BytesToAddress([]byte{2}))
			Expect(found).To(BeFalse())
			Expect(pc).To(BeNil())
		})

		It("should correctly find and return statelful precompiles", func() {
			spc, found := pm.Exists(addr2)
			Expect(found).To(BeTrue())
			Expect(spc).ToNot(BeNil())

			spc, found = pm.Exists(common.BytesToAddress([]byte{2}))
			Expect(found).To(BeFalse())
			Expect(spc).To(BeNil())
		})
	})
})
