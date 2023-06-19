#!/bin/bash

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

jq -c '.[]' cachedOutput.json >formattedCachedOutput.json

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

#   return the values for use in the send_rpc_requests() function
echo "${txHashes[*]}"
echo "${blockNums[*]}"
echo "${blockHashes[*]}"

rm -rf cachedOutput.json
