#!/bin/bash

NODE_PID=0

# Function to start the binary executable in a separate process
start_node() {
    ./cosmos/init.sh &
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
    kill $NODE_PID
    wait $NODE_PID 2>/dev/null
}

# Function to perform transactions
perform_transactions() {
    # use tx spam script
}

# Function to send RPC queries
send_rpc_queries() {
    # use rpc spam script
}

# Start the node
start_node

# Perform transactions
perform_transactions

# Send RPC queries
send_rpc_queries

# Stop the node
stop_node

# Start the node again
restart_node

# Retry the RPC queries
send_rpc_queries

# Stop the node
stop_node
