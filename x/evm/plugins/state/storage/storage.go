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

package storage

import (
	"fmt"

	"github.com/berachain/stargazer/lib/errors"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// Compile-time type assertions.
var _ libtypes.Cloneable[Slots] = Slots{}
var _ fmt.Stringer = Slots{}

// `Storage` represents the account Storage map as a slice of single key value
// State pairs. This helps to ensure that the Storage map can be iterated over
// deterministically.
type Slots []Slot

// `ValidateBasic` performs basic validation of the Storage data structure.
// It checks for duplicate keys and calls `ValidateBasic` on each `State`.
func (s Slots) ValidateBasic() error {
	seenSlots := make(map[string]bool)
	for i, slot := range s {
		if seenSlots[slot.Key] {
			return errors.Wrapf(ErrInvalidState, "duplicate state key %d: %s", i, slot.Key)
		}

		if err := slot.ValidateBasic(); err != nil {
			return err
		}

		seenSlots[slot.Key] = true
	}
	return nil
}

// `String` implements `fmt.Stringer`.
func (s Slots) String() string {
	var str string
	for _, slot := range s {
		str += fmt.Sprintf("%s\n", slot.String())
	}

	return str
}

// `Clone` implements `types.Cloneable`.
func (s Slots) Clone() Slots {
	cpy := make(Slots, len(s))
	copy(cpy, s)

	return cpy
}
