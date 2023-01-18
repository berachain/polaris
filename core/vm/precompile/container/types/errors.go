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

package types

import "errors"

var (
	// `ErrIncompleteFnAndGas` is returned when a `FnAndGas` has missing, or nil, fields.
	ErrIncompleteFnAndGas = errors.New("incomplete FnAndGas passed in for precompile")

	// `ErrAbiSigInvalid` is returned when a user-provided ABI signature (`FnAndGas.AbiSig`) does
	// not match the Go-Ethereum style function signatures. Please check
	// core/vm/precompile/container/types.go for more information.
	ErrAbiSigInvalid = errors.New("user-provided ABI signature invalid: ")

	// `ErrEthEventNotRegistered` is returned when an incoming Cosmos event is not mapped to any
	// registered Ethereum event.
	ErrEthEventNotRegistered = errors.New("this Ethereum event was not registered for Cosmos event")

	// `ErrEthEventAlreadyRegistered` is returned when an already registered Ethereum event is
	// being registered again.
	ErrEthEventAlreadyRegistered = errors.New("this Ethereum event is already registered")

	// `ErrStateDBNotSupported` is returned when the state DB is not compatible for running
	// stateful precompiles.
	ErrStateDBNotSupported = errors.New("given StateDB is not compatible for running stateful precompiles")

	// `ErrPrecompileMethodNotFound` is returned when the Precompile method is not found.
	ErrPrecompileMethodNotFound = errors.New("precompile method not found in contract ABI")

	// `ErrPrecompileHasNoMethods` is returned when a stateful container function is invoked but no
	// precompile methods were registered.
	ErrPrecompileHasNoMethods = errors.New("the stateful precompile has no methods to run")
)
