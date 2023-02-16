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

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin", func() {
	var ctx sdk.Context
	var p *plugin
	var blockGasMeter storetypes.GasMeter
	var txGasLimit = uint64(1000)

	BeforeEach(func() {
		// new block
		blockGasMeter = storetypes.NewGasMeter(uint64(5000))
		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		p = utils.MustGetAs[*plugin](NewPluginFrom(ctx))
	})

	It("correctly consume, refund, and report cumulative in the same block", func() {
		// tx 1
		err := p.SetTxGasLimit(txGasLimit)
		Expect(err).To(BeNil())
		err = p.TxConsumeGas(500)
		Expect(err).To(BeNil())
		Expect(p.TxGasUsed()).To(Equal(uint64(500)))
		Expect(p.TxGasRemaining()).To(Equal(uint64(500)))

		p.TxRefundGas(250)
		Expect(p.TxGasUsed()).To(Equal(uint64(250)))
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(250)))
		blockGasMeter.ConsumeGas(250, "") // finalize tx 1

		p.Prepare(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 2
		err = p.SetTxGasLimit(txGasLimit)
		Expect(err).To(BeNil())
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(250)))
		err = p.TxConsumeGas(1000)
		Expect(err).To(BeNil())
		Expect(p.TxGasUsed()).To(Equal(uint64(1000)))
		Expect(p.TxGasRemaining()).To(Equal(uint64(0)))
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(1250)))
		blockGasMeter.ConsumeGas(1000, "") // finalize tx 2

		p.Prepare(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 3
		err = p.SetTxGasLimit(txGasLimit)
		Expect(err).To(BeNil())
		Expect(p.CumulativeGasUsed()).To(Equal(uint64(1250)))
		err = p.TxConsumeGas(1000) // tx 3 should fail but no error here (250 over block limit)
		Expect(err).To(BeNil())
		Expect(p.TxGasUsed()).To(Equal(uint64(1000)))
		Expect(p.TxGasRemaining()).To(Equal(uint64(0)))
		// Expect(p.CumulativeGasUsed()).To(Equal(uint64(2000)))             // total is 2250, but capped at 2000
		// Expect(func() { blockGasMeter.ConsumeGas(1000, "") }).To(Panic()) // finalize tx 3
	})

	It("should error on overconsumption in tx", func() {
		err := p.SetTxGasLimit(txGasLimit)
		Expect(err).To(BeNil())
		err = p.TxConsumeGas(1000)
		Expect(err).To(BeNil())
		err = p.TxConsumeGas(1)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should error on uint64 overflow", func() {
		p.blockGasMeter = storetypes.NewInfiniteGasMeter()
		err := p.SetTxGasLimit(math.MaxUint64)
		Expect(err).To(BeNil())
		err = p.TxConsumeGas(math.MaxUint64)
		Expect(err).To(BeNil())
		err = p.TxConsumeGas(1)
		Expect(err.Error()).To(Equal("gas uint64 overflow"))
	})
})
