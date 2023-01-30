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

package plugin

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
)

// `initCapacity` is the initial capacity of the `refund`'s snapshot stack.
const initCapacity = 16

// Compile-time assertion that `refund` implements `Base`.
var _ Base = (*refund)(nil)

// `refund` is a `Store` that tracks the refund counter.
type refund struct {
	ds.Stack[uint64] // snapshot stack
}

// `NewRefund` creates and returns a `refund`.
func NewRefund() *refund { //nolint: revive // its ok.
	return &refund{
		Stack: stack.New[uint64](initCapacity),
	}
}

// `Get` returns the current value of the refund counter.
func (rs *refund) GetRefund() uint64 {
	return rs.Peek()
}

// `Set` sets the refund counter to the given `gas`.
func (rs *refund) AddRefund(gas uint64) {
	rs.Push(rs.Peek() + gas)
}

// `Sub` subtracts the given `gas` from the refund counter.
func (rs *refund) SubRefund(gas uint64) {
	rs.Push(rs.Peek() - gas)
}

// `Snapshot` returns the current size of the refund counter, which is used to
// revert the refund counter to a previous value.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (rs *refund) Snapshot() int {
	return rs.Size()
}

// `RevertToSnapshot` reverts the refund counter to the value at the given `snap`.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (rs *refund) RevertToSnapshot(snap int) {
	rs.PopToSize(snap)
}
