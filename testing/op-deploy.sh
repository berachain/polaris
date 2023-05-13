#!/bin/bash

cd ~/op-stack-deployment/optimism/packages/contracts-bedrock

# Step 1: Generate some Keys
output=$(npx hardhat rekey)

# Parse the output using awk and store the values in variables
mnemonic=$(echo "$output" | awk '/Mnemonic:/ { print $2, $3, $4, $5, $6, $7, $8, $9 }')
admin_address=$(echo "$output" | awk '/Admin:/ { print $2 }')
admin_private_key=$(echo "$output" | awk '/Admin:/ { getline; print $3 }')
proposer_address=$(echo "$output" | awk '/Proposer:/ { print $2 }')
proposer_private_key=$(echo "$output" | awk '/Proposer:/ { getline; print $3 }')
batcher_address=$(echo "$output" | awk '/Batcher:/ { print $2 }')
batcher_private_key=$(echo "$output" | awk '/Batcher:/ { getline; print $3 }')
sequencer_address=$(echo "$output" | awk '/Sequencer:/ { print $2 }')
sequencer_private_key=$(echo "$output" | awk '/Sequencer:/ { getline; print $3 }')

# Print the variables
echo "Mnemonic: $mnemonic"
echo "Admin Address: $admin_address"
echo "Admin Private Key: $admin_private_key"
echo "Proposer Address: $proposer_address"
echo "Proposer Private Key: $proposer_private_key"
echo "Batcher Address: $batcher_address"
echo "Batcher Private Key: $batcher_private_key"
echo "Sequencer Address: $sequencer_address"
echo "Sequencer Private Key: $sequencer_private_key"

# Step 2: Copy .env.example to .env
cp .env.example .env

# Step3: Fill out environment variables in .env file
L1_RPC_URL="http://localhost:8545"  # Replace with your L1 node URL
PRIVATE_KEY_DEPLOYER="$admin_private_key"  # Replace with the private key of the Admin account
CHAIN_ID=69420

# Replace the values in .env file
sed -i.bak "s|^L1_RPC=.*|L1_RPC=$L1_RPC_URL|g" .env && rm .env.bak
sed -i.bak "s|^PRIVATE_KEY_DEPLOYER=.*|PRIVATE_KEY_DEPLOYER=$PRIVATE_KEY_DEPLOYER|g" .env && rm .env.bak

# Add CHAIN_ID to .env file
echo "" >> .env
echo "# Set the L1 ChainID" >> .env
echo "CHAIN_ID=$CHAIN_ID" >> .env

echo "Sending 100 ether to all addresses..."
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $admin_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $proposer_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $batcher_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $sequencer_address --value 100ether


# Step 3: Deploy L1 smart contracts
npx hardhat deploy --network deployer --tags l1

echo "L1 smart contract deployment completed."