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
	"math/big"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/vm"
)

// Compile-time interface check.
var _ vm.CanTransferFunc = CanTransfer
var _ vm.TransferFunc = Transfer

// `CanTransfer` checks whether there are enough funds in the address' account to make a transfer.
// NOTE: This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(sdb vm.StateDB, addr common.Address, amount *big.Int) bool {
	return sdb.GetBalance(addr).Cmp(amount) >= 0
}

// `Transfer` subtracts amount from sender and adds amount to recipient using the `vm.StateDB`.
func Transfer(sdb vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	// We use `TransferBalance` to use the same logic as the native transfer in x/bank.
	sdb.(state.ExtStateDB).TransferBalance(sender, recipient, amount)
}
