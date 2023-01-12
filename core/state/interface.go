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

package state

import (
	"math/big"

	"github.com/berachain/stargazer/common"
	coretypes "github.com/berachain/stargazer/core/types"
)

// ExtStateDB defines an extension to the interface provided by the go-ethereum codebase to
// support additional state transition functionalities. In particular it supports getting the
// cosmos sdk context for natively running stateful precompiled contracts.
type StargazerStateDB interface {
	GethStateDB

	// TransferBalance transfers the balance from one account to another
	TransferBalance(common.Address, common.Address, *big.Int)

	// GetSavedErr returns the error saved in the statedb
	GetSavedErr() error

	// GetLogs returns the logs generated during the transaction
	Logs() []*coretypes.Log

	// Commit writes the state to the underlying multi-store
	Commit() error

	// PrepareForTransition prepares the statedb for a new transition
	// by setting the block hash, tx hash, tx index and tx log index.
	PrepareForTransition(common.Hash, common.Hash, uint, uint)
}
