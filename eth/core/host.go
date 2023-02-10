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

package core

import libtypes "github.com/berachain/stargazer/lib/types"

// =============================================================================
// Mandatory Plugins
// =============================================================================

// The following plugins MUST be implemented by the chain running Stargazer EVM and exposed via the
// `StargazerHostChain` interface. All plugins should be resettable with a given context.
type (
	GasPlugin interface {
		// `GasPlugin` implements `libtypes.Resettable`. Calling Reset() MUST reset the GasPlugin to a
		// default state.
		libtypes.Resettable

		// `ConsumeGas` MUST consume the supplied amount of gas. It MUST not panic due to a GasOverflow
		// and must return core.ErrOutOfGas if the amount of gas remaining is less than the amount
		// requested.
		ConsumeGas(uint64) error

		// `RefundGas` MUST refund the supplied amount of gas. It MUST not panic.
		RefundGas(uint64)

		// `GasRemaining` MUST return the amount of gas remaining. It MUST not panic.
		GasRemaining() uint64

		// `GasUsed` MUST return the amount of gas used during the current transaction. It MUST not panic.
		GasUsed() uint64

		// `CumulativeGasUsed` MUST return the amount of gas used during the current block. The value returned
		// MUST include any gas consumed during this transaction. It MUST not panic.
		CumulativeGasUsed() uint64

		// `MaxFeePerGas` MUST set the maximum amount of gas that can be consumed by the meter. It MUST not panic, but
		// instead, return an error, if the new gas limit is less than the currently consumed amount of gas.
		SetGasLimit(uint64) error
	}
)
