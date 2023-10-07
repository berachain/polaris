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

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="polaris-2061"
MONIKER="localtestnet"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the ./bin/polard instance
HOMEDIR="/"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG_TOML=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
# make build

# Remove the previous folder

 # Set moniker and chain-id (Moniker can be anything, chain-id must be an integer)
polard init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"
# Set client config
polard config set client keyring-backend $KEYRING --home "$HOMEDIR"
polard config set client chain-id "$CHAINID" --home "$HOMEDIR"

# If keys exist they should be deleted
for KEY in "${KEYS[@]}"; do
    polard keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
done

ETH_GENESIS_SOURCE=$(cat genesis.json)
JSON_FILE="genesis.json"
temp_genesis=$(cat $GENESIS)
echo "eth_gen dump: "
echo $ETH_GENESIS_SOURCE


# # # Change eth_genesis in config/genesis.json
# # # TODO FIX TO SETUP APP.TOML STUFF
updated_genesis=$(echo "$temp_genesis" | jq --argjson eth_gen "$ETH_GENESIS_SOURCE" '.app_state["evm"] = $eth_gen')
echo "$updated_genesis" > "$GENESIS"

# echo "UPDATED FILE"

# Change parameter token denominations to abera
jq '.app_state["staking"]["params"]["bond_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["mint"]["params"]["mint_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

# Dump genesis
# echo "genesis state:"
# cat $GENESIS

# Allocate genesis accounts (cosmos formatted addresses)
for KEY in "${KEYS[@]}"; do
    polard genesis add-genesis-account $KEY 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
done

# Sign genesis transaction
polard genesis gentx ${KEYS[0]} 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
## In case you want to create multiple validators at genesis
## 1. Back to `./bin/polard keys add` step, init more keys
## 2. Back to `./bin/polard add-genesis-account` step, add balance for those
## 3. Clone this ~/../bin/polard home directory into some others, let's say `~/.cloned./bin/polard`
## 4. Run `gentx` in each of those folders
## 5. Copy the `gentx-*` folders under `~/.cloned./bin/polard/config/gentx/` folders into the original `~/../bin/polard/config/gentx`

# Collect genesis tx
polard genesis collect-gentxs --home "$HOMEDIR"

# Run this to ensure everything worked and that the genesis file is setup correctly
polard genesis validate-genesis --home "$HOMEDIR"

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
fi

# Read values from JSON
CHAIN_ID=$(jq '.config.chainId' $JSON_FILE)
HOMESTEAD_BLOCK=$(jq '.config.homesteadBlock' $JSON_FILE)
DAO_FORK_BLOCK=$(jq '.config.daoForkBlock' $JSON_FILE)
DAO_FORK_SUPPORT=$(jq '.config.daoForkSupport' $JSON_FILE)
EIP150_BLOCK=$(jq '.config.eip150Block' $JSON_FILE)
EIP155_BLOCK=$(jq '.config.eip155Block' $JSON_FILE)
EIP158_BLOCK=$(jq '.config.eip158Block' $JSON_FILE)
BYZANTIUM_BLOCK=$(jq '.config.byzantiumBlock' $JSON_FILE)
CONSTANTINOPLE_BLOCK=$(jq '.config.constantinopleBlock' $JSON_FILE)
PETERSBURG_BLOCK=$(jq '.config.petersburgBlock' $JSON_FILE)
ISTANBUL_BLOCK=$(jq '.config.istanbulBlock' $JSON_FILE)
BERLIN_BLOCK=$(jq '.config.berlinBlock' $JSON_FILE)
LONDON_BLOCK=$(jq '.config.londonBlock' $JSON_FILE)
MUIR_GLACIER_BLOCK=$(jq '.config.muirGlacierBlock' $JSON_FILE)
ARROW_GLACIER_BLOCK=$(jq '.config.arrowGlacierBlock' $JSON_FILE)
GRAY_GLACIER_BLOCK=$(jq '.config.grayGlacierBlock' $JSON_FILE)
MERGE_NETSPLIT_BLOCK=$(jq '.config.mergeNetsplitBlock' $JSON_FILE)
SHANGHAI_TIME=$(jq '.config.shanghaiTime' $JSON_FILE)
TERMINAL_TOTAL_DIFFICULTY=$(jq '.config.terminalTotalDifficulty' $JSON_FILE)
TERMINAL_TOTAL_DIFFICULTY_PASSED=$(jq '.config.terminalTotalDifficultyPassed' $JSON_FILE)

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

# sed -i 's/"null"/null/g' $APP_TOML
# sed -i 's/"null"/null/g' $APP_TOML
# cat "POSTSED"
# cat $APP_TOML
cat $APP_TOML
echo "BINGBONG"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)m
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR"