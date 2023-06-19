#!/bin/bash

LOGLEVEL="info"
HOMEDIR="/.polaris"
TRACE=""

# Start the node
./scripts/compatriot-init.sh &
NODE_PID=$!
echo "start_node()" $NODE_PID

sleep 10 # wait for the node to start

# Perform transactions
./scripts/spam-tx.sh

# Send RPC requests
./scripts/rpc-requests.sh cached

# Stop the node
pkill -f "polard start"
kill -9 $NODE_PID

echo "node stopped"

# Start the node again
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR" &
NODE_PID=$!

sleep 10 # wait for the node to start

# Send the RPC requests again
./scripts/rpc-requests.sh non-cached

# Stop the node
kill -9 $NODE_PID

# Diff errors if there are differences
diff -r cached non-cached
