#!/bin/bash

RELATIVE_PATH=scripts/
EXEC=compatriot-init.sh
SEND_TXS=spam-tx.sh
REQUESTS=rpc-requests.sh
NODE_PID=0

# Function to start the binary executable in a separate process
start_node() {
    ./scripts/compatriot-init.sh &
    NODE_PID=$!
    sleep 3
}

# Function to restart the node using the built polard binary
restart_node() {
    polard start --api.enable --home ./bin/polard &
    NODE_PID=$!
    sleep 3
}

# Function to stop the node process
stop_node() {
    kill -9 $NODE_PID
    wait $NODE_PID 2>/dev/null
}

# Function to send transactions defined by
# the tx spam script
send_transactions() {
    # use tx spam script
    ./scripts/spam-tx.sh
}

# Function to send RPC queries defined by
# the rpc requests script
send_rpc_requests() {
    # use rpc spam script
    ./scripts/rpc-requests.sh
}

# start_node

send_transactions

send_rpc_requests

# restart_node

# stop_node

# send_rpc_requests

# stop_node