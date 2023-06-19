#!/bin/bash

# cd "../../"
# mage build
# mage start &

# save the block hash and block number from these transactions
# build historical plugin function calls with the block hash and block number data
# send these requests to the node, save it in a .json file

# shut down the node + nuke the caches

# start the node again
# send the request to the node again
# save these results in a different .json file

# diff the two files

PRIV_KEY=0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306
RECEIVER=0x00000000000000000000000000000000DeaDBeef

# sleep 10 # wait for network to boot up
txReceipts=()

# Send 10 transactions and append JSON objects to the array
for i in {1..10}; do
  txReceipt=$(cast send --private-key $PRIV_KEY $RECEIVER --rpc-url http://127.0.0.1:8545/ --json)
  txReceipts+=("$txReceipt")
done
# Join the array elements with commas and enclose in brackets
json_output="[$(
  IFS=,
  echo "${txReceipts[*]}"
)]"
# Write the JSON output to the file
echo "$json_output" >cachedOutput.json
jq -r '.' cachedOutput.json >temp.json && mv temp.json cachedOutput.json

txHashes=()
blockNums=()
blockHashes=()

jq -c '.[]' cachedOutput.json >tmp.json

while IFS= read -r tx; do
  # do stuff with $tx
  txHash=$(echo "$tx" | jq -r ".transactionHash")
  blockNum=$(echo "$tx" | jq -r ".blockNumber")
  blockHash=$(echo "$tx" | jq -r ".blockHash")
  # Output the extracted value
  txHashes+=("$txHash")
  blockNums+=("$blockNum")
  blockHashes+=("$blockHash")
done <tmp.json

rm tmp.json

cachedTxByHash=()
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

cachedReceiptsByHash=()
#		// StoreBlock stores the given block.
#		StoreBlock(*types.Block) error
#		// StoreReceipts stores the receipts for the given block hash.
#		StoreReceipts(common.Hash, types.Receipts) error
#		// StoreTransactions stores the transactions for the given block hash.
#		StoreTransactions(uint64, common.Hash, types.Transactions) error
