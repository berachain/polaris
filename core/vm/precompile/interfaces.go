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
	"github.com/berachain/stargazer/types/abi"
)

// `HasEvents` defines the minimum set of required functions for a stateful precompile contract
// to implement if it wants to emit each of its Cosmos emitted events as an Eth event log.
type HasEvents interface {
	// `CosmosEventTypes` should return a slice of the Cosmos event type strings (under_score)
	// formatting for all Cosmos events the stateful precompile wishes to emit as Eth events.
	CosmosEventTypes() []string

	// `ABIEvents` should return a map of Eth event names (should be CamelCase formatted) to
	// geth abi `Event` structs. Note: this can be directly loaded from the `Events` field of a
	// geth ABI struct, which is built from a solidity library, interface or contract.
	ABIEvents() map[string]abi.Event
}
