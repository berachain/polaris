<a href="https://berachain.com">
  <img src="https://wallpaperaccess.com/full/1736439.jpg" width="100%">
</a>
<h1 class="center">
  Stargazer
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

The easiest way to install a Cosmos-SDK Blockchain running Stargazer is to download a pre-built binary. You can find the latest binaries on the [releases](https://github.com/stargazer/releases) page.

### From Source

**Step 1: Install Golang**

Go v1.20+ or higher is required for Stargazer

1. Install [Go 1.20+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   sudo apt-get install golang -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   go run build/setup.go
   ```

   For Mac:

   ```sh
   brew install golang
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin
   go run build/setup.go
   ```

2. Confirm your Go installation by checking the version:

   ```sh
   go version
   ```


**Step 2: Get Stargazer source code**

Clone the `stargazer` repo from the [official repo](https://github.com/berachain/stargazer/) and check out the `main` branch for the latest stable release.

```bash
git clone https://github.com/berachain/stargazer
cd stargazer
git checkout main
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