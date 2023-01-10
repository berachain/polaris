// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

// . "github.com/onsi/ginkgo/v2"
// . "github.com/onsi/gomega"

// func TestPrecompileRegistry(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Precompile Registry Tests")
// }

// var _ = Describe("Precompile Registry Tests", func() {
// 	var ctx sdk.Context
// 	var pm *precompile.Registry
// 	var addr common.Address
// 	var mpc *vm.MockContract

// 	BeforeEach(func() {
// 		ctx, _, _, _ = testutil.SetupMinimalKeepers()
// 		var mockAbi abi.ABI
// 		err := mockAbi.UnmarshalJSON([]byte(mock.InterfaceMetaData.ABI))
// 		Expect(err).To(BeNil())
// 		mpc = vm.NewMockPrecompile(&mockAbi)

// 		pm = precompile.NewRegistry(testutil.EvmKey)
// 		addr = testutil.Alice
// 	})

// 	Describe("Test Inject And Get Precompile", func() {
// 		It("should correctly store and load factory precompile", func() {
// 			err := pm.Inject(ctx, addr, mpc)
// 			Expect(err).To(BeNil())

// 			getPrecompile := pm.GetPrecompileFn(ctx)
// 			_, found := getPrecompile(common.BigToAddress(big.NewInt(20)))
// 			Expect(found).To(BeFalse())

// 			pc, found := getPrecompile(addr)
// 			Expect(found).To(BeTrue())
// 			fpc, ok := pc.(precompile.FactoryContract)
// 			Expect(ok).To(BeTrue())
// 			Expect(fpc.Name()).To(Equal(vm.MockContractName))
// 		})
// 	})

// 	Describe("Test Add contract with events ", func() {
// 		It("should correctly register events for registered modules", func() {
// 			moduleAddr := common.BytesToAddress(authtypes.NewModuleAddress("test").Bytes())
// 			pm.RegisterModule("test", mpc)
// 			event := sdk.NewEvent("cosmos_event_type")
// 			log, err := pm.GetEventsRegistry().BuildEthLog(&event)
// 			Expect(err).To(BeNil())
// 			Expect(log.Address).To(Equal(moduleAddr))
// 		})
// 	})
// })
