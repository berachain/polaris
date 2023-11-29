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
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the ./build/bin/polard instance
HOMEDIR="./.tmp/polard"
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
make build

overwrite="N"
# fi
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else	
overwrite="Y"
fi

# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# # Remove the previous folder
	rm -rf "$HOMEDIR"

	# # Set moniker and chain-id (Moniker can be anything, chain-id must be an integer)
	./build/bin/polard init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	# Set client config
	./build/bin/polard config set client keyring-backend $KEYRING --home "$HOMEDIR"
	./build/bin/polard config set client chain-id "$CHAINID" --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		./build/bin/polard keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
	done

	# Change parameter token denominations to abera
	jq '.app_state["staking"]["params"]["bond_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["mint_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		./build/bin/polard genesis add-genesis-account $KEY 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
	done

	# Test Account
	# absurd surge gather author blanket acquire proof struggle runway attract cereal quiz tattoo shed almost sudden survey boring film memory picnic favorite verb tank
	# 0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306
	./build/bin/polard genesis add-genesis-account cosmos1yrene6g2zwjttemf0c65fscg8w8c55w58yh8rl 69000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"

	# Sign genesis transaction
	./build/bin/polard genesis gentx ${KEYS[0]} 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
	## In case you want to create multiple validators at genesis
	## 1. Back to `./build/bin/polard keys add` step, init more keys
	## 2. Back to `./build/bin/polard add-genesis-account` step, add balance for those
	## 3. Clone this ~/../build/bin/polard home directory into some others, let's say `~/.cloned./build/bin/polard`
	## 4. Run `gentx` in each of those folders
	## 5. Copy the `gentx-*` folders under `~/.cloned./build/bin/polard/config/gentx/` folders into the original `~/../build/bin/polard/config/gentx`

	# Collect genesis tx
	./build/bin/polard genesis collect-gentxs --home "$HOMEDIR"

	ADDRESS=$(jq -r '.address' $HOMEDIR/config/priv_validator_key.json)
	PUB_KEY=$(jq -r '.pub_key' $HOMEDIR/config/priv_validator_key.json)
	jq --argjson pubKey "$PUB_KEY" '. + {"validators": [{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1000", "name": "Rollkit Sequencer"}]}' "$GENESIS" > temp.json && mv temp.json "$GENESIS"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	./build/bin/polard genesis validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# copy centralized sequencer address into genesis.json
# Note: validator and sequencer are used interchangeably here
ADDRESS=$(jq -r '.address' .tmp/polard/config/priv_validator_key.json)
PUB_KEY=$(jq -r '.pub_key' .tmp/polard/config/priv_validator_key.json)
jq --argjson pubKey "$PUB_KEY" '. + {"validators": [{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1000", "name": "Rollkit Sequencer"}]}' .tmp/polard/config/genesis.json > temp.json && mv temp.json .tmp/polard/config/genesis.json

# set the auth token for DA bridge node
AUTH_TOKEN=$(docker exec $(docker ps -q)  celestia bridge --node.store /home/celestia/bridge/ auth admin)

# set the data availability layer's block height from local-celestia-devnet
DA_BLOCK_HEIGHT=$(docker exec $(docker ps -q) celestia header local-head --token $AUTH_TOKEN | jq '.result.header.height' -r)
echo $DA_BLOCK_HEIGHT

# set a random namespace to post data to DA
NAMESPACE_ID=$(openssl rand -hex 10)
echo $NAMESPACE_ID

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
./build/bin/polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR" --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26658","timeout":60000000000,"gas_limit":6000000,"fee":600000,"auth_token":"'$AUTH_TOKEN'"}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT
