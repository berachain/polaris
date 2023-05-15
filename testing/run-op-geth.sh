#!/bin/bash

cd ~/op-stack-deployment/op-geth

SEQ_KEY="e2d186bc65327b8840f0032434bdd585a7cdf915d28eba7b9699725cf0bda197"
L1_RPC="http://localhost:8545"
RPC_KIND="any"

export SEQ_ADDR="0xedd88fff6ed74050f93685ea9ef3a79d92fa850e"
export SEQ_KEY="e2d186bc65327b8840f0032434bdd585a7cdf915d28eba7b9699725cf0bda197"
export BATCHER_KEY="26b1aa5dc21c47ddca7f2144c324971428f00add9d371fa0671e073975fd00ce"
export PROPOSER_KEY="db25fa341cd14ba2f2e96c7b62ec6813068abec70f2b6ffaa1bbfeb89efabaf2"
export RPC_KIND="any"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0"

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