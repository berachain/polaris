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

package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// `maxIndexedArgs` is the maximum number of indexed arguments allowed in an Ethereum event log.
const maxIndexedArgs = 3

type (
	ABI       = abi.ABI
	Argument  = abi.Argument
	Arguments = abi.Arguments
	Event     = abi.Event
)

var (
	MakeTopics  = abi.MakeTopics
	NewEvent    = abi.NewEvent
	NewType     = abi.NewType
	ToCamelCase = abi.ToCamelCase
	JSON        = abi.JSON
)

// `ToMixedCase` converts a under_score formatted string to mixedCase format (camelCase with the
// first letter lowercase). This function is inspired by the geth `abi.ToCamelCase“ function.
func ToMixedCase(input string) string {
	parts := strings.Split(input, "_")
	for i, s := range parts {
		if i > 0 && len(s) > 0 {
			parts[i] = strings.ToUpper(s[:1]) + s[1:]
		}
	}
	return strings.Join(parts, "")
}

// `GetIndexed` extracts indexed arguments from a set of arguments. Will panic if more than 3
// indexed arguments are provided by the inputs ABI.
func GetIndexed(args abi.Arguments) (abi.Arguments, error) {
	var indexed abi.Arguments
	for _, arg := range args {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if len(indexed) > maxIndexedArgs {
		return nil, ErrTooManyIndexedArgs
	}

	return indexed, nil
}
