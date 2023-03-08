# Deploying Liquid Staking Contract Using Forge Foundry

**Step 1: Get the validator address and convert to hex address**

Set this value in the `Deploy.s.sol` solidity forge script.

```sh
./bin/polard query staking validators
```

**Step 2: Run the Script To Deploy**

Export the rpc url to your environment.

```sh
export ETH_RPC_URL=http://localhost:1317/eth/rpc
```

Run the script to deploy the contract, can change the private key to one of your choosing, underneath we use the private key in the `init.sh` file.

```sh
forge script src/examples/Deploy.s.sol:Deploy --broadcast  --private-key 0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306 --rpc-url $ETH_RPC_URL --gas-limit 10000000
```
