#!/bin/bash

cd ~/op-stack-deployment/optimism/op-proposer

SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
L1_RPC="http://localhost:8545"
DA_RPC="http://localhost:26659"
RPC_KIND="basic"

export SEQ_ADDR="0xed960280ba229deb3eb8ddf968a1dc5019c378ba"
export SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
export BATCHER_KEY="d30f6dfe88e12a4303038fd00da6715da50e750421d68da3d35d22d0e20d5952"
export PROPOSER_KEY="4df8ff5f78e5bef31c25150aa2dbb91351379a7a3a7e9c41662b5bf362882e0c"
export RPC_KIND="basic"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0x3D52168A3408d32a38496FDE80458777846e217B"

./bin/op-proposer \
    --poll-interval 5s \
    --rpc.port 8560 \
    --rollup-rpc http://localhost:8547 \
    --l2oo-address $L2OO_ADDR \
    --private-key $PROPOSER_KEY \
    --l1-eth-rpc $L1_RPC \
    --da-rpc $DA_RPC