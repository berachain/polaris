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

CHAINID="polaris-2061"
MONIKER="localtestnet"

LOGLEVEL="info"
# Set dedicated home directory for the ./bin/polard instance
HOMEDIR="/"
# to trace evm 
#TRACE="--trace"
TRACE=""

# used to exit on first error (any non-zero exit code)
set -e

GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json
cp $GENESIS $TMP_GENESIS

# set the current time as genesis time
CURRENT_TIME=$(date -u +"%Y-%m-%dT%H:%M:%S.000000Z")
cat "$GENESIS" | jq --arg curr_time "$CURRENT_TIME" '.genesis_time = $curr_time' > "$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

## In case you want to create multiple validators at genesis
## 1. Back to `./bin/polard keys add` step, init more keys
## 2. Back to `./bin/polard add-genesis-account` step, add balance for those
## 3. Clone this ~/../bin/polard home directory into some others, let's say `~/.cloned./bin/polard`
## 4. Run `gentx` in each of those folders
## 5. Copy the `gentx-*` folders under `~/.cloned./bin/polard/config/gentx/` folders into the original `~/../bin/polard/config/gentx`

# Run this to ensure everything worked and that the genesis file is setup correctly
polard genesis validate-genesis --home "$HOMEDIR"

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)m
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR"
