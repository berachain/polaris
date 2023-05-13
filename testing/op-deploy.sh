#!/bin/bash

# Step 3: Fill out environment variables in .env file
L1_RPC_URL="http://localhost:8545"  # Replace with your L1 node URL
L1_CHAINID=69420
L2_CHAINID=69421

# Step 0: Copy the examples to the contracts-bedrock
cp example.deployer.json ~/op-stack-deployment/optimism/packages/contracts-bedrock/deploy-config/deployer.json
cp example.deployer.ts ~/op-stack-deployment/optimism/packages/contracts-bedrock/deploy-config/deployer.ts
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

# Replace the values in .env file
PRIVATE_KEY_DEPLOYER="$admin_private_key"  # Replace with the private key of the Admin account
DISABLE_LIVE_DEPLOYER=false
sed -i.bak "s|^L1_RPC=.*|L1_RPC=$L1_RPC_URL|g" .env && rm .env.bak
sed -i.bak "s|^PRIVATE_KEY_DEPLOYER=.*|PRIVATE_KEY_DEPLOYER=$PRIVATE_KEY_DEPLOYER|g" .env && rm .env.bak
sed -i.bak "s|^DISABLE_LIVE_DEPLOYER=.*|DISABLE_LIVE_DEPLOYER=$DISABLE_LIVE_DEPLOYER|g" .env && rm .env.bak
# Add CHAIN_ID to .env file
echo "" >> .env
echo "# Set the L1 ChainID" >> .env
echo "CHAIN_ID=$L1_CHAINID" >> .env
cat .env

echo "Sending 100 ether to all addresses..."
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $admin_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $proposer_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $batcher_address --value 100ether
cast send --private-key=fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 $sequencer_address --value 100ether

# Update deploy-config/deployer.json with addresses from reky
sed -i.bak "s|\"finalSystemOwner\": \"ADMIN\"|\"finalSystemOwner\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"portalGuardian\": \"ADMIN\"|\"portalGuardian\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"controller\": \"ADMIN\"|\"controller\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"p2pSequencerAddress\": \"SEQUENCER\"|\"p2pSequencerAddress\": \"$sequencer_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"batchInboxAddress\": \"0xff00000000000000000000000000000000042069\"|\"batchInboxAddress\": \"$batcher_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"batchSenderAddress\": \"BATCHER\"|\"batchSenderAddress\": \"$batcher_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l2OutputOracleProposer\": \"PROPOSER\"|\"l2OutputOracleProposer\": \"$proposer_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l2OutputOracleChallenger\": \"ADMIN\"|\"l2OutputOracleChallenger\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"proxyAdminOwner\": \"ADMIN\"|\"proxyAdminOwner\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"baseFeeVaultRecipient\": \"ADMIN\"|\"baseFeeVaultRecipient\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l1FeeVaultRecipient\": \"ADMIN\"|\"l1FeeVaultRecipient\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"sequencerFeeVaultRecipient\": \"ADMIN\"|\"sequencerFeeVaultRecipient\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"governanceTokenOwner\": \"ADMIN\"|\"governanceTokenOwner\": \"$admin_address\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak

# Get L1 Info
output=$(cast block finalized | grep -E "(timestamp|hash|number)")

# Parse the output using awk and store the values in variables
hash=$(echo "$output" | awk '/hash/ { print $2 }')
number=$(echo "$output" | awk '/number/ { print $2 }')
timestamp=$(echo "$output" | awk '/timestamp/ { print $2 }')

# Print the variables
echo "Hash: $hash"
echo "Number: $number"
echo "Timestamp: $timestamp"

# Update deploy-config/deployer.json file with the values
sed -i.bak "s|\"l1StartingBlockTag\": \"BLOCKHASH\"|\"l1StartingBlockTag\": \"$hash\"|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l2OutputOracleStartingTimestamp\": TIMESTAMP|\"l2OutputOracleStartingTimestamp\": $timestamp|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l1ChainID\": L1_CHAINID|\"l1ChainID\": $L1_CHAINID|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak
sed -i.bak "s|\"l2ChainID\": L2_CHAINID|\"l2ChainID\": $L2_CHAINID|g" deploy-config/deployer.json && rm deploy-config/deployer.json.bak

# Print the updated JSON file
cat deploy-config/deployer.json
cat .env
source .env

export DISABLE_LIVE_DEPLOYER=false

# Step 4: Deploy L1 smart contracts
npx hardhat deploy --network deployer --tags l1 --reset
echo "L1 smart contract deployment completed."