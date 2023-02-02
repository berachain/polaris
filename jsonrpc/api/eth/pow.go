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

import "github.com/ethereum/go-ethereum/common/hexutil"

const (
	// `MethodEthHashrate` is the method name of `eth_hashrate`.
	MethodEthHashrate = "eth_hashrate"
	// `MethodEthMining` is the method name of `eth_mining`.
	MethodEthMining = "eth_mining"
)

// `Hashrate` returns 0 since there is no mining in Tendermint.
func (api *api) Hashrate() hexutil.Uint64 {
	api.logger.Debug(MethodEthHashrate)
	return 0
}

// `Mining` returns false since there is no mining in Tendermint.
func (api *api) Mining() bool {
	api.logger.Debug(MethodEthMining)
	return false
}
