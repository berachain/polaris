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

package precompile

import (
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/types/abi"
)

// `EventType` is the name of an Ethereum event, which is equivalent to the CamelCase version of
// its corresponding event's `Type`.
type EventType string

// `LogRegistry` builds Ethereum logs from Cosmos events.
type LogRegistry struct {
	// `eventTypesToLogs` is a map of `EventType`s to `*log.PrecompileLog` for all supported
	// events.
	eventTypesToLogs map[EventType]*log.PrecompileLog
}

// `NewLogRegistry` creates and returns a new, empty `LogRegistry`.
func NewLogRegistry() *LogRegistry {
	return &LogRegistry{
		eventTypesToLogs: make(map[EventType]*log.PrecompileLog),
	}
}

// `RegisterEvent` registers an Ethereum event as a precompile log.
func (lr *LogRegistry) RegisterEvent(
	moduleEthAddress common.Address,
	abiEvent abi.Event,
	customModuleAttributes log.ValueDecoders,
) error {
	if _, found := lr.eventTypesToLogs[EventType(abiEvent.Name)]; found {
		return errors.Wrap(ErrEthEventAlreadyRegistered, abiEvent.Name)
	}

	var err error
	lr.eventTypesToLogs[EventType(abiEvent.Name)], err = log.NewPrecompileLog(
		moduleEthAddress,
		abiEvent,
		customModuleAttributes,
	)
	return err
}

// `GetPrecompileLog` returns the precompile log corresponding to the given event.
func (lr *LogRegistry) GetPrecompileLog(eventType string) *log.PrecompileLog {
	return lr.eventTypesToLogs[EventType(abi.ToCamelCase(eventType))]
}
