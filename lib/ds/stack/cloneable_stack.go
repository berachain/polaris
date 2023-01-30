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
package stack

import (
	"github.com/berachain/stargazer/lib/ds"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `cloneableStack` is a struct that holds a slice of CacheEntry instances.
type cloneableStack[T libtypes.Cloneable[T]] struct {
	// The `cloneableStack` is a `ds.Stack`.
	ds.Stack[T]
}

// `NewCloneable` creates and returns a new cloneableStack instance.
func NewCloneable[T libtypes.Cloneable[T]](capacity int) cloneableStack[T] { //nolint:revive // it's ok.
	return cloneableStack[T]{
		New[T](capacity),
	}
}

// `Clone` returns a cloned journal by deep copyign each CacheEntry.
//
// `Clone` implements `CloneableStack[T]`.
func (cs cloneableStack[T]) Clone() ds.CloneableStack[T] {
	newcloneableStack := NewCloneable[T](cs.Capacity())
	for i := 0; i < cs.Size(); i++ {
		newcloneableStack.Push(cs.PeekAt(i).Clone())
	}
	return newcloneableStack
}
