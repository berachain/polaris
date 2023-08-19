<h1 align="center"> Berachain x Celestia x OP Stack ❄️🔭 </h1>

## Directory Structure

> Polaris utilizes [go workspaces](https://go.dev/doc/tutorial/workspaces) to break up the repository into logical sections, helping to reduce cognitive overhead.

<pre>
🔭 Polaris 🔭
├── <a href="./contracts">contracts</a>: Contracts and bindings for Polaris (and hosts).
├── <a href="./docs">docs</a>: Documentation for Polaris.
├── <a href="./cosmos">cosmos</a>: Polaris integrated into a Cosmos-SDK based chain.
├── <a href="./e2e">e2e</a>: End-to-end testing utilities.
├── <a href="./eth">eth</a>: The Core of the Polaris Ethereum Framework.
├── <a href="./lib">lib</a>: A collection of libraries used throughout the repo.
├── <a href="./magefiles">magefiles</a>: Build scripts and utils.
</pre>

## Build & Test

[Golang 1.20+](https://go.dev/doc/install) and [Foundry](https://book.getfoundry.sh/getting-started/installation) are required for Polaris.

1. Install [Go 1.20+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   cd $HOME
   sudo apt-get install golang jq -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   cd $HOME
   brew install go jq
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Install and Set Foundry:

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   foundryup
   ```

3. Clone, Setup and Test:

   ```sh
   cd $HOME
   git clone https://github.com/kobakaku/polaris
   git checkout rollkit-v0.50.0-beta.0
   go run magefiles/setup/setup.go
   mage cosmos:test
   ```

4. Start a local development network:

   ```sh
   mage start
   ```

5. Start Celestia Local Devnet:

   ```sh
   docker run --platform linux/amd64 -p 26658:26658 -p 26659:26659 ghcr.io/rollkit/local-celestia-devnet:v0.11.0-rc8
   ```

6. Put Auth Token into cosmos/init.sh

This auth key is required to authorize rollkit to post to the DA.

![sleep](assets/step2.png)

Place it in `cosmos/init.sh`

![sleep](assets/step2.1.png)

7. The following private key has funds on the Polaris Chain

```bash
Address: 0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4
PrivateKey: 0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306
```

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

The network will have an Ethereum JSON-RPC server running at `http://localhost:8545` and a Tendermint RPC server running at `http://localhost:26657`.
