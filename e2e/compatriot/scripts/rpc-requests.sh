#!/bin/bash

# retrieve the txHashes, blockHashes and blockNums from the cachedOutput.json file
txHashes=$(jq -r '.[].transactionHash' txReceipts.json)
blockHashes=$(jq -r '.[].blockHash' txReceipts.json)
blockNums=$(jq -r '.[].blockNumber' txReceipts.json)

rm -rf txReceipts.json # we don't need it anymore, remove it

# gather receipts for each txHash and add response to array
cachedReceiptsByHash=()
for txh in $txHashes; do
    #		GetTransactionByHash(common.Hash) (*types.TxLookupEntry, error)
    #		// GetReceiptByHash returns the receipts at the given block hash.
    receipt=$(cast rpc eth_getTransactionByHash "$txh")
    cachedReceiptsByHash+=("$receipt")
done

# gather blocks for each blockNum and add response to array
cachedBlockByNumber=()
for bn in $blockNums; do
    #	  GetBlockByNumber(uint64) (*types.Block, error)
    #		// GetBlockByHash returns the block at the given block hash.
    blockByNumber=$(cast rpc eth_getBlockByNumber "$bn" "false")
    cachedBlockByNumber+=("$blockByNumber")
done

# gather blocks for each blockHash and add response to array
cachedBlockByHash=()
for bh in $blockHashes; do
    #		GetBlockByHash(common.Hash) (*types.Block, error)
    #		// GetTransactionByHash returns the transaction lookup entry at the given transaction
    #		// hash.
    blockByHash=$(cast rpc eth_getBlockByHash "$bh" "false")
    cachedBlockByHash+=("$blockByHash")
done

# Join the array elements with commas and enclose in brackets
receiptsByHashArray="["
for ((i=0; i<${#cachedReceiptsByHash[@]}; i++)); do
    receiptsByHashArray+="$(printf '%s\n' "${cachedReceiptsByHash[i]}")"
    if [ $i -lt $((${#cachedReceiptsByHash[@]} - 1)) ]; then
        receiptsByHashArray+=","
    fi
done
receiptsByHashArray+="]"
# write the json array to a file
echo "$receiptsByHashArray" | jq '.' > receiptsByHash.json








echo "________________________________________"
# echo "cachedBlockByNumber: ${cachedBlockByNumber[@]}" | jq .
# echo "________________________________________"
# echo "cachedBlockByHash: ${cachedBlockByHash[@]}" | jq .
# echo "________________________________________" |
