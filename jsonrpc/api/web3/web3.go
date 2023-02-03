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

package web3

import (
	"github.com/berachain/stargazer/lib/common/hexutil"
	"github.com/berachain/stargazer/lib/crypto"
	libtypes "github.com/berachain/stargazer/lib/types"
	"go.uber.org/zap/zapcore"
)

const (
	// `MethodClientVersion` is the method name of `ClientVersion`.
	MethodClientVersion = "web3_clientVersion"
	// `MethodSha3` is the method name of `Sha3`.
	MethodSha3 = "web3_sha3"
)

// `api` is the Web3 API.
type api struct {
	logger libtypes.Logger[zapcore.Field]
}

// `NewAPI` returns a new `api` object.
func NewAPI(logger libtypes.Logger[zapcore.Field]) *api { //nolint: revive // by design.
	return &api{
		logger,
	}
}

// `ClientVersion` returns the client version.
func (api *api) ClientVersion() string {
	api.logger.Debug(MethodClientVersion)
	return "stargazer" // TODO: implement
}

// `Sha3` returns the keccak-256 hash of the supplied input.
func (api *api) Sha3(input string) hexutil.Bytes {
	api.logger.Debug(MethodSha3)
	return crypto.Keccak256(hexutil.Bytes(input))
}
