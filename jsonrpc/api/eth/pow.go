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

package eth

import (
	"github.com/berachain/stargazer/lib/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// `MethodEthHashrate` is the method name of `eth_hashrate`.
	MethodEthHashrate = "eth_hashrate"
	// `MethodEthMining` is the method name of `eth_mining`.
	MethodEthMining = "eth_mining"
)

// `Hashrate` returns 0 since there is no mining in CometBFT.
func (api *api) Hashrate() hexutil.Uint64 {
	api.logger.Debug(MethodEthHashrate)
	return 0
}

// `Mining` returns false since there is no mining in CometBFT.
func (api *api) Mining() bool {
	api.logger.Debug(MethodEthMining)
	return false
}

// `GetWork` returns nil since there is no mining in CometBFT.
func (api *api) GetWork() ([]hexutil.Bytes, error) {
	return nil, nil
}

// `SubmitWork` returns false since there is no mining in CometBFT.
func (api *api) SubmitWork(nonce hexutil.Uint64, headerHash, mixDigest hexutil.Bytes) bool {
	return false
}

// `SubmitHashrate` returns false since there is no mining in CometBFT.
func (api *api) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	return false
}
