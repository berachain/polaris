#!/bin/bash
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# See the file LICENSE for licensing terms.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


KEY="mykey"
CHAINID="berachain_420-1"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
TRACE="--trace"

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.berad*

./bin/berad config keyring-backend $KEYRING
./bin/berad config chain-id $CHAINID

# if $KEY exists it should be deleted
./bin/berad keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

##############################################
# Setup 5 User Accounts
##############################################
echo 12345678 | ./bin/berad keys unsafe-import-eth-key --keyring-backend $KEYRING account0 0xe521154ebe9733c29baa4f6f232cb6e7a8928b3bd85e14e95dff9fa8ca8f72b0

echo 12345678 | ./bin/berad keys unsafe-import-eth-key --keyring-backend $KEYRING account1 0x577b84e5765243ce57cece5893c993297ca78f255c2b68e3d50c9d8a2213c821

echo 12345678 | ./bin/berad keys unsafe-import-eth-key --keyring-backend $KEYRING account2 0x315545448acb3083e144687a0ac7d515233d4464bb8eb00cd3feeda6e7a285c6

echo 12345678 | ./bin/berad keys unsafe-import-eth-key --keyring-backend $KEYRING account3 0xb189234091a05a861de7cabfb91ea8aa71bdfd35a4084df9cbfe76d50e95383b

echo 12345678 | ./bin/berad keys unsafe-import-eth-key --keyring-backend $KEYRING account4 0xa5a61114976c416b41788b985007459211f33e9d68612af12768e609b923a6e2

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
./bin/berad init $MONIKER --chain-id $CHAINID


##############################################
# Configuration Module Params
##############################################

# Staking Params
cat $HOME/.berad/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="abgt"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="672h"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Bank
cat $HOME/.berad/config/genesis.json | jq '.app_state["bank"]["send_enabled"][0]["denom"]="abgt"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["bank"]["send_enabled"][0]["enabled"]=false' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

cat $HOME/.berad/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="abgt"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Bera Inflation Control
cat $HOME/.berad/config/genesis.json | jq '.app_state["bic"]["params"]["blocks_per_year"]="10519200"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["bic"]["params"]["community_tax"]="0.15"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["bic"]["minter"]["inflation"]="0.06"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Epochs
cat $HOME/.berad/config/genesis.json | jq '.app_state["epochs"]["epochs"][0]["identifier"]="berachain_epoch_identifier"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json
cat $HOME/.berad/config/genesis.json | jq '.app_state["epochs"]["epochs"][0]["duration"]="672h"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Evm
cat $HOME/.berad/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="abera"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Feemarket
cat $HOME/.berad/config/genesis.json | jq '.app_state["feemarket"]["params"]["base_fee"]="225000000000"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json

# Set gas limit in genesis
cat $HOME/.berad/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.berad/config/tmp_genesis.json && mv $HOME/.berad/config/tmp_genesis.json $HOME/.berad/config/genesis.json


##############################################
# Setup Remaining Genesis Info
##############################################

# Allocate genesis accounts (cosmos formatted addresses)
./bin/berad genesis add-genesis-account $KEY 1000000000000000000000000abgt,100000000000abera --keyring-backend $KEYRING
./bin/berad genesis add-genesis-account account0 250000000abgt,100000000000000000000000abera --keyring-backend $KEYRING
./bin/berad genesis add-genesis-account account1 250000000abgt,100000000000000000000000abera --keyring-backend $KEYRING
./bin/berad genesis add-genesis-account account2 250000000abgt,100000000000000000000000abera --keyring-backend $KEYRING
./bin/berad genesis add-genesis-account account3 250000000abgt,100000000000000000000000abera --keyring-backend $KEYRING
./bin/berad add-genesis-account account4 100000000000000000000000abgt,100000000000000000000000abera --keyring-backend $KEYRING


# Sign genesis transaction
./bin/berad genesis gentx $KEY 1000000000000000000000000abgt --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
./bin/berad genesis collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
./bin/berad genesis validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Copy our custom toml validator config.
cp build/networks/devnet/app.toml $HOME/.berad/config/app.toml
cp build/networks/devnet/config.toml $HOME/.berad/config/config.toml

##############################################
# Start Chain
##############################################

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
./bin/berad start --pruning=nothing --evm.tracer=json $TRACE --log_level $LOGLEVEL --minimum-gas-prices=21000abera --json-rpc.api eth,txpool,personal,net,debug,web3,miner --api.enable --api.enabled-unsafe-cors --chain-id $CHAINID