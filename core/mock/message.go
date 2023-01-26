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

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func NewEmptyMessage() *MessageMock {
	m := new(MessageMock)
	m.FromFunc = func() common.Address {
		return common.Address{}
	}
	m.GasPriceFunc = func() *big.Int {
		return big.NewInt(0)
	}
	m.GasFunc = func() uint64 {
		return 0
	}
	m.GasFeeCapFunc = func() *big.Int {
		return big.NewInt(0)
	}
	m.ValueFunc = func() *big.Int {
		return big.NewInt(0)
	}
	m.DataFunc = func() []byte {
		return []byte{}
	}
	m.ToFunc = func() *common.Address {
		return nil
	}
	m.AccessListFunc = func() coretypes.AccessList {
		return nil
	}
	return m
}
