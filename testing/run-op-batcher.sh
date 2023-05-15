#!/bin/bash

cd ~/op-stack-deployment/optimism/op-batcher

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

./bin/op-batcher \
    --l2-eth-rpc=http://localhost:8545 \
    --rollup-rpc=http://localhost:8547 \
    --poll-interval=1s \
    --sub-safety-margin=6 \
    --num-confirmations=1 \
    --safe-abort-nonce-too-low-count=3 \
    --resubmission-timeout=30s \
    --rpc.addr=0.0.0.0 \
    --rpc.port=8548 \
    --rpc.enable-admin \
    --max-channel-duration=1 \
    --l1-eth-rpc=$L1_RPC \
    --private-key=$BATCHER_KEY