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
//
//nolint:ireturn // ManagerI returns interfaces by design.
package journal

import (
	ds "github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/types"
)

// `ManagerI` is an interface that defines the methods that a journal manager must implement.
// Journal managers support holding cache entries and reverting to a certain index.
type ManagerI[T any] interface {
	// The journal manager is a stack.
	ds.StackI[CacheEntry]

	// `ManagerI` implements `Cloneable`.
	types.Cloneable[T]
}

// Compile-time check to ensure `Manager` implements `ManagerI`.
var _ ManagerI[*Manager] = (*Manager)(nil)

// `Manager` is a struct that holds a slice of CacheEntry instances.
type Manager struct {
	*ds.Stack[CacheEntry]
}

// `NewManager` creates and returns a new Manager instance with an empty journal.
func NewManager() *Manager {
	return &Manager{
		ds.NewStack[CacheEntry](),
	}
}

// `RevertToSize` does not modify the journal if `newSize` is invalid.
//
// `RevertToSize` implements `ManagerI`.
func (jm *Manager) PopToSize(newSize int) {
	// Revert and discard all journal entries after and including newSize.
	for i := jm.Size() - 1; i >= newSize; i-- {
		jm.Stack.PeekAt(i).Revert()
	}
	// Call parent
	jm.Stack.PopToSize(newSize)
}

// `Clone` returns a cloned journal by deep copying each CacheEntry.
//
// `Clone` implements `ManagerI[*Manager]`.
func (jm *Manager) Clone() *Manager {
	newManager := NewManager()
	for i := 0; i < jm.Size(); i++ {
		newManager.Push(jm.PeekAt(i).Clone())
	}
	return newManager
}
