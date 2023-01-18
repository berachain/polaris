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

import (
	"math/big"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"cosmossdk.io/errors"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/**
 * 	Welcome to Stateful Precompiled Contracts!
 *	To build a stateful precompile, you must follow these steps:
 *		1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *		2) Build a Go precompile contract, which implements the interface's methods.
 *   		A) This precompile contract should expose the ABI's `Methods`, which can be generated
 *      	   via Go-Ethereum's abi package. These methods are of type `abi.Method`.
 *   		B) This precompile contract should also expose the `PrecompileMethod`s. A
 *             `PrecompileMethod` includes the `Method`, which is the direct implementation of a
 *             corresponding ABI method, the `Method`'s `RequiredGas`, and the ABI signature. Do
 *             NOT provide the `abiMethod` as this will be auto-populated.
 **/

// `Method` is a type of function that stateful precompiled contract will implement. Each `Method`
// should directly correspond to an ABI method.
type Method func(
	ctx sdk.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (ret []any, err error)

// `getName` uses `reflect` and `runtime` to get the Go function's name.
func (pfn Method) getName() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(pfn).Pointer()).Name()
	if brokenUpName := strings.Split(fullName, "."); len(brokenUpName) > 1 {
		return brokenUpName[1]
	}
	return fullName
}

// `PrecompileMethod` is a struct that contains the required information for the EVM to execute a
// stateful precompiled contract method.
type PrecompileMethod struct {
	// `AbiSig` returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256".
	AbiSig string

	// `AbiMethod` is the ABI `Method` struct corresponding to this precompile method. NOTE: this
	// field should be left empty (as nil) as this will automatically be populated by the
	// corresponding interface's ABI.
	AbiMethod *abi.Method

	// `Execute` is the function which will execute the logic of the implemented method.
	Execute Method

	// `RequiredGas` is the amount of gas (as a `uint64`) used up by the execution of `Execute`.
	RequiredGas uint64
}

var (
	funcNameRegex = regexp.MustCompile(`^[a-zA-Z_$]{1,}[a-zA-Z0-9_$]*`)
	typeRegex     = regexp.MustCompile(`^[a-z]+[0-9]*$`)
)

// `ValidateBasic` returns an error if this a precompile `PrecompileMethod` has invalid fields.
func (pm *PrecompileMethod) ValidateBasic() error {
	// ensure all required fields are nonempty
	if len(pm.AbiSig) == 0 || pm.AbiMethod != nil || pm.Execute == nil || pm.RequiredGas == 0 {
		return ErrIncompleteFnAndGas
	}

	// validate user-defined abi signature (AbiSig) according to geth ABI signature definition
	// check only 1 `(` exists in the string
	nameAndArgs := strings.Split(pm.AbiSig, "(")
	if len(nameAndArgs) != 2 { //nolint:gomnd // this constant -- 2 -- will never change.
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not contain exactly 1 '('",
			pm.Execute.getName(),
		)
	}
	// check that the method name is valid according to Solidity
	if name := nameAndArgs[0]; !funcNameRegex.MatchString(name) {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not have a valid method name",
			pm.Execute.getName(),
		)
	}
	// check that only 1 `)` exists and its the last character
	args := strings.Split(nameAndArgs[1], ")")
	if len(args) != 2 || len(args[1]) > 0 {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not does not end with 1 ')'",
			pm.Execute.getName(),
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
				pm.Execute.getName(),
			)
		}
	}

	return nil
}

// `PrecompileMethods` is a type that represents a list of precompile methods. This is what a
// stateful precompiled contract implementation should expose.
type PrecompileMethods []*PrecompileMethod
