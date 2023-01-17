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

import "errors"

var (
	// `ErrIncompleteFnAndGas` is returned when a `FnAndGas` has missing, or nil, fields.
	ErrIncompleteFnAndGas = errors.New("incomplete FnAndGas passed in for precompile")

	// `ErrAbiSigInvalid` is returned when a user-provided ABI signature (`FnAndGas.AbiSig`) does
	// not match the Go-Ethereum style function signatures. Please check
	// core/vm/precompile/types.go for more information.
	ErrAbiSigInvalid = errors.New("user-provided ABI signature invalid: ")

	// `ErrEthEventNotRegistered` is returned when an incoming Cosmos event is not mapped to any
	// registered Ethereum event.
	ErrEthEventNotRegistered = errors.New("the Ethereum event was not registered for Cosmos event")
)
