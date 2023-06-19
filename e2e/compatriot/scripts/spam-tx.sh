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
jsonArray="["
for ((i=0; i<${#txReceipts[@]}; i++)); do
    jsonArray+="$(printf '%s\n' "${txReceipts[i]}")"
    if [ $i -lt $((${#txReceipts[@]} - 1)) ]; then
        jsonArray+=","
    fi
done
jsonArray+="]"

# write the json array to a file
echo "$jsonArray" | jq '.' > txReceipts.json
