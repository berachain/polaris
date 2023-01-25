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

package types

import (
	"fmt"
	"strings"

	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/gointerfaces"
)

// Compile-time interface assertions.
var _ gointerfaces.Cloneable[State] = &State{}
var _ fmt.Stringer = Storage{}

// `NewState` creates a new State instance.
func NewState(key, value common.Hash) State {
	return State{
		Key:   key.Hex(),
		Value: value.Hex(),
	}
}

// `ValidateBasic` checks to make sure the key is not empty.
func (s State) ValidateBasic() error {
	if strings.TrimSpace(s.Key) == "" {
		return errors.Wrapf(ErrInvalidState, "key cannot be empty %s", s.Key)
	}

	return nil
}

// `Clone` implements `types.Cloneable`.
func (s State) Clone() State {
	return State{
		Key:   s.Key,
		Value: s.Value,
	}
}
