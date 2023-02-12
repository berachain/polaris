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
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/common/hexutil"
)

// `CompiliedContract` is a contract that has been compiled.
type CompiliedContract struct {
	ABI ABI
	Bin hexutil.Bytes
}

// `BuildCompiledContract` builds a `CompiledContract` from an ABI string and a bytecode string.
func BuildCompiledContract(abiStr, bytecode string) CompiliedContract {
	var parsedAbi ABI
	if err := parsedAbi.UnmarshalJSON([]byte(abiStr)); err != nil {
		panic(err)
	}
	return CompiliedContract{
		ABI: parsedAbi,
		Bin: common.Hex2Bytes(bytecode),
	}
}
