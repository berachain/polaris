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
package journal

import (
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/store/snapkv/internal/cache"
)

// Compile-time check to ensure `entry` implements `libtypes.Cloneable[*Entry]`.
var _ libtypes.Cloneable[*Entry] = (*Entry)(nil)

// `entry` is a struct that contains information needed to set a value in a cache.
type Entry struct {
	Key  string       // key of the value to be set.
	Prev *cache.Value // Deep copy of object in cache map.
}

// `NewEntry` creates a new `entry` object for the given `store`, `key`, and `prev`
// cache value.
func NewEntry(key string, prev *cache.Value) *Entry {
	// create a deep copy of the prev field, if it is not nil.
	if prev != nil {
		prev = prev.Clone()
	}

	return &Entry{
		Key:  key,
		Prev: prev,
	}
}

// `Clone` creates a deep copy of an Entry.
//
// `Clone` implements `libtypes.Cloneable[*Entry]`.
func (ce Entry) Clone() *Entry {
	// Return a new entry object with the same Store and key fields as the original,
	// and the prev field set to the deep copy of the original prev field (or nil if the original
	// was nil).
	return NewEntry(ce.Key, ce.Prev)
}
