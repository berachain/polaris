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
package cachekv

import (
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `CacheEntry` is an interface for journal entries.
type CacheEntry interface {
	// `CacheEntry` implements `Cloneable`.
	libtypes.Cloneable[CacheEntry]

	// `Key` returns the key of the entry.
	Key() string

	// `Prev` returns the previous value of the entry.
	Prev() *cacheValue
}

// Compile-time check to ensure `cacheEntry` implements `journal.CacheEntry`.
var _ CacheEntry = (*cacheEntry)(nil)

// `cacheEntry` is a struct that contains information needed to set a value in a cache.
type cacheEntry struct {
	key  string      // key of the value to be set.
	prev *cacheValue // Deep copy of object in cache map.
}

// `newCacheEntry` creates a new `cacheEntry` object for the given `store`, `key`, and `prev`
// cache value.
func newCacheEntry(key string, prev *cacheValue) *cacheEntry {
	// create a deep copy of the prev field, if it is not nil.
	if prev != nil {
		prev = prev.Clone()
	}

	return &cacheEntry{
		key:  key,
		prev: prev,
	}
}

func (ce *cacheEntry) Key() string {
	return ce.key
}

func (ce *cacheEntry) Prev() *cacheValue {
	return ce.prev
}

// `Clone` creates a deep copy of the cacheEntry object.
// The deep copy contains the same Store and key fields as the original,
// and a deep copy of the prev field, if it is not nil.s
//
// `Clone` implements `journal.cacheEntry`.
//
//nolint:nolintlint,ireturn // by design.
func (ce *cacheEntry) Clone() CacheEntry {
	// Return a new cacheEntry object with the same Store and key fields as the original,
	// and the prev field set to the deep copy of the original prev field (or nil if the original
	// was nil).
	return newCacheEntry(ce.key, ce.prev)
}
