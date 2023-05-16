#!/bin/bash

cd ~/op-stack-deployment/optimism/op-batcher

SEQ_KEY="3b6b715c09a45bd370e89d3f03e99b739463b81a65aebd6e24c4b3c037b21f89"
L1_RPC="http://localhost:8545"
RPC_KIND="any"

export SEQ_ADDR="0xcd1c0aaca7fb0c528cf8744bfd4efe1c55694754"
export SEQ_KEY="3b6b715c09a45bd370e89d3f03e99b739463b81a65aebd6e24c4b3c037b21f89"
export BATCHER_KEY="cb985d350b1300767a3a596bfd061de22bbb0fa7e86f2495916315a5d31723a5"
export PROPOSER_KEY="044f9ab10c335a0e537debe3dd5543b17d09528f61ac483a2e584229d582e0f0"
export RPC_KIND="any"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0x8b495AAbAaD833c7433270a79594029912C7e480"

./bin/op-batcher \
    --l2-eth-rpc=http://localhost:7545 \
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