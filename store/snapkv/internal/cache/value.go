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

package cache

import libtypes "github.com/berachain/stargazer/lib/types"

// Compile-time assertion that `Value` implements `types.Cloneable`.
var _ libtypes.Cloneable[*Value] = (*Value)(nil)

// `Value` represents a cached value in the cachekv store.
// If dirty is true, it indicates the cached value is different from the underlying value.
type Value struct {
	Value []byte
	Dirty bool
}

// `NewValue` creates a new `cacheValue` object with the given `value` and `dirty` flag.
func NewValue(bz []byte, d bool) *Value {
	return &Value{
		Value: bz,
		Dirty: d,
	}
}

// `Clone` implements `types.Cloneable`.
func (cv *Value) Clone() *Value {
	// Return a new cacheValue with the same value and dirty flag
	if cv.Value == nil {
		return NewValue(nil, cv.Dirty)
	}
	bz := make([]byte, len(cv.Value))
	copy(bz, cv.Value)
	return NewValue(bz, cv.Dirty)
}
