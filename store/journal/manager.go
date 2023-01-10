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

package journal

// `ManagerI` is an interface that defines the methods that a journal manager
// must implement.
type ManagerI interface {
	// `Append` adds a new CacheEntry instance to the journal. The Size
	// method returns the current number of entries in the journal.
	Append(ce CacheEntry)
	Size() int
	Get(i int) CacheEntry
	RevertToSize(newSize int)
	Clone() ManagerI
}

// `Manager` is a struct that holds an array of CacheEntry instances.
type Manager struct {
	journal []CacheEntry
}

// `NewManager` creates and returns a new Manager instance with an empty
// journal.
func NewManager() *Manager {
	return &Manager{}
}

// `Append` implements `ManagerI`.
func (jm *Manager) Append(ce CacheEntry) {
	jm.journal = append(jm.journal, ce)
}

// `Size` returns the current number of entries in the journal.
func (jm *Manager) Size() int {
	return len(jm.journal)
}

// `Get` returns the CacheEntry instance at the given index.
func (jm *Manager) Get(i int) CacheEntry {
	return jm.journal[i]
}

// `RevertToSize` reverts and discards all journal entries after and including
// the given size.
func (jm *Manager) RevertToSize(newSize int) {
	// Revert and discard all journal entries after and including newSize.
	for j := len(jm.journal) - 1; j >= newSize; j-- {
		jm.journal[j].Revert()
	}

	// Discard all journal entries after and including newSize, such that
	// now len(jm.journal) == newSize.
	jm.journal = jm.journal[:newSize]
}

// `Clone` creates and returns a new Manager instance with a cloned journal.
func (jm *Manager) Clone() *Manager {
	newJournal := make([]CacheEntry, len(jm.journal))
	for i := 0; i < len(jm.journal); i++ {
		newJournal[i] = jm.journal[i].Clone()
	}
	return &Manager{
		journal: newJournal,
	}
}
