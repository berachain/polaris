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
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = Describe("plugin", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext()
		p = utils.MustGetAs[*plugin](NewPluginFrom(ctx))
	})

	It("correctly consume, refund gas and reset", func() {
		p.SetGasLimit(uint64(1000))
		p.ConsumeGas(500)
		Expect(p.GasUsed()).To(Equal(uint64(500)))
		Expect(p.GasRemaining()).To(Equal(uint64(500)))
	})
})
