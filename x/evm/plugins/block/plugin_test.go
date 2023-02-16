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

package block

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/common"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block Plugin", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		p = utils.MustGetAs[*plugin](NewPlugin(&mockSHG{}))
		p.Prepare(ctx)
	})

	It("should give the constant base fee", func() {
		Expect(p.BaseFee()).To(Equal(bf))
	})

	It("should get the header at current height", func() {
		header := p.GetStargazerHeaderAtHeight(ctx.BlockHeight())
		Expect(header.Hash()).To(Equal(header.Header.Hash()))
		Expect(header.TxHash).To(Equal(common.BytesToHash(ctx.BlockHeader().DataHash)))
	})
})

// MOCKS BELOW.

type mockSHG struct {
	calls int
}

func (m *mockSHG) GetStargazerHeader(ctx sdk.Context, height int64) (*coretypes.StargazerHeader, bool) {
	m.calls++
	return nil, false
}
