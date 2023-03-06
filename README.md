<a href="https://berachain.com">
  <img src="https://wallpaperaccess.com/full/1736439.jpg" width="100%">
</a>
<h1 class="center">
  Polaris
</h1>

<div>
  <a href="https://codecov.io/gh/berachain/stargazer" > 
    <img src="https://codecov.io/gh/berachain/stargazer/branch/main/graph/badge.svg?token=0DYQAKBGVW"/> 
  </a>
  <a href="https://pkg.go.dev/pkg.berachain.dev/stargazer">
    <img src="https://pkg.go.dev/badge/pkg.berachain.dev/stargazer.svg" alt="Go Reference">
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

## Installation of a Cosmos-SDK Blockchain

### From Binary

The easiest way to install a Cosmos-SDK Blockchain running Polaris is to download a pre-built binary. You can find the latest binaries on the [releases](https://github.com/stargazer/releases) page.

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

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

**Step 2: Get Polaris source code**

Clone the `stargazer` repo from the [official repo](https://github.com/berachain/stargazer/) and check
out the `main` branch for the latest stable release.
Build the binary.

```bash
cd $HOME
git clone https://github.com/berachain/stargazer
cd stargazer
git checkout main
go run build/setup.go
```

**Step 3: Build the Node Software**

Run the following command to install `stargazerd` to your `GOPATH` and build the node. `stargazerd` is the node daemon and CLI for interacting with a stargazer node.

```bash
mage install
```

**Step 4: Verify your installation**

Verify your installation with the following command:

```bash
stargazerd version --long
```

A successful installation will return the following:

```bash
name: berachain
server_name: stargazerd
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
