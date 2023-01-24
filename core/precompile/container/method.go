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

	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/types/abi"
)

/**
 * 	Welcome to Stateful Precompiled Contracts! To build a stateful precompile, you must follow
 *  these steps:
 *	  1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *	  2) Build a Go precompile contract, which implements the interface's methods.
 *       A) This precompile contract should expose the ABI's `Methods`, which can be generated via
 *          Go-Ethereum's abi package. These methods are of type `abi.Method`.
 *   	 B) This precompile contract should also expose the `Method`s. A `Method` includes the
 *          `Executable`, which is the direct implementation of a corresponding ABI method, the
 *          `Executable`'s `RequiredGas`, and the ABI signature. Do NOT provide the `abiMethod` as
 *          this field will be automatically populated.
 **/

// `funcNamePart` is the part of a runtime function name that is of relevance.
const funcNamePart = 2

// `Executable` is a type of function that stateful precompiled contract will implement. Each
// `Executable` should directly correspond to an ABI method.
type executable func(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (ret []any, err error)

// `GetName` uses `reflect` and `runtime` to get the Go function's name.
func (e executable) getName() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(e).Pointer()).Name()
	if brokenUpName := strings.Split(fullName, "."); len(brokenUpName) > funcNamePart {
		return brokenUpName[funcNamePart]
	}
	return fullName
}

type Method struct {
	abiMethod   *abi.Method
	abiSig      string
	requiredGas uint64
	execute     executable
}

var (
	funcNameRegex = regexp.MustCompile(`^[a-zA-Z_$]{1,}[a-zA-Z0-9_$]*`)
	typeRegex     = regexp.MustCompile(`^[a-z]+[0-9]*$`)
)

// `ValidateBasic` returns an error if this a precompile `Method` has invalid fields.
func (m *Method) ValidateBasic() error {
	// ensure all required fields are nonempty
	if len(m.abiSig) == 0 || m.abiMethod != nil || m.execute == nil {
		return ErrIncompleteMethod
	}

	// validate user-defined abi signature (abiSig) according to geth ABI signature definition
	// check only 1 `(` exists in the string
	nameAndArgs := strings.Split(m.abiSig, "(")
	if len(nameAndArgs) != 2 { //nolint:gomnd // the constant 2 will never change.
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not contain exactly 1 '('",
			m.execute.getName(),
		)
	}
	// check that the method name is valid according to Solidity
	if name := nameAndArgs[0]; !funcNameRegex.MatchString(name) {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not have a valid method name",
			m.execute.getName(),
		)
	}
	// check that only 1 `)` exists and its the last character
	args := strings.Split(nameAndArgs[1], ")")
	if len(args) != 2 || len(args[1]) > 0 {
		return errors.Wrapf(
			ErrAbiSigInvalid,
			"%s does not does not end with 1 ')'",
			m.execute.getName(),
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
				m.execute.getName(),
			)
		}
	}

	return nil
}

type Methods []*Method
