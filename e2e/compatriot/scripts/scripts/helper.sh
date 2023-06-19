#!/bin/bash

RELATIVE_PATH=scripts/
EXEC=compatriot-init.sh
SEND_TXS=txs.sh
REQUESTS=rpc-requests.sh
NODE_PID=0

# Function to start the binary executable in a separate process
start_node() {
    run "${$RELATIVE_PATH}${EXEC}" &
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
    run "${$RELATIVE_PATH}${SEND_TXS}"
}

# Function to send RPC queries defined by
# the rpc requests script
send_rpc_requests() {
    # use rpc spam script
    run "${$RELATIVE_PATH}${REQUESTS}"
}
