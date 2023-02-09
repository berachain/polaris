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
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/registry"
	libtypes "github.com/berachain/stargazer/lib/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `Factory` is a `PrecompileLogFactory` that builds Ethereum logs from Cosmos events. All Ethereum
// events must be registered with the factory before it can build logs during state transitions.
type Factory struct {
	// `events` is a registry of precompile logs, indexed by the Cosmos event type.
	events libtypes.Registry[string, *precompileLog]
	// `customValueDecoders` is a map of Cosmos attribute keys to attribute value decoder
	// functions for custom events.
	customValueDecoders ValueDecoders
}

// `NewFactory` returns a `Factory` with an empty events registry and custom value decoders map.
func NewFactory() *Factory {
	return &Factory{
		events:              registry.NewMap[string, *precompileLog](),
		customValueDecoders: make(ValueDecoders),
	}
}

// `RegisterEvent` registers an Ethereum event, and optionally its custom attribute value decoders,
// with the factory.
func (f *Factory) RegisterEvent(
	moduleEthAddress common.Address, abiEvent abi.Event, customValueDecoders ValueDecoders,
) {
	// register the ABI Event as a precompile log
	_ = f.events.Register(newPrecompileLog(moduleEthAddress, abiEvent))
	// register the event's custom value decoders, if any are provided
	for attr, decoder := range customValueDecoders {
		f.customValueDecoders[attr] = decoder
	}
}

// `Build` builds an Ethereum log from a Cosmos event.
//
// `Build` implements `events.PrecompileLogFactory`.
func (f *Factory) Build(event *sdk.Event) (*coretypes.Log, error) {
	// get the precompile log for the Cosmos event type
	pl := f.events.Get(event.Type)
	if pl == nil {
		return nil, ErrEthEventNotRegistered
	}

	var err error

	// validate the Cosmos event attributes
	if err = validateAttributes(pl, event); err != nil {
		return nil, err
	}

	// build the Ethereum log
	log := &coretypes.Log{
		Address: pl.precompileAddr,
	}
	if log.Topics, err = f.makeTopics(pl, event); err != nil {
		return nil, err
	}
	if log.Data, err = f.makeData(pl, event); err != nil {
		return nil, err
	}

	return log, nil
}
