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
	"math/big"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
)

//go:generate moq -out ./message.mock.go -pkg mock ../ Message

func NewEmptyMessage() *MessageMock {
	mockedMessage := &MessageMock{
		AccessListFunc: func() types.AccessList {
			return nil
		},
		DataFunc: func() []byte {
			return nil
		},
		FromFunc: func() common.Address {
			return common.Address{}
		},
		GasFunc: func() uint64 {
			return 0
		},
		GasFeeCapFunc: func() *big.Int {
			return big.NewInt(0)
		},
		GasPriceFunc: func() *big.Int {
			return big.NewInt(0)
		},
		GasTipCapFunc: func() *big.Int {
			return big.NewInt(0)
		},
		IsFakeFunc: func() bool {
			return false
		},
		NonceFunc: func() uint64 {
			return 0
		},
		ToFunc: func() *common.Address {
			return nil
		},
		ValueFunc: func() *big.Int {
			return big.NewInt(0)
		},
	}
	return mockedMessage
}
