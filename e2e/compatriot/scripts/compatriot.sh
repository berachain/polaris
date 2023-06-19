#!/bin/bash

# Include the functions script which
# contain all the helper functions
source ./scripts/helper.sh

# Start the node
start_node

# Perform transactions
send_transactions

# Send RPC requests
send_rpc_requests

# Stop the node
stop_node

# Start the node again
restart_node
t
# Retry the RPC requests
send_rpc_requests

# Stop the node
stop_node
