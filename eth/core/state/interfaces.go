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

package state

import (
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	libtypes "github.com/berachain/stargazer/lib/types"
)

type LogsPlugin interface {
	// `LogsPlugin` implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// `AddLog` adds a log to the state
	AddLog(*coretypes.Log)
	// `BuildLogsAndClear` returns the logs of the tx with the given metadata
	BuildLogsAndClear(common.Hash, common.Hash, uint, uint) []*coretypes.Log
}

// `RefundPlugin` is a `Store` that tracks the refund counter.
type RefundPlugin interface {
	// `RefundPlugin` implements `libtypes.Controllable`.
	libtypes.Controllable[string]
	// `GetRefund` returns the current value of the refund counter.
	GetRefund() uint64
	// `AddRefund` sets the refund counter to the given `gas`.
	AddRefund(gas uint64)
	// `SubRefund` subtracts the given `gas` from the refund counter.
	SubRefund(gas uint64)
}
