#!/bin/bash
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

# Constants
KEYS[0]="dev0"
CHAINID="polaris-2061"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="secp256k1"
LOGLEVEL="info"
HOMEDIR="/"
ETH_GENESIS_JSON="genesis.json"
ETH_GENESIS_SOURCE=$(cat genesis.json)
CONFIG_TOML=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# Used to exit on first error (any non-zero exit code)
set -e

 # Set moniker and chain-id (Moniker can be anything, chain-id must be an integer)
polard init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

# Set client config
polard config set client keyring-backend $KEYRING --home "$HOMEDIR"
polard config set client chain-id "$CHAINID" --home "$HOMEDIR"

# If keys exist they should be deleted
for KEY in "${KEYS[@]}"; do
    polard keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
done

# Update cosmos genesis.json by copying the ethereum genesis.json into the cosmos genesis.
temp_genesis=$(cat $GENESIS)
updated_genesis=$(echo "$temp_genesis" | jq --argjson eth_gen "$ETH_GENESIS_SOURCE" '.app_state["evm"] = $eth_gen')
echo "$updated_genesis" > "$GENESIS"


# Change parameter token denominations to abera
jq '.app_state["staking"]["params"]["bond_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["mint"]["params"]["mint_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

# Allocate genesis accounts (cosmos formatted addresses)
for KEY in "${KEYS[@]}"; do
    polard genesis add-genesis-account $KEY 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
done

# Sign genesis transaction
polard genesis gentx ${KEYS[0]} 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"

# Collect genesis tx
polard genesis collect-gentxs --home "$HOMEDIR"

# Run this to ensure everything worked and that the genesis file is setup correctly
polard genesis validate-genesis --home "$HOMEDIR"

# Update app.toml with the correct ethereum chain config stuff.
CHAIN_ID=$(jq '.config.chainId' $ETH_GENESIS_JSON)
HOMESTEAD_BLOCK=$(jq '.config.homesteadBlock' $ETH_GENESIS_JSON)
DAO_FORK_BLOCK=$(jq '.config.daoForkBlock' $ETH_GENESIS_JSON)
DAO_FORK_SUPPORT=$(jq '.config.daoForkSupport' $ETH_GENESIS_JSON)
EIP150_BLOCK=$(jq '.config.eip150Block' $ETH_GENESIS_JSON)
EIP155_BLOCK=$(jq '.config.eip155Block' $ETH_GENESIS_JSON)
EIP158_BLOCK=$(jq '.config.eip158Block' $ETH_GENESIS_JSON)
BYZANTIUM_BLOCK=$(jq '.config.byzantiumBlock' $ETH_GENESIS_JSON)
CONSTANTINOPLE_BLOCK=$(jq '.config.constantinopleBlock' $ETH_GENESIS_JSON)
PETERSBURG_BLOCK=$(jq '.config.petersburgBlock' $ETH_GENESIS_JSON)
ISTANBUL_BLOCK=$(jq '.config.istanbulBlock' $ETH_GENESIS_JSON)
BERLIN_BLOCK=$(jq '.config.berlinBlock' $ETH_GENESIS_JSON)
LONDON_BLOCK=$(jq '.config.londonBlock' $ETH_GENESIS_JSON)
MUIR_GLACIER_BLOCK=$(jq '.config.muirGlacierBlock' $ETH_GENESIS_JSON)
ARROW_GLACIER_BLOCK=$(jq '.config.arrowGlacierBlock' $ETH_GENESIS_JSON)
GRAY_GLACIER_BLOCK=$(jq '.config.grayGlacierBlock' $ETH_GENESIS_JSON)
MERGE_NETSPLIT_BLOCK=$(jq '.config.mergeNetsplitBlock' $ETH_GENESIS_JSON)
SHANGHAI_TIME=$(jq '.config.shanghaiTime' $ETH_GENESIS_JSON)
TERMINAL_TOTAL_DIFFICULTY=$(jq '.config.terminalTotalDifficulty' $ETH_GENESIS_JSON)
TERMINAL_TOTAL_DIFFICULTY_PASSED=$(jq '.config.terminalTotalDifficultyPassed' $ETH_GENESIS_JSON)

# Update values in TOML using sed command for Linux with -i and empty string argument
sed -i "s/chain-id = .*/chain-id = \"$CHAIN_ID\"/" $APP_TOML
sed -i "s/homestead-block = .*/homestead-block = \"$HOMESTEAD_BLOCK\"/" $APP_TOML
sed -i "s/dao-fork-block = .*/dao-fork-block = $DAO_FORK_BLOCK/" $APP_TOML
sed -i "s/dao-fork-support = .*/dao-fork-support = $DAO_FORK_SUPPORT/" $APP_TOML
sed -i "s/eip150-block = .*/eip150-block = \"$EIP150_BLOCK\"/" $APP_TOML
sed -i "s/eip155-block = .*/eip155-block = \"$EIP155_BLOCK\"/" $APP_TOML
sed -i "s/eip158-block = .*/eip158-block = \"$EIP158_BLOCK\"/" $APP_TOML
sed -i "s/byzantium-block = .*/byzantium-block = \"$BYZANTIUM_BLOCK\"/" $APP_TOML
sed -i "s/constantinople-block = .*/constantinople-block = \"$CONSTANTINOPLE_BLOCK\"/" $APP_TOML
sed -i "s/petersburg-block = .*/petersburg-block = \"$PETERSBURG_BLOCK\"/" $APP_TOML
sed -i "s/istanbul-block = .*/istanbul-block = \"$ISTANBUL_BLOCK\"/" $APP_TOML
sed -i "s/berlin-block = .*/berlin-block = \"$BERLIN_BLOCK\"/" $APP_TOML
sed -i "s/london-block = .*/london-block = \"$LONDON_BLOCK\"/" $APP_TOML
sed -i "s/muir-glacier-block = .*/muir-glacier-block = \"$MUIR_GLACIER_BLOCK\"/" $APP_TOML
sed -i "s/arrow-glacier-block = .*/arrow-glacier-block = \"$ARROW_GLACIER_BLOCK\"/" $APP_TOML
sed -i "s/gray-glacier-block = .*/gray-glacier-block = \"$GRAY_GLACIER_BLOCK\"/" $APP_TOML
sed -i "s/merge-netsplit-block = .*/merge-netsplit-block = \"$MERGE_NETSPLIT_BLOCK\"/" $APP_TOML
sed -i "s/shanghai-time = .*/shanghai-time = \"$SHANGHAI_TIME\"/" $APP_TOML
sed -i "s/terminal-total-difficulty = .*/terminal-total-difficulty = \"$TERMINAL_TOTAL_DIFFICULTY\"/" $APP_TOML
sed -i "s/terminal-total-difficulty-passed = .*/terminal-total-difficulty-passed = $TERMINAL_TOTAL_DIFFICULTY_PASSED/" $APP_TOML


# Adjust timeouts for CometBFT: 
# TODO: these values are sensitive due to a race condition in the json-rpc ports opening.
# If the JSON-RPC opens before the first block is committed, hive tests will start failing.
# This needs to be fixed before mainnet as its ghetto af.
sed -i 's/timeout_propose = "3s"/timeout_propose = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_prevote = "1s"/timeout_prevote = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_precommit = "1s"/timeout_precommit = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/g' $CONFIG_TOML
sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "2s"/g' $CONFIG_TOML

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)m
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR"