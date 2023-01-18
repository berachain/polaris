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

// `Method` is a type of function that a stateful precompiled contract should implement.
type Method func(
	ctx sdk.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (ret []any, err error)

// `getFuncName` uses `reflect` and `runtime` to get the Go function's name.
func (pfn Method) getFuncName() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(pfn).Pointer()).Name()
	if brokenUpName := strings.Split(fullName, "."); len(brokenUpName) > 1 {
		return brokenUpName[1]
	}
	return fullName
}

// `PrecompileMethod` is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract function.
type PrecompileMethod struct {
	// `AbiSig` returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256".
	AbiSig string

	// `AbiMethod` is the ABI `Method` struct corresponding to this precompile function. NOTE: this
	// field should be left empty (as nil) as this will automatically be populated by the
	// corresponding interface's ABI.
	AbiMethod *abi.Method

	// `Execute` is the function which will execute the logic of the precompile function.
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
	// ensure all fields are nonempty
	if len(pm.AbiSig) == 0 || pm.AbiMethod != nil || pm.Execute == nil || pm.RequiredGas == 0 {
		return ErrIncompleteFnAndGas
	}

	// validate user-defined abi signature (AbiSig) according to geth ABI signature definition
	// check only 1 `(` exists in the string
	nameAndArgs := strings.Split(pm.AbiSig, "(")
	if len(nameAndArgs) != 2 { //nolint:gomnd // this constant, 2, will never change.
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"function %s does not contain exactly 1 '('",
			pm.Execute.getFuncName(),
		)
	}
	// check that the function name is valid according to Solidity
	if name := nameAndArgs[0]; !funcNameRegex.MatchString(name) {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"function %s does not have a valid function name",
			pm.Execute.getFuncName(),
		)
	}
	// check that only 1 `)` exists and its the last character
	args := strings.Split(nameAndArgs[1], ")")
	if len(args) != 2 || len(args[1]) > 0 {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"function %s does not does not end with 1 ')'",
			pm.Execute.getFuncName(),
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
				"function %s has incorrect argument types",
				pm.Execute.getFuncName(),
			)
		}
	}

	// TODO: validate `AbiMethod` ?

	return nil
}

// `PrecompileMethods` is a type that represents a list of functions and gas. This is what
// stateful precompiled contract should expose.
type PrecompileMethods []*PrecompileMethod
