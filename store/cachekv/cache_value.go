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

// `cValue` represents a cached value.
// If dirty is true, it indicates the cached value is different from the underlying value.
type cValue struct {
	value []byte
	dirty bool
}

// `NewCValue` creates a new cValue object with the given value and dirty flag.
func NewCValue(v []byte, d bool) *cValue { //nolint: revive // todo:revist.
	return &cValue{
		value: v,
		dirty: d,
	}
}

// `Dirty` returns the dirty flag of the cValue object.
func (cv *cValue) Dirty() bool {
	return cv.dirty
}

// `Value` returns the value of the cValue object.
func (cv *cValue) Value() []byte {
	return cv.value
}

// `deepCopy` creates a new cValue object with the same value and dirty flag as the original
// cValue object. This function is used to create a deep copy of the prev field in
// DeleteCacheValue and SetCacheValue objects, so that modifications to the original prev value do
// not affect the cloned DeleteCacheValue or SetCacheValue object.
func (cv *cValue) deepCopy() *cValue {
	// Return a new cValue with the same value and dirty flag
	return &cValue{
		value: append([]byte(nil), cv.value...),
		dirty: cv.dirty,
	}
}
