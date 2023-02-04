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

package events

import (
	libtypes "github.com/berachain/stargazer/lib/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TODO: determine appropriate events journal capacity.
	initJournalCapacity = 32
	managerRegistryKey  = `events`
)

type manager struct {
	*sdk.EventManager // pointer to the event manager floating aroundmon the context.
}

// `NewManager` creates and returns a controllable event manager from the given Cosmos SDK context.
func NewManager(em *sdk.EventManager) libtypes.Controllable[string] {
	return &manager{
		EventManager: em,
	}
}

// `Registry` implements `libtypes.Registrable`.
func (m *manager) RegistryKey() string {
	return managerRegistryKey
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (m *manager) Snapshot() int {
	return len(m.Events())
}

// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (m *manager) RevertToSnapshot(id int) {
	temp := m.Events()
	*m.EventManager = *sdk.NewEventManager()
	m.EmitEvents(temp[:id])
}

// `Finalize` implements `libtypes.Finalizable`.
func (m *manager) Finalize() {}
