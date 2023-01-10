package vm

import (
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/berachain/stargazer/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Fn func(
	ctx sdk.Context,
	evm StargazerEVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) (ret []any, err error)

// uses reflect, runtime to get the function name.
func (pfn Fn) Name() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(pfn).Pointer()).Name()
	if brokenUpName := strings.Split(fullName, "."); len(brokenUpName) > 1 {
		return brokenUpName[1]
	}
	return fullName
}

type FnAndGas struct {
	// AbiSig returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256"
	AbiSig string

	Func        Fn
	RequiredGas uint64
}

var (
	funcNameRegex = regexp.MustCompile(`^[a-zA-Z_$]{1,}[a-zA-Z0-9_$]*`)
	typeRegex     = regexp.MustCompile(`^[a-z]+[0-9]*$`)
)

// RequireValid panics if this PrecompileFnAndGas has invalid fields.
func (fg *FnAndGas) RequireValid() {
	// ensure all fields are nonempty
	if len(fg.AbiSig) == 0 || fg.Func == nil || fg.RequiredGas == 0 {
		panic("incomplete PrecompileFnAndGas passed in for precompile")
	}

	// validate user-defined abi signature (AbiSig) according to geth ABI signature definition
	// check only 1 `(` exists in the string
	splitOnBracketsSize := 2
	nameAndArgs := strings.Split(fg.AbiSig, "(")
	if len(nameAndArgs) != splitOnBracketsSize {
		panic(
			fmt.Sprintf(
				"user-provided ABI signature for function %s does not contain exactly 1 '('",
				fg.Func.Name(),
			),
		)
	}
	// check that the function name is valid according to Solidity
	if name := nameAndArgs[0]; !funcNameRegex.MatchString(name) {
		panic(
			fmt.Sprintf(
				"user-provided ABI signature for function %s does not have a valid function name",
				fg.Func.Name(),
			),
		)
	}
	// check that only 1 `)` exists and its the last character
	args := strings.Split(nameAndArgs[1], ")")
	if len(args) != 2 || len(args[1]) > 0 {
		panic(
			fmt.Sprintf(
				"user-provided ABI signature for function %s does not does not end with 1 ')'",
				fg.Func.Name(),
			),
		)
	}
	// if no args are provided, sig is valid
	if len(args[0]) == 0 {
		return
	}
	// check that each provided type is valid if some args are provided
	types := strings.Split(args[0], ",")
	for _, t := range types {
		if len(t) == 0 || !typeRegex.MatchString(t) {
			panic(
				fmt.Sprintf(
					"user-provided ABI signature for function %s has incorrect argument types",
					fg.Func.Name(),
				),
			)
		}
	}
}

type FnsAndGas []*FnAndGas

type Getter func(addr common.Address) (PrecompiledContract, bool)

// StatefulContrats are defined in-memory.
type StatefulContract interface {
	PrecompiledContract
	GetFunctionsAndGas() FnsAndGas
}

// FactoryContract are defined via a kvstore.
type FactoryContract interface {
	StatefulContract
	Name() string
}
