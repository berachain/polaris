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
	"encoding/json"

	"github.com/berachain/stargazer/eth/params"
	"github.com/cosmos/gogoproto/types"
)

// `DefaultParams` contains the default values for all parameters.
func DefaultParams() *Params {
	return &Params{
		EvmDenom:    "abera",
		ExtraEIPs:   []int64{},
		ChainConfig: chainConfig{*params.DefaultChainConfig}.ToProtoStruct(),
	}
}

// `chainConfig` is a wrapper around `params.ChainConfig` to provide compatibility
// with gogoproto types.Struct.
type chainConfig struct {
	params.ChainConfig
}

// `ToProtoStruct` marshals the `params.ChainConfig` to JSON and then unmarshals it to
// a `params.ChainConfig`.
func (c chainConfig) ToProtoStruct() types.Struct {
	// Marshal to JSON
	bz, err := json.Marshal(params.DefaultChainConfig)
	if err != nil {
		panic(err)
	}

	// Unmarshal to the protobuf struct.
	var chainConfig types.Struct
	err = chainConfig.Unmarshal(bz)
	if err != nil {
		panic(err)
	}
	return chainConfig
}
