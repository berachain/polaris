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

package block

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth/core"
	coretypes "github.com/berachain/polaris/eth/core/types"
	errorslib "github.com/berachain/polaris/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// prevHeaderHashes is the number of previous header hashes being stored on chain.
const prevHeaderHashes = 256

// GetHeaderByNumber returns the header at the given height, using the plugin's query context.
//
// GetHeaderByNumber implements core.BlockPlugin.
func (p *plugin) GetHeaderByNumber(number uint64) (*ethtypes.Header, error) {
	bz, err := p.readHeaderBytes(number)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeaderByNumber: failed to readHeaderBytes")
	}
	if bz == nil {
		return nil, core.ErrHeaderNotFound
	}

	header, err := coretypes.UnmarshalHeader(bz)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeaderByNumber: failed to unmarshal")
	}

	if header.Number.Uint64() > number {
		return nil, errorslib.Wrapf(
			err,
			"GetHeader: header number mismatch, requested %d, got %d ",
			number, header.Number.Uint64(),
		)
	}

	return header, nil
}

// GetHeaderByHash returns the header specified by the given block hash
//
// GetHeaderByHash implements core.BlockPlugin.
func (p *plugin) GetHeaderByHash(hash common.Hash) (*ethtypes.Header, error) {
	numBz := p.ctx.MultiStore().GetKVStore(p.storekey).Get(hash.Bytes())
	if numBz == nil {
		return nil, core.ErrHeaderNotFound
	}
	return p.GetHeaderByNumber(new(big.Int).SetBytes(numBz).Uint64())
}

// StoreHeader implements core.BlockPlugin.
func (p *plugin) StoreHeader(header *ethtypes.Header) error {
	headerHash := header.Hash()
	headerBz, err := coretypes.MarshalHeader(header)
	if err != nil {
		return errorslib.Wrap(err, "SetHeader: failed to marshal header")
	}

	blockHeight := header.Number.Int64()
	if blockHeight != p.ctx.BlockHeight() {
		return fmt.Errorf(
			"StoreHeader: block height mismatch, got %d, expected %d",
			blockHeight, p.ctx.BlockHeight(),
		)
	}

	// write genesis header
	if blockHeight == 0 {
		return p.writeGenesisHeaderBytes(headerHash, headerBz)
	}

	kvstore := p.ctx.MultiStore().GetKVStore(p.storekey)
	// set header key
	kvstore.Set([]byte{types.HeaderKey}, headerBz)

	// rotate previous header hashes
	if pruneHeight := blockHeight - prevHeaderHashes; pruneHeight > 0 {
		hashKey := headerHashKeyForHeight(pruneHeight)
		pruneHash := kvstore.Get(hashKey)
		kvstore.Delete(hashKey)
		kvstore.Delete(pruneHash)
	}
	kvstore.Set(headerHashKeyForHeight(blockHeight), headerHash.Bytes())
	kvstore.Set(headerHash.Bytes(), header.Number.Bytes())

	return nil
}

// readHeaderBytes reads the header at the given height, using the plugin's query context for
// non-genesis blocks.
func (p *plugin) readHeaderBytes(number uint64) ([]byte, error) {
	// if number requested is 0, get the genesis block header
	if number == 0 {
		return p.readGenesisHeaderBytes(), nil
	}

	// try fetching the query context for a historical block header
	if p.getQueryContext == nil {
		return nil, errors.New("GetHeader: getQueryContext is nil")
	}

	// TODO: ensure we aren't differing from geth / hiding errors here.
	// TODO: the GTE may be hiding a larger issue with the timing of the NewHead channel stuff.
	// Investigate and hopefully remove this GTE.
	if number > uint64(p.ctx.BlockHeight()) {
		// cannot retrieve future block header
		number = uint64(p.ctx.BlockHeight())
	}

	ctx, err := p.getQueryContext()(int64(number), false)
	if err != nil {
		return nil, errorslib.Wrap(err, "GetHeader: failed to use query context")
	}

	// Unmarshal the header at IAVL height from its context kv store.
	return ctx.MultiStore().GetKVStore(p.storekey).Get([]byte{types.HeaderKey}), nil
}

// writeGenesisHeaderBytes writes the genesis header to the kvstore.
//
//	GenesisHeaderKey --> Header bytes
//	Header Hash      --> 0
func (p *plugin) writeGenesisHeaderBytes(headerHash common.Hash, headerBz []byte) error {
	p.ctx.MultiStore().GetKVStore(p.storekey).Set([]byte{types.GenesisHeaderKey}, headerBz)
	p.ctx.MultiStore().GetKVStore(p.storekey).Set(headerHash.Bytes(), new(big.Int).Bytes())
	return nil
}

// readGenesisHeaderBytes returns the header bytes at the genesis key.
func (p *plugin) readGenesisHeaderBytes() []byte {
	return p.ctx.MultiStore().GetKVStore(p.storekey).Get([]byte{types.GenesisHeaderKey})
}
