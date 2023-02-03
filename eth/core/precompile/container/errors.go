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

package container

import "errors"

var (
	// `ErrOutOfGas` is returned when a precompile container execution runs out of gas.
	ErrOutOfGas = errors.New("out of gas")

	// `ErrIncompatibleStateDB` is returned when an incompatible statedb is passed in.
	ErrIncompatibleStateDB = errors.New("statedb is not compatible with Stargazer")

	// `ErrInvalidContext` is returned when a nil context is passed in.
	ErrInvalidContext = errors.New("the context passed in is nil")

	// `ErrIncompleteMethod` is returned when a `FnAndGas` has missing, or nil, required fields.
	ErrIncompleteMethod = errors.New("incomplete precompile Method")

	// `ErrAbiSigInvalid` is returned when a user-provided ABI signature (`FnAndGas.AbiSig`) does
	// not match the Go-Ethereum style function signatures. Please check
	// core/vm/precompile/container/types.go for more information.
	ErrAbiSigInvalid = errors.New("user-provided ABI signature invalid: ")

	// `ErrMethodNotFound` is returned when the precompile method is not found.
	ErrMethodNotFound = errors.New("precompile method not found in contract ABI")

	// `ErrContainerHasNoMethods` is returned when a stateful container function is invoked but no
	// precompile methods were registered.
	ErrContainerHasNoMethods = errors.New("the stateful precompile has no methods to run")

	// `ErrInvalidInputToPrecompile` is returned when a precompile container receives invalid
	// input.
	ErrInvalidInputToPrecompile = errors.New("input bytes to precompile container are invalid")

	// `ErrWrongContainerFactory` is returned when the wrong precompile container factory is used
	// to build a precompile contract.
	ErrWrongContainerFactory = errors.New("this precompile contract implementation is not implemented")

	// `ErrNoPrecompileMethodForABIMethod` is returned when no precompile method is provided for a
	// corresponding ABI method.
	ErrNoPrecompileMethodForABIMethod = errors.New("this ABI method does not have a corresponding precompile method")
)
