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

package mock

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
)

const testBaseFee = 69

//go:generate moq -out ./host.mock.go -pkg mock ../ StargazerHostChain

func NewMockHost() *StargazerHostChainMock {
	// make and configure a mocked core.StargazerHostChain
	mockedStargazerHostChain := &StargazerHostChainMock{
		StargazerHeaderAtHeightFunc: func(contextMoqParam context.Context, v uint64) *types.StargazerHeader {
			return &types.StargazerHeader{
				Header: &types.Header{
					Number:  big.NewInt(int64(v)),
					BaseFee: big.NewInt(testBaseFee),
				},
				CachedHash: common.Hash{123},
			}
		},
		CumulativeGasUsedFunc: func(contextMoqParam context.Context, gasUsed uint64) uint64 {
			return 0
		},
	}
	return mockedStargazerHostChain
}
