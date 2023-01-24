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

import (
	"context"
	"math/big"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/types/abi"
)

/**
 * 	Welcome to Stateful Precompiled Contracts! To build a stateful precompile, you must implement
 *  the `StatefulPrecompileImpl` interface in `interfaces.go`; below are the suggested steps to
 *  follow:
 *	  1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *	  2) Build a Go precompile contract, which implements the interface's methods.
 *       A) This precompile contract should expose the ABI's `Methods`, which can be generated via
 *          Go-Ethereum's abi package. These methods are of type `abi.Method`.
 *   	 B) This precompile contract should also expose the `Method`s. A `Method` includes the
 *          `executable`, which is the direct implementation of a corresponding ABI method, the
 *          `executable`'s `RequiredGas`, and the ABI signature. Do NOT provide the `AbiMethod` as
 *          this field will be automatically populated.
 **/

// `funcNamePart` is the part of a runtime function name that is of relevance.
const funcNamePart = 2

// `Executable` is a type of function that stateful precompiled contract will implement. Each
// `Executable` should directly correspond to an ABI method.
type Executable func(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (ret []any, logs []*coretypes.Log, err error)

// `GetName` uses `reflect` and `runtime` to get the Go function's name.
func (e Executable) getName() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(e).Pointer()).Name()
	if brokenUpName := strings.Split(fullName, "."); len(brokenUpName) > funcNamePart {
		return brokenUpName[funcNamePart]
	}
	return fullName
}

// `Method` is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract method.
type Method struct {
	// `AbiMethod` is the ABI `Methods` struct corresponding to this precompile's executable. NOTE:
	// this field should be left empty (as nil) as this will automatically be populated by the
	// corresponding interface's ABI.
	AbiMethod *abi.Method

	// `AbiSig` returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256".
	AbiSig string

	// `Execute` is the precompile's executable which will execute the logic of the implemented
	// ABI method.
	Execute Executable

	// `RequiredGas` is the amount of gas (as a `uint64`) used up by the execution of `Execute`.
	// This field is optional; if left empty, the precompile's executable should consume gas using
	// the native gas meter.
	RequiredGas uint64
}

var (
	funcNameRegex = regexp.MustCompile(`^[a-zA-Z_$]{1,}[a-zA-Z0-9_$]*`)
	typeRegex     = regexp.MustCompile(`^[a-z]+[0-9]*$`)
)

// `ValidateBasic` returns an error if this a precompile `Method` has invalid fields.
func (m *Method) ValidateBasic() error {
	// ensure all required fields are nonempty
	if len(m.AbiSig) == 0 || m.AbiMethod != nil || m.Execute == nil {
		return ErrIncompleteMethod
	}

	// validate user-defined abi signature (AbiSig) according to geth ABI signature definition
	// check only 1 `(` exists in the string
	nameAndArgs := strings.Split(m.AbiSig, "(")
	if len(nameAndArgs) != 2 { //nolint:gomnd // the constant 2 will never change.
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not contain exactly 1 '('",
			m.Execute.getName(),
		)
	}
	// check that the method name is valid according to Solidity
	if name := nameAndArgs[0]; !funcNameRegex.MatchString(name) {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not have a valid method name",
			m.Execute.getName(),
		)
	}
	// check that only 1 `)` exists and its the last character
	args := strings.Split(nameAndArgs[1], ")")
	if len(args) != 2 || len(args[1]) > 0 {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not does not end with 1 ')'",
			m.Execute.getName(),
		)
	}
	// if no args are provided, sig is valid
	if len(args[0]) == 0 {
		return nil
	}
	// check that each provided type is valid if some args are provided
	types := strings.Split(args[0], ",")
	for _, t := range types {
		if len(t) == 0 || !typeRegex.MatchString(t) {
			return errors.Wrapf(
				ErrAbiSigInvalid,
				"%s has incorrect argument types",
				m.Execute.getName(),
			)
		}
	}

	return nil
}

// `Methods` is a type that represents a list of precompile methods. This is what a stateful
// precompiled contract implementation should expose.
type Methods []*Method
