#!/bin/bash

cd ~/op-stack-deployment/optimism/op-proposer

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

./bin/op-proposer \
    --poll-interval 5s \
    --rpc.port 8560 \
    --rollup-rpc http://localhost:8547 \
    --l2oo-address $L2OO_ADDR \
    --private-key $PROPOSER_KEY \
    --l1-eth-rpc $L1_RPC