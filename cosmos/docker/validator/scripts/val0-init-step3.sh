KEY1="val0"
KEYRING="test"
HOMEDIR="/root/.berad"
VAL_JSON="$HOMEDIR/config/validator.json"

# Generating a JSON string (https://stackoverflow.com/a/48470227)
validator_json_string=$(
  jq --null-input \
    '{
        "pubkey": {"@type":"/cosmos.crypto.ed25519.PubKey","key":"oWg2ISpLF405Jcm2vXV+2v4fnjodh6aafuIdeoW+rUw="},
        "amount": "1000000stake",
        "moniker": "myvalidator",
        "identity": "optional identity signature (ex. UPort or Keybase)",
        "website": "validator's (optional) website",
        "security": "validator's (optional) security contact email",
        "details": "validator's (optional) details",
        "commission-rate": "0.1",
        "commission-max-rate": "0.2",
        "commission-max-change-rate": "0.01",
        "min-self-delegation": "1"
    }'
)

# Creating the JSON file
echo $validator_json_string > $VAL_JSON

berad tx staking create-validator $VAL_JSON --from $KEY1 --home "$HOMEDIR"
