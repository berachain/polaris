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
	"github.com/berachain/stargazer/core/vm/precompile/event"
	"github.com/berachain/stargazer/types/abi"
)

// `HasEvents` is an interface that enforces the required function for a stateful precompile
// contract to implement if it wants to emit some (or all) of its Cosmos module's events as
// Ethereum event logs.
type HasEvents interface {
	// `ABIEvents` should return a map of Ethereum event names (should be CamelCase formatted) to
	// geth abi `Event` structs. NOTE: this can be directly loaded from the `Events` field of a
	// geth ABI struct, which can be built for a solidity library, interface, or contract.
	ABIEvents() map[string]abi.Event
}

// `HasCustomModuleEvents` is an interface that enforces the required functions for a stateful
// precompile contract to implement if it wants to emit some (or all) of its custom Cosmos module's
// events as Ethereum event logs.
type HasCustomModuleEvents interface {
	HasEvents

	// `CustomValueDecoders` should return a map of Cosmos event attribute keys to value decoder
	// functions for all supported events in the custom Cosmos module.
	CustomValueDecoders() map[EventType]event.ValueDecoders
}
