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
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
)

// `initCapacity` is the initial capacity of the `RefundPlugin`'s snapshot stack.
const initCapacity = 32

// Compile-time assertion to ensure `RefundPlugin` adheres to `Store`.
var _ Store = (*RefundPlugin)(nil)

// `RefundPlugin` is a `Store` that tracks the refund counter.
type RefundPlugin struct {
	ds.Stack[uint64] // snapshot stack
}

// `NewRefund` creates and returns a `RefundPlugin`.
func NewRefund() *RefundPlugin {
	return &RefundPlugin{
		Stack: stack.New[uint64](initCapacity),
	}
}

// `Name` returns the name of the `RefundPlugin`.
func (rs *RefundPlugin) Name() string {
	return "refund"
}

// `Get` returns the current value of the refund counter.
func (rs *RefundPlugin) Get() uint64 {
	return rs.Peek()
}

// `Set` sets the refund counter to the given `amount`.
func (rs *RefundPlugin) Add(amount uint64) {
	rs.Push(rs.Peek() + amount)
}

// `Sub` subtracts the given `amount` from the refund counter.
func (rs *RefundPlugin) Sub(amount uint64) {
	rs.Push(rs.Peek() - amount)
}

// `Snapshot` returns the current size of the refund counter, which is used to
// revert the refund counter to a previous value.
func (rs *RefundPlugin) Snapshot() int {
	return rs.Size()
}

// `RevertToSnapshot` reverts the refund counter to the value at the given `snap`.
func (rs *RefundPlugin) RevertToSnapshot(snap int) {
	rs.PopToSize(snap)
}
