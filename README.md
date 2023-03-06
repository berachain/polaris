<a href="https://berachain.com">
  <img src="./docs/web/public/header.png" width="100%">
</a>
<h1 class="center">
</h1>

<div>
  <a href="https://codecov.io/gh/berachain/polaris" > 
    <img src="https://codecov.io/gh/berachain/polaris/branch/main/graph/badge.svg?token=0DYQAKBGVW"/> 
  </a>
  <a href="https://pkg.go.dev/pkg.berachain.dev/polaris">
    <img src="https://pkg.go.dev/badge/pkg.berachain.dev/polaris.svg" alt="Go Reference">
  </a>
  <a href="https://magefile.org"> 
    <img alt="Built with Mage" src="https://magefile.org/badge.svg" />
  </a>
  <a href="https://discord.gg/berachain">
    <img alt="Discord" src="https://img.shields.io/discord/924442927399313448.svg?label=discord&color=7289da" />
  </a>
  <a href="https://twitter.com/berachain">
    <img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/berachain">
  </a>
</div>
&nbsp;

## Polaris ❄️🔭

Polaris introduces the new standard of intergrating EVM into your blockchain project. With improvements to speed, security, reliability, and an extended set of features, Polaris will be able to support the next generation of decentralized applications while offering a compelling alternative to existing implementations.

Polaris VM is a blockchain framework built on top of the Cosmos SDK that offers a full-featured EVM with full interoperability to the Cosmos ecosystem. It achieves this through the use of various Stateful Precompiles built into the chain that act as gateways to the greater Cosmos framework. This allows EVM users to perform Cosmos native operations such as voting on governance, delegating to validators, and even communicating with other chains through IBC. This design allows us to maintain the native EVM user experience without sacrifices, providing true interoperability between the Cosmos ecosystem and EVM.

# 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

## Installation

### From Binary

The easiest way to install a Cosmos-SDK Blockchain running Polaris is to download a pre-built binary. You can find the latest binaries on the [releases](https://github.com/polaris/releases) page.

### From Source

**Step 1: Install Golang & Foundry**

Go v1.20+ or higher is required for Polaris

1. Install [Go 1.20+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   cd $HOME
   sudo apt-get install golang -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   cd $HOME
   brew install go
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Confirm your Go installation by checking the version:

   ```sh
   go version
   ```

[Foundry](https://book.getfoundry.sh/getting-started/installation) is required for Polaris

3. Install Foundry:
polaris
   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

**Step 2: Get Polaris source code**

Clone the `polaris` repo from the [official repo](https://github.com/berachain/polaris/) and check
out the `main` branch for the latest stable release.
Build the binary.

```bash
cd $HOME
git clone https://github.com/berachain/polaris
cd polaris
git checkout main
go run build/setup.go
```

**Step 3: Build the Node Software**

Run the following command to install `polard` to your `GOPATH` and build the node. `polard` is the node daemon and CLI for interacting with a polaris node.

```bash
mage install
```

**Step 4: Verify your installation**

Verify your installation with the following command:

```bash
polard version --long
```

A successful installation will return the following:

```bash
name: berachain
server_name: polard
version: <x.x.x>
commit: <Commit hash>
build_tags: netgo,ledger
go: go version go1.20.4 darwin/amd64
```

## Running a Local Network

After ensuring dependecies are installed correctly, run the following command to start a local development network.
```bash
mage start
```
