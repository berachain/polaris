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
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
)

const (
	// `initCapacity` is the initial capacity of the `refund`'s snapshot stack.
	initCapacity      = 16
	refundRegistryKey = `refund`
)

// `refund` is a `Store` that tracks the refund counter.
type refund struct {
	ds.Stack[uint64] // snapshot stack
}

// `NewRefund` creates and returns a `refund`.
func NewRefund() state.RefundPlugin {
	return &refund{
		Stack: stack.New[uint64](initCapacity),
	}
}

func (r *refund) RegistryKey() string {
	return refundRegistryKey
}

// `Get` returns the current value of the refund counter.
func (r *refund) GetRefund() uint64 {
	return r.Peek()
}

// `Set` sets the refund counter to the given `gas`.
func (r *refund) AddRefund(gas uint64) {
	r.Push(r.Peek() + gas)
}

// `Sub` subtracts the given `gas` from the refund counter.
func (r *refund) SubRefund(gas uint64) {
	r.Push(r.Peek() - gas)
}

// `Snapshot` returns the current size of the refund counter, which is used to
// revert the refund counter to a previous value.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (r *refund) Snapshot() int {
	return r.Size()
}

// `RevertToSnapshot` reverts the refund counter to the value at the given `snap`.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (r *refund) RevertToSnapshot(snap int) {
	r.PopToSize(snap)
}

// `Write` implements `libtypes.Controllable`.
func (r *refund) Write() {}
