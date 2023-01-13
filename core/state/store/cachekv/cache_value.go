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

package cachekv

import "github.com/berachain/stargazer/types"

// Compile-time assertion that `cacheValue` implements `types.Cloneable`.
var _ types.Cloneable[*cacheValue] = (*cacheValue)(nil)

// `cacheValue` represents a cached value in the cachekv store.
// If dirty is true, it indicates the cached value is different from the underlying value.
type cacheValue struct {
	value []byte
	dirty bool
}

// `newCacheValue` creates a new `cacheValue` object with the given `value` and `dirty` flag.
func newCacheValue(v []byte, d bool) *cacheValue { 
	return &cacheValue{
		value: v,
		dirty: d,
	}
}

// `Clone` implements `types.Cloneable`.
func (cv *cacheValue) Clone() *cacheValue {
	// Return a new cacheValue with the same value and dirty flag
	return newCacheValue(append([]byte(nil), cv.value...), cv.dirty)
}
