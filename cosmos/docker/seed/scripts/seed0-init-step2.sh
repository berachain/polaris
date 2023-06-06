HOMEDIR="/pv/.polard"

polard genesis collect-gentxs --home "$HOMEDIR"

polard genesis validate-genesis --home "$HOMEDIR"
