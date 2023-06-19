#!/bin/bash

send_rpc_requests() {

    txHashes=()
    blockNums=()
    blockHashes=()

    while IFS= read -r tx; do
        txHashes+=("$txHash")
        blockNums+=("$blockNum")
        blockHashes+=("$blockHash")
    done <cachedFormattedOutput.json

    cachedReceiptsByHash=()
    for txh in "${txHashes[@]}"; do
        #		GetTransactionByHash(common.Hash) (*types.TxLookupEntry, error)
        #		// GetReceiptByHash returns the receipts at the given block hash.
        receipt=$(cast rpc eth_getTransactionByHash "$txh" "false")
        cachedReceiptsByHash+=("$receipt")
    done

    cachedBlockByNumber=()
    for bn in "${blockNums[@]}"; do
        #	  GetBlockByNumber(uint64) (*types.Block, error)
        #		// GetBlockByHash returns the block at the given block hash.
        blockByNumber=$(cast rpc eth_getBlockByNumber "$bn" "false")
        cachedBlockByNumber+=("$blockByNumber")
    done

    cachedBlockByHash=()
    for bh in "${blockHashes[@]}"; do
        #		GetBlockByHash(common.Hash) (*types.Block, error)
        #		// GetTransactionByHash returns the transaction lookup entry at the given transaction
        #		// hash.
        blockByHash=$(cast rpc eth_getBlockByHash "$bh" "false")
        cachedBlockByHash+=("$blockByHash")
    done
}
