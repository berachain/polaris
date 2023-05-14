#!/bin/bash

cd ~/op-stack-deployment/optimism/op-node

SEQ_KEY="d75524d55b28a7babcf40a19a09a813e58c73d0b001edf2df8fea4ce78c86d25"
L1_RPC="http://localhost:8545"
RPC_KIND="any"

export SEQ_ADDR="0x6ccbe65895459635bc4fec22a53f8f935d6e31f8"
export SEQ_KEY="f533a590e17bec876ba042096c5e789a2040824e6a1597aa8e50e7c45f1e188e"
export BATCHER_KEY="3ed8a4f7fb30ca082ba82f42ba232ea57e1cdff6b28d03a95fb48055816ae8fe"
export PROPOSER_KEY="d3d836dd0328fa6367582e362a2ccc272144dde621dde15c66619f4c67459a23"
export RPC_KIND="any"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0x12e4A33ff887D5626f2992315c900D6EB8818169"

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