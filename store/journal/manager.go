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

import "github.com/berachain/stargazer/types"

// `ManagerI` is an interface that defines the methods that a journal manager must implement.
// Journal managers support holding `CacheEntry`s and reverting to a certain index.
type ManagerI[T any] interface {
	// `Append` adds a new `CacheEntry` instance to the journal. The Size method returns the current
	// number of entries in the journal.
	Append(ce CacheEntry)

	// `Size` returns the current number of entries in the journal.
	Size() int

	// `Get` returns the `CacheEntry` instance at the given index.
	Get(i int) CacheEntry

	// `RevertToSize` reverts and discards all journal entries after and including the given size.
	RevertToSize(newSize int)

	// `ManagerI` implements `Cloneable`.
	types.Cloneable[T]
}

// Compile-time check to ensure `Manager` implements `ManagerI`.
var _ ManagerI[*Manager] = (*Manager)(nil)

// `Manager` is a struct that holds a slice of `CacheEntry` instances.
type Manager struct {
	journal []CacheEntry
}

// `NewManager` creates and returns a new Manager instance with an empty journal.
func NewManager() *Manager {
	return &Manager{
		journal: make([]CacheEntry, 0),
	}
}

// `Append` implements `ManagerI`.
func (jm *Manager) Append(ce CacheEntry) {
	jm.journal = append(jm.journal, ce)
}

// `Size` implements `ManagerI`.
func (jm *Manager) Size() int {
	return len(jm.journal)
}

// `Get` returns nil if index `i` is invalid.
//
// `Get` implements `ManagerI`.
func (jm *Manager) Get(i int) CacheEntry {
	if i < 0 || i >= len(jm.journal) {
		return nil
	}
	return jm.journal[i]
}

// `RevertToSize` does not modify the journal if `newSize` is invalid.
//
// `RevertToSize` implements `ManagerI`.
func (jm *Manager) RevertToSize(newSize int) {
	if newSize > len(jm.journal) {
		return
	}

	// Revert and discard all journal entries after and including newSize.
	for j := len(jm.journal) - 1; j >= newSize; j-- {
		jm.journal[j].Revert()
	}

	// Discard all journal entries after and including newSize, such that now
	// len(jm.journal) == newSize.
	jm.journal = jm.journal[:newSize]
}

// `Clone` returns a cloned journal by deep copying each `CacheEntry`.
//
// `Clone` implements `ManagerI[*Manager]`.
func (jm *Manager) Clone() *Manager {
	newJournal := make([]CacheEntry, len(jm.journal))
	for i := 0; i < len(jm.journal); i++ {
		newJournal[i] = jm.journal[i].Clone()
	}
	return &Manager{
		journal: newJournal,
	}
}
