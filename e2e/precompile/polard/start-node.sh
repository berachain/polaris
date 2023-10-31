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
# Set dedicated home directory for the ./build/bin/polard instance
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

# echo "UPDATED FILE"

# Change parameter token denominations to abera
jq '.app_state["staking"]["params"]["bond_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["gov"]["params"]["min_deposit"][0]["amount"]="1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["gov"]["params"]["expedited_min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["gov"]["params"]["expedited_min_deposit"][0]["amount"]="2"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS";
jq '.app_state["mint"]["params"]["mint_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

# Dump genesis
# echo "genesis state:"
# cat $GENESIS

# Allocate genesis accounts (cosmos formatted addresses)
for KEY in "${KEYS[@]}"; do
    polard genesis add-genesis-account $KEY 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
done

# alice cosmos1dgtgps0vxwt90hu6f3cceqypc5k664czp95ank
# bob cosmos1h08vp7xt40nks7d0mlg47duyv54ewdxr0p0f44
# charlie cosmos14nqnr8l8y2se3uu47qtyqehdfccfgwdlmshdpp

# Give alice, bob and charlie some bank tokens.
polard genesis add-genesis-account cosmos1dgtgps0vxwt90hu6f3cceqypc5k664czp95ank 1000000000000000000abera,1000000000000000000asupply,1000000000000000000atoken,12345bAKT,1000000000000000000bATOM,24690bOSMO,1000000000000000000stake  --keyring-backend $KEYRING --home "$HOMEDIR"
polard genesis add-genesis-account cosmos1h08vp7xt40nks7d0mlg47duyv54ewdxr0p0f44 100abera,100atoken,1000000000000000000stake --keyring-backend $KEYRING --home "$HOMEDIR"
polard genesis add-genesis-account cosmos14nqnr8l8y2se3uu47qtyqehdfccfgwdlmshdpp 1000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"

# Give alice, bob and charlie some evm tokens.
jq '.app_state["evm"]["alloc"]["6A1680c1Ec339657df9a4c718C8081C52daD5702"]["balance"]="0x4563918244f40000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["evm"]["alloc"]["bBcec0f8cBAbe76879AfdfD15F3784652B9734C3"]["balance"]="0x4563918244f40000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["evm"]["alloc"]["acc1319Fe722A198F395F0164066ED4E309439Bf"]["balance"]="0x4563918244f40000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

# Sign genesis transaction
polard genesis gentx ${KEYS[0]} 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
## In case you want to create multiple validators at genesis
## 1. Back to `./build/bin/polard keys add` step, init more keys
## 2. Back to `./build/bin/polard add-genesis-account` step, add balance for those
## 3. Clone this ~/../build/bin/polard home directory into some others, let's say `~/.cloned./build/bin/polard`
## 4. Run `gentx` in each of those folders
## 5. Copy the `gentx-*` folders under `~/.cloned./build/bin/polard/config/gentx/` folders into the original `~/../build/bin/polard/config/gentx`

# Collect genesis tx
polard genesis collect-gentxs --home "$HOMEDIR"

# Run this to ensure everything worked and that the genesis file is setup correctly
polard genesis validate-genesis --home "$HOMEDIR"

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
fi

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
