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

import (
	"context"

	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

// `Host` provides the EVM with the underlying application and consensus engine to run.
type Host interface {
	Application
	Consensus
}

// `Application` defines the required function for a specific application layer to implement.
type Application interface {
	// `GasMeter` should return the application's native gas meter.
	GasMeter(context.Context) StargazerGasMeter
}

// `Consensus` defines the required functions for a consensus engine to implement.
type Consensus interface {
	// `GetBlockHashFunc` should return a block hash function getter for a particular block. It is
	// used by the BLOCKHASH EVM op code.
	GetBlockHashFunc(context.Context) vm.GetHashFunc

	// `GetCoinbase` gets the coinbase address from the consensus engine.
	GetCoinbase(context.Context) (common.Address, error)
}

// `StargazerGasMeter` defines the required function for an application's native gas meter.
type StargazerGasMeter interface {
	ConsumeGas(gas uint64, descriptor string)
}
