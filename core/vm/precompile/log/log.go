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
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

// `PrecompileLog` contains the required data for a Cosmos precompile contract to produce an
// Ethereum formatted event log.
type PrecompileLog struct {
	// `address` is the Ethereum address which represents a Cosmos module's account address.
	precompileAddr common.Address

	// `id` is the Ethereum event ID, to be used as an Ethereum event's first topic.
	id common.Hash

	// `indexedInputs` holds an Ethereum event's indexed arguments, emitted as event topics.
	indexedInputs abi.Arguments

	// `nonIndexedInputs` holds an Ethereum event's non-indexed arguments, emitted as event data.
	nonIndexedInputs abi.Arguments
}

// `NewPrecompileLog` returns a new `PrecompileLog` with the given `precompileAddress`, `abiEvent`,
// and optional `customValueDecoders`.
func NewPrecompileLog(
	precompileAddr common.Address,
	abiEvent abi.Event,
) (*PrecompileLog, error) {
	indexedInputs, err := abi.GetIndexed(abiEvent.Inputs)
	if err != nil {
		return nil, err
	}
	pe := &PrecompileLog{
		precompileAddr:   precompileAddr,
		id:               abiEvent.ID,
		indexedInputs:    indexedInputs,
		nonIndexedInputs: abiEvent.Inputs.NonIndexed(),
	}
	return pe, nil
}

// `GetPrecompileAddress` returns the address of the precompile contract in which
// this event is to be emitted from.
func (pe *PrecompileLog) GetPrecompileAddress() common.Address {
	return pe.precompileAddr
}

// `ID` returns the event ID.
func (pe *PrecompileLog) ID() common.Hash {
	return pe.id
}

// `IndexedInputs` returns the indexed arguments of the event.
func (pe *PrecompileLog) IndexedInputs() abi.Arguments {
	return pe.indexedInputs
}

// `NonIndexedInputs` returns the non-indexed arguments of the event.
func (pe *PrecompileLog) NonIndexedInputs() abi.Arguments {
	return pe.nonIndexedInputs
}
