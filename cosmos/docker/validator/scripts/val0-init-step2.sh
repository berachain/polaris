KEY1="val0"
KEYRING="test"
HOMEDIR="/root/.berad"

berad genesis add-genesis-account $KEY1 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
