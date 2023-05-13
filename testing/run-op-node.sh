#!/bin/bash

cd ~/op-stack-deployment/optimism/op-node

SEQ_KEY="65eeadfde237124aa929afe60076f7c612583d254b925e9ddaeee566acf1223a"
L1_RPC="http://localhost:8545"
RPC_KIND="basic"

./bin/op-node \
	--l2=http://localhost:8551 \
	--l2.jwt-secret=./jwt.txt \
	--sequencer.enabled \
	--sequencer.l1-confs=3 \
	--verifier.l1-confs=3 \
	--rollup.config=./rollup.json \
	--rpc.addr=0.0.0.0 \
	--rpc.port=8547 \
	--p2p.disable \
	--rpc.enable-admin \
	--p2p.sequencer.key=$SEQ_KEY \
	--l1=$L1_RPC \
	--l1.rpckind=$RPC_KIND