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
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/utils/slice"
)

// `Builder` is used to build a bloom filter for a block.
type Builder struct {
	receipts []*coretypes.Receipt
}

// `NewBuilder` returns a new bloom builder.
func NewBuilder() *Builder {
	return &Builder{
		receipts: slice.Make[*coretypes.Receipt](),
	}
}

// `AddReceiptToBlock` adds the tx receipt to the block bloom builder.
func (bb *Builder) AddReceiptToBlock(r *coretypes.Receipt) {
	bb.receipts = append(bb.receipts, r)
}

// `GetBloom` returns the currently built bloom filter.
func (bb *Builder) BuildBloom() coretypes.Bloom {
	return coretypes.CreateBloom(bb.receipts)
}
