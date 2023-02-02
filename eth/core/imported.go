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
	"github.com/ethereum/go-ethereum/core"
)

type (
	// `ExecutionResult` is the result of executing a transaction.
	ExecutionResult = core.ExecutionResult

	// `Message` contains data used ype used to execute transactions.
	Message = core.Message
)

var (
	// `EthIntrinsicGas` returns the intrinsic gas required to execute a transaction.
	EthIntrinsicGas = core.IntrinsicGas

	// `NewEVMTxContext` creates a new context for use in the EVM.
	NewEVMTxContext = core.NewEVMTxContext
)

var (
	// `ErrIntrinsicGas` is the error returned when the intrinsic gas is higher than the gas limit.
	ErrIntrinsicGas = core.ErrIntrinsicGas

	// `ErrInsufficientFundsForTransfer` is the error returned when the account does not have enough funds to transfer.
	ErrInsufficientFundsForTransfer = core.ErrInsufficientFundsForTransfer

	// `ErrInsufficientFunds` is the error returned when the account does not have enough funds to execute the transaction.
	ErrInsufficientFunds = core.ErrInsufficientFunds

	// `ErrInsufficientBalanceForGas` is the error return when gas required to execute a transaction overflows.
	ErrGasUintOverflow = core.ErrGasUintOverflow
)
