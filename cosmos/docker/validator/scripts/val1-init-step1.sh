KEY2="val1"
CHAINID="brickchain-666"
MONIKER1="val-1"
KEYRING="test"
KEYALGO="eth_secp256k1"
HOMEDIR="/root/.berad"

berad init $MONIKER1 -o --chain-id $CHAINID --home "$HOMEDIR"

berad config set client keyring-backend $KEYRING --home "$HOMEDIR"

berad keys add $KEY1 --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
  