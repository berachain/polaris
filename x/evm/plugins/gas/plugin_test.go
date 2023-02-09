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

package gas

import (
	"math"

	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = Describe("plugin", func() {
	var ctx sdk.Context
	var p *plugin
	var blockGasMeter sdk.GasMeter
	var txGasLimit = uint64(1000)

	BeforeEach(func() {
		// new block
		blockGasMeter = sdk.NewGasMeter(uint64(2000))
		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		p = utils.MustGetAs[*plugin](NewPluginFrom(ctx))
	})

	It("correctly consume, refund, and report cumulative in the same block", func() {
		// tx 1
		p.SetGasLimit(txGasLimit)
		err := p.ConsumeGas(500)
		Expect(err).To(BeNil())
		Expect(p.GasUsed()).To(Equal(uint64(500)))
		Expect(p.GasRemaining()).To(Equal(uint64(500)))

		p.RefundGas(250)
		Expect(p.GasUsed()).To(Equal(uint64(250)))
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(250)))
		blockGasMeter.ConsumeGas(250, "") // finalize tx 1

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 2
		p.SetGasLimit(txGasLimit)
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(250)))
		err = p.ConsumeGas(1000)
		Expect(err).To(BeNil())
		Expect(p.GasUsed()).To(Equal(uint64(1000)))
		Expect(p.GasRemaining()).To(Equal(uint64(0)))
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(1250)))
		blockGasMeter.ConsumeGas(1000, "") // finalize tx 2

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 3
		p.SetGasLimit(txGasLimit)
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(1250)))
		err = p.ConsumeGas(1000) // tx 3 should fail but no error here (250 over block limit)
		Expect(err).To(BeNil())
		Expect(p.GasUsed()).To(Equal(uint64(1000)))
		Expect(p.GasRemaining()).To(Equal(uint64(0)))
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(2000)))             // total is 2250, but capped at 2000
		Expect(func() { blockGasMeter.ConsumeGas(1000, "") }).To(Panic()) // finalize tx 3
	})

	It("should error on overconsumption in tx", func() {
		p.SetGasLimit(txGasLimit)
		err := p.ConsumeGas(1000)
		Expect(err).To(BeNil())
		err = p.ConsumeGas(1)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should error on uint64 overflow", func() {
		p.SetGasLimit(math.MaxUint64)
		err := p.ConsumeGas(math.MaxUint64)
		Expect(err).To(BeNil())
		err = p.ConsumeGas(1)
		Expect(err.Error()).To(Equal("gas uint64 overflow"))
	})
})
