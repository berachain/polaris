#!/bin/bash

export SEQ_ADDR="0x7530B10263a0861D29Fad72f45a3713E59Dc1cD2"
export SEQ_KEY="6baf56caba428fccf3e25826ec9e577be665fecca0d98547a6d50ae162ffad87"
export BATCHER_KEY="379df969ecfffac36c393117fedad6cdee462b213f979ef4cf70942e431c3506"
export PROPOSER_KEY="e583e1c48a2c31cce469a9bf54f3c03efeb1938a3a18e9ce773301a136ecd11c"
export RPC_KIND="basic"
export L1_RPC="http://localhost:8545"
cd ~/op-stack-deployment/optimism/op-node

go run cmd/main.go genesis l2 \
    --deploy-config ../packages/contracts-bedrock/deploy-config/getting-started.json \
    --deployment-dir ../packages/contracts-bedrock/deployments/getting-started/ \
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
