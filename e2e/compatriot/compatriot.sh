#!/bin/bash

LOGLEVEL="info"
HOMEDIR="/.polaris"
TRACE=""

declare NODE_PID

# Start the node
./scripts/compatriot-init.sh &
NODE_PID=$!
echo "start_node()" $NODE_PID

sleep 10 # wait for the node to start

# # Perform transactions
# ./scripts/spam-tx.sh

# # Send RPC requests
# ./scripts/rpc-requests.sh

# Stop the node
kill %1

# Start the node again
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001abera --home "$HOMEDIR" &
NODE_PID=$!
echo "restart_node()" $NODE_PID

sleep 10 # wait for the node to start

# # Send the RPC requests again
# ./scripts/rpc-requests.sh

jobs

# Stop the node
kill %2
