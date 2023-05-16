#!/bin/bash

cd ~/op-stack-deployment/op-geth

SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
L1_RPC="http://localhost:8545"
RPC_KIND="any"

export SEQ_ADDR="0xed960280ba229deb3eb8ddf968a1dc5019c378ba"
export SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
export BATCHER_KEY="d30f6dfe88e12a4303038fd00da6715da50e750421d68da3d35d22d0e20d5952"
export PROPOSER_KEY="4df8ff5f78e5bef31c25150aa2dbb91351379a7a3a7e9c41662b5bf362882e0c"
export RPC_KIND="any"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0x3D52168A3408d32a38496FDE80458777846e217B"

./build/bin/geth \
	--datadir ./datadir \
	--http \
	--http.corsdomain="*" \
	--http.vhosts="*" \
	--http.addr=0.0.0.0 \
    --http.port=7545 \
	--http.api=web3,debug,eth,txpool,net,engine \
	--ws \
	--ws.addr=0.0.0.0 \
	--ws.port=7546 \
	--ws.origins="*" \
	--ws.api=debug,eth,txpool,net,engine \
	--syncmode=full \
	--gcmode=archive \
	--nodiscover \
	--maxpeers=0 \
	--networkid=42069 \
	--authrpc.vhosts="*" \
	--authrpc.addr=0.0.0.0 \
	--authrpc.port=8551 \
	--authrpc.jwtsecret=./jwt.txt \
	--rollup.disabletxpoolgossip=true \
	--password=./datadir/password \
	--allow-insecure-unlock \
	--mine \
	--miner.etherbase=$SEQ_ADDR \
	--unlock=$SEQ_ADDR