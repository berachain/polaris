// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package txpool

// Mempool metrics.
const (
	MetricKeyCometPrefix = "polaris_cometbft_"

	MetricKeyCometPoolTxs  = "polaris_cometbft_comet_pool_txs"
	MetricKeyCometLocalTxs = "polaris_cometbft_local_txs"

	MetricKeyTimeShouldEject           = "polaris_cometbft_time_should_eject"
	MetricKeyAnteEjectedTxs            = "polaris_cometbft_ante_ejected_txs"
	MetricKeyAnteShouldEjectInclusion  = "polaris_cometbft_ante_should_eject_included"
	MetricKeyAnteShouldEjectExpiredTx  = "polaris_cometbft_ante_should_eject_expired"
	MetricKeyAnteShouldEjectPriceLimit = "polaris_cometbft_ante_should_eject_price_limit"

	MetricKeyTxPoolPending = "polaris_cometbft_txpool_pending"
	MetricKeyTxPoolQueue   = "polaris_cometbft_txpool_queue"

	MetricKeyMempoolFull     = "polaris_cometbft_mempool_full"
	MetricKeyMempoolSize     = "polaris_cometbft_mempool_size"
	MetricKeyMempoolKnownTxs = "polaris_cometbft_mempool_known_txs"

	MetricKeyBroadcastTxs     = "polaris_cometbft_broadcast_txs"
	MetricKeyBroadcastFailure = "polaris_cometbft_broadcast_failure"
	MetricKeyBroadcastRetry   = "polaris_cometbft_broadcast_retry"
)
