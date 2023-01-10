// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type (
	ABI          = abi.ABI
	Argument     = abi.Argument
	Arguments    = abi.Arguments
	Event        = abi.Event
	FunctionType = abi.FunctionType
	Method       = abi.Method
	Type         = abi.Type
)

const (
	IntTy        = abi.IntTy
	UintTy       = abi.UintTy
	BoolTy       = abi.BoolTy
	StringTy     = abi.StringTy
	SliceTy      = abi.SliceTy
	ArrayTy      = abi.ArrayTy
	TupleTy      = abi.TupleTy
	AddressTy    = abi.AddressTy
	FixedBytesTy = abi.FixedBytesTy
	BytesTy      = abi.BytesTy
	HashTy       = abi.HashTy
	FixedPointTy = abi.FixedPointTy
	FunctionTy   = abi.FunctionTy
)

var (
	Function    = abi.Function
	NewMethod   = abi.NewMethod
	NewEvent    = abi.NewEvent
	NewType     = abi.NewType
	ToCamelCase = abi.ToCamelCase
	MakeTopics  = abi.MakeTopics
)

// converts under_score formatted string to mixedCase (first letter lowercase, CamelCase)
// works similarly as geth abi.ToCamelCase function.
func ToMixedCase(input string) string {
	parts := strings.Split(input, "_")
	for i, s := range parts {
		if i > 0 && len(s) > 0 {
			parts[i] = strings.ToUpper(s[:1]) + s[1:]
		}
	}
	return strings.Join(parts, "")
}
