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

package log

import (
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// Compile-time assertion.
var _ libtypes.Registrable[string] = (*precompileLog)(nil)

// `precompileLog` contains the required data for a precompile contract to produce an Ethereum
// compatible event log.
type precompileLog struct {
	// `eventType` is the corresponding Cosmos event type for this precompile log.
	eventType string
	// `address` is the Ethereum address used as the `Address` field for the Ethereum log.
	precompileAddr common.Address
	// `id` is the Ethereum event ID, to be used as an Ethereum event's first topic.
	id common.Hash
	// `indexedInputs` holds an Ethereum event's indexed arguments, emitted as event topics.
	indexedInputs abi.Arguments
	// `nonIndexedInputs` holds an Ethereum event's non-indexed arguments, emitted as event data.
	nonIndexedInputs abi.Arguments
}

// `newPrecompileLog` returns a new `precompileLog` with the given `precompileAddress` and
// `abiEvent`. It separates the indexed and non-indexed arguments of the event.
func newPrecompileLog(precompileAddr common.Address, abiEvent abi.Event) *precompileLog {
	return &precompileLog{
		eventType:        abi.ToUnderScore(abiEvent.Name),
		precompileAddr:   precompileAddr,
		id:               abiEvent.ID,
		indexedInputs:    abi.GetIndexed(abiEvent.Inputs),
		nonIndexedInputs: abiEvent.Inputs.NonIndexed(),
	}
}

// `RegistryKey` returns the Cosmos event type for the precompile log.
//
// `RegistryKey` implements `libtypes.Registrable`.
func (l *precompileLog) RegistryKey() string {
	return l.eventType
}
