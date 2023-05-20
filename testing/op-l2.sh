#!/bin/bash

SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
L1_RPC="http://localhost:8545"
RPC_KIND="basic"

export SEQ_ADDR="0xed960280ba229deb3eb8ddf968a1dc5019c378ba"
export SEQ_KEY="73bead3e3e588be3ae77598160d6a5017acdb008973f21166571e400a4b51246"
export BATCHER_KEY="d30f6dfe88e12a4303038fd00da6715da50e750421d68da3d35d22d0e20d5952"
export PROPOSER_KEY="4df8ff5f78e5bef31c25150aa2dbb91351379a7a3a7e9c41662b5bf362882e0c"
export RPC_KIND="basic"
export L1_RPC="http://localhost:8545"
export L2OO_ADDR="0x3D52168A3408d32a38496FDE80458777846e217B"
cd ~/op-stack-deployment/optimism/op-node

go run cmd/main.go genesis l2 \
    --deploy-config ../packages/contracts-bedrock/deploy-config/deployer.json \
    --deployment-dir ../packages/contracts-bedrock/deployments/deployer/ \
    --outfile.l2 genesis.json \
    --outfile.rollup rollup.json \
    --l1-rpc $L1_RPC

openssl rand -hex 32 > jwt.txt

cp genesis.json ~/op-stack-deployment/op-geth
cp jwt.txt ~/op-stack-deployment/op-geth


cd ~/op-stack-deployment/op-geth
mkdir datadir
echo "pwd" > datadir/password

echo $SEQ_KEY > datadir/block-signer-key

./build/bin/geth account import --datadir=datadir --password=datadir/password datadir/block-signer-key

build/bin/geth init --datadir=datadir genesis.json
