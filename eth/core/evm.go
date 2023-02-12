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

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
)

// Compile-time type assertion.
var _ vm.CanTransferFunc = canTransfer
var _ vm.TransferFunc = transfer

// `canTransfer` checks whether there are enough funds in the address' account to make a transfer.
// NOTE: This does not take the necessary gas in to account to make the transfer valid.
func canTransfer(sdb vm.GethStateDB, addr common.Address, amount *big.Int) bool {
	return sdb.GetBalance(addr).Cmp(amount) >= 0
}

// `transfer` subtracts amount from sender and adds amount to recipient using a `vm.GethStateDB`.
func transfer(sdb vm.GethStateDB, sender, recipient common.Address, amount *big.Int) {
	utils.MustGetAs[vm.StargazerStateDB](sdb).TransferBalance(sender, recipient, amount)
}
