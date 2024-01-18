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

package eth

import (
	"fmt"

	"github.com/berachain/polaris/lib/utils"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/ethereum/go-ethereum/metrics"
)

// Available geth txpool metrics
const (
	// // Metrics for the pending pool
	// MetricKeyTxpoolPendingDiscardMeter = "txpool/pending/discard"
	// MetricKeyPendingReplaceMeter       = "txpool/pending/replace"
	// MetricKeyPendingRateLimitMeter     = "txpool/pending/ratelimit" // Dropped due to rate limiting
	// MetricKeyPendingNofundsMeter       = "txpool/pending/nofunds"   // Dropped due to out-of-funds

	// // Metrics for the queued pool
	// MetricKeyTxpoolQueuedDiscardMeter   = "txpool/queued/discard"
	// MetricKeyTxpoolQueuedReplaceMeter   = "txpool/queued/replace"
	// MetricKeyTxpoolQueuedRateLimitMeter = "txpool/queued/ratelimit" // Dropped due to rate limiting
	// MetricKeyTxpoolQueuedNofundsMeter   = "txpool/queued/nofunds"   // Dropped due to out-of-funds
	// MetricKeyTxpoolQueuedEvictionMeter  = "txpool/queued/eviction"  // Dropped due to lifetime

	// // General tx metrics
	// MetricKeyTxpoolKnownTxMeter       = "txpool/known"
	// MetricKeyTxpoolValidTxMeter       = "txpool/valid"
	// MetricKeyTxpoolInvalidTxMeter     = "txpool/invalid"
	// MetricKeyTxpoolUnderpricedTxMeter = "txpool/underpriced"
	// MetricKeyTxpoolOverflowedTxMeter  = "txpool/overflowed"

	// // throttleTxMeter counts how many transactions are rejected
	// // due to too-many-changes between txpool reorgs.
	// MetricKeyTxpoolThrottleTxMeter = "txpool/throttle"
	// // reorgDurationTimer measures how long time a txpool reorg takes.
	// MetricKeyTxpoolReorgDurationTimer = "txpool/reorgtime"
	// // dropBetweenReorgHistogram counts how many drops we experience between two reorg runs.
	// // It is expected that this number is pretty low, since txpool reorgs happen very frequently.
	// MetricKeyTxpoolDropBetweenReorgHistogram = "txpool/dropbetweenreorg"

	MetricKeyTxpoolPendingGauge = "txpool/pending"
	MetricKeyTxpoolQueuedGauge  = "txpool/queued"
	MetricKeyTxpoolLocalGauge   = "txpool/local"
	MetricKeyTxpoolSlotsGauge   = "txpool/slots"

	// MetricKeyTxpoolReheapTimer = "txpool/reheap"
)

type Metrics struct {
	gauges map[string]metrics.Gauge
}

func NewMetrics() *Metrics {
	return &Metrics{
		gauges: make(map[string]metrics.Gauge),
	}
}

func (m *Metrics) Register(name string) error {
	gauge, ok := utils.GetAs[metrics.Gauge](metrics.DefaultRegistry.Get(name))
	if !ok {
		return fmt.Errorf("could not get eth meter for %s", name)
	}
	m.gauges[name] = gauge
	return nil
}

func (m *Metrics) Report() {
	for name, gauge := range m.gauges {
		snapshot := gauge.Snapshot()
		telemetry.SetGauge(float32(snapshot.Value()), name)
	}
}

var TxPoolMetrics = []string{
	MetricKeyTxpoolPendingGauge,
	MetricKeyTxpoolQueuedGauge,
}
