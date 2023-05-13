#!/bin/bash

SEQ_PRIVATE_KEY="65eeadfde237124aa929afe60076f7c612583d254b925e9ddaeee566acf1223a"

cd ~/op-stack-deployment/optimism/op-node
L1_RPC="http://localhost:8545"

go run cmd/main.go genesis l2 \
    --deploy-config ../packages/contracts-bedrock/deploy-config/deployer.json \
    --deployment-dir ../packages/contracts-bedrock/deployments/deployer/ \
    --outfile.l2 genesis.json \
    --outfile.rollup rollup.json \
    --l1-rpc $L1_RPC

openssl rand -hex 32 > jwt.txt

# cp genesis.json ~/op-stack-deployment/op-geth
cp jwt.txt ~/op-stack-deployment/op-geth


cd ~/op-stack-deployment/op-geth
mkdir datadir
echo "pwd" > datadir/password

echo $SEQ_PRIVATE_KEY > datadir/block-signer-key

./build/bin/geth account import --datadir=datadir --password=datadir/password datadir/block-signer-key

build/bin/geth init --datadir=datadir genesis.json
