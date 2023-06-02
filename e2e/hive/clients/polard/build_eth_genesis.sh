#!/bin/bash

# Extract eth_genesis from config/genesis.json
eth_genesis=$(cat config/genesis.json | jq -r '.app_state.evm.params.eth_genesis')

# Remove escape characters and format the JSON
formatted_eth_genesis=$(echo -n "$eth_genesis" | sed 's/\\//g' | jq '.')

# Create a new file called genesis.json with formatted_eth_genesis
echo "$formatted_eth_genesis" > genesis.json

echo "Successfully created ethereum genesis.json"
