// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
