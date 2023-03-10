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

package node

import (
	"time"

	"github.com/rs/zerolog/log"

	"pkg.berachain.dev/polaris/playground/pkg/chain"
)

// `Runner` is the interface for the node runner.
type Runner interface {
	Start() error
}

// `runner` is the main node runner.
type runner struct {
	blocktime  time.Duration
	rpcService RPCService
	stop       chan struct{}
	chain      *chain.Playground
}

// `NewRunner` creates a new node runner.
func NewRunner(blocktime time.Duration) Runner {
	// Setup RPC
	rpcService := NewRPCService()

	// Setup Mempool
	mempool := NewMempool()
	return &runner{
		blocktime:  blocktime,
		rpcService: rpcService,
		stop:       make(chan struct{}),
		chain: chain.NewPlayground(
			mempool,
		),
	}
}

// `Start` starts the node.
func (r *runner) Start() error {
	for {
		select {
		case <-r.stop:
			log.Info().Msg("chain shutting down")
			close(r.stop)
			return nil
		case <-time.After(r.blocktime):
			log.Error().Msg("producing block")
			if err := r.chain.ProduceBlock(); err != nil {
				log.Error().Err(err).Msg("failed to produce block")
				return err
			}
		}
	}
}

// `Stop` stops the node.
func (r *runner) Stop() {
	r.stop <- struct{}{}
}
