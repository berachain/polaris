apk add bash jq

KEY2="seed1"
CHAINID="brickchain-666"
MONIKER2="seed-1"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
HOMEDIR="/pv/.polard"
TRACE=""
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json


polard init $MONIKER2 -o --chain-id $CHAINID --home "$HOMEDIR"

polard config set client keyring-backend $KEYRING --home "$HOMEDIR"

polard keys add $KEY2 --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
