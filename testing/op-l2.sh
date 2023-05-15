#!/bin/bash

SEQ_KEY="e2d186bc65327b8840f0032434bdd585a7cdf915d28eba7b9699725cf0bda197"

L1_RPC="http://localhost:8545"
RPC_KIND="any"
cd ~/op-stack-deployment/optimism/op-node
L1_RPC="http://localhost:8545"

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
