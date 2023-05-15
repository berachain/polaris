#!/bin/bash

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
