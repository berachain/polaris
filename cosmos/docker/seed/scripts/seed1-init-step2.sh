KEY2="seed1"
KEYRING="test"
CHAINID="brickchain-666"
HOMEDIR="/pv/.polard"

polard genesis add-genesis-account $KEY2 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"

polard genesis gentx $KEY2 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
