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

package api

// List of JSON-RPC methods
//
// +-----------------------------------------+-----------+
// |                 Method                  | Supported |
// +-----------------------------------------+-----------+
// | web3_clientVersion                      |     Y     |
// | web3_sha3                               |           |
// | net_version                             |     Y     |
// | net_listening                           |     Y     |
// | net_peerCount                           |     Y     |
// | eth_protocolVersion                     |           |
// | eth_syncing                             |           |
// | eth_coinbase                            |           |
// | eth_mining                              |     N     |
// | eth_hashrate                            |     N     |
// | eth_gasPrice                            |           |
// | eth_accounts                            |           |
// | eth_blockNumber                         |     Y     |
// | eth_getBalance                          |           |
// | eth_getStorageAt                        |           |
// | eth_getTransactionCount                 |           |
// | eth_getBlockTransactionCountByHash      |           |
// | eth_getBlockTransactionCountByNumber    |           |
// | eth_getUncleCountByBlockHash            |           |
// | eth_getUncleCountByBlockNumber          |           |
// | eth_getCode                             |           |
// | eth_sign                                |           |
// | eth_signTransaction                     |           |
// | eth_sendTransaction                     |           |
// | eth_sendRawTransaction                  |           |
// | eth_call                                |           |
// | eth_estimateGas                         |           |
// | eth_getBlockByHash                      |           |
// | eth_getBlockByNumber                    |           |
// | eth_getTransactionByHash                |           |
// | eth_getTransactionByBlockHashAndIndex   |           |
// | eth_getTransactionByBlockNumberAndIndex |           |
// | eth_getTransactionReceipt               |           |
// | eth_getUncleByBlockHashAndIndex         |           |
// | eth_getUncleByBlockNumberAndIndex       |           |
// | eth_newFilter                           |           |
// | eth_newBlockFilter                      |           |
// | eth_newPendingTransactionFilter         |           |
// | eth_uninstallFilter                     |           |
// | eth_getFilterChanges                    |           |
// | eth_getFilterLogs                       |           |
// | eth_getLogs                             |           |
// | eth_getWork                             |     N     |
// | eth_submitWork                          |     N     |
// | eth_submitHashrate                      |     N     |
// +-----------------------------------------+-----------+
