// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
package cache

import (
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `Entry` is an interface for journal entries.
type Entry interface {
	// `Entry` implements `Cloneable`.
	libtypes.Cloneable[Entry]

	// `Key` returns the key of the entry.
	Key() string

	// `Prev` returns the previous value of the entry.
	Prev() *Value
}

// Compile-time check to ensure `entry` implements `journal.Entry`.
var _ Entry = (*entry)(nil)

// `entry` is a struct that contains information needed to set a value in a cache.
type entry struct {
	key  string // key of the value to be set.
	prev *Value // Deep copy of object in cache map.
}

// `newEntry` creates a new `entry` object for the given `store`, `key`, and `prev`
// cache value.
func newEntry(key string, prev *Value) *entry {
	// create a deep copy of the prev field, if it is not nil.
	if prev != nil {
		prev = prev.Clone()
	}

	return &entry{
		key:  key,
		prev: prev,
	}
}

// `Key` returns the key of the entry.
func (ce *entry) Key() string {
	return ce.key
}

// `Prev` returns the previous value of the entry.
func (ce *entry) Prev() *Value {
	return ce.prev
}

// `Clone` creates a deep copy of the entry object.
// The deep copy contains the same Store and key fields as the original,
// and a deep copy of the prev field, if it is not nil.s
//
// `Clone` implements `journal.entry`.
//
//nolint:nolintlint,ireturn // by design.
func (ce *entry) Clone() Entry {
	// Return a new entry object with the same Store and key fields as the original,
	// and the prev field set to the deep copy of the original prev field (or nil if the original
	// was nil).
	return newEntry(ce.key, ce.prev)
}
