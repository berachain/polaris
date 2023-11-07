<h1 align="center"> Polaris Monorepo ❄️🔭 </h1>

*The project is still work in progress, see the [disclaimer below](#-warning-under-construction-).*

<div>
  <a href="https://codecov.io/gh/berachain/polaris" target="_blank">
    <img src="https://codecov.io/gh/berachain/polaris/branch/main/graph/badge.svg?token=5SYYGUS8GW"/>
  </a>
  <a href="https://pkg.go.dev/github.com/berachain/polaris" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/berachain/polaris.svg" alt="Go Reference">
  </a>
  <a href="https://t.me/polaris_devs" target="_blank">
    <img alt="Telegram Chat" src="https://img.shields.io/endpoint?color=neon&logo=telegram&label=chat&url=https%3A%2F%2Ftg.sumanjay.workers.dev%2Fpolaris_devs">
  </a>
  <a href="https://twitter.com/berachain" target="_blank">
    <img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/berachain">
  <a href="https://discord.gg/berachain">
   <img src="https://img.shields.io/discord/984015101017346058?color=%235865F2&label=Discord&logo=discord&logoColor=%23fff" alt="Discord">
  </a>
</div>

## What is Polaris?

Introducing Polaris, the revolutionary framework designed to simplify the integration of an Ethereum Virtual Machine (EVM) into your application. Polaris is built with a clean, easy-to-integrate API that eliminates the need for developers to spend time hacking together their own EVM integration solutions. Our framework is highly modular, allowing you to choose the components that best fit your needs and integrate an EVM environment into virtually any application.

Polaris is built with several core principles in mind:

1. **Modularity**: Each component is developed as a distinct package, complete with thorough testing, documentation, and benchmarking. You can use these components individually or combine them to create innovative EVM integrations.
2. **Configurability**: We want Polaris to be accessible to as many teams and use cases as possible. To support this, our framework is highly configurable, allowing you to tailor it to your specific needs.
3. **Performance**: In today's competitive crypto landscape, performance is key. Polaris is optimized to deliver the highest levels of performance and efficiency.
4. **Contributor Friendliness**: We believe that open collaboration is key to driving innovation in blockchain development. While Polaris is currently licensed under BUSL-1.1, we plan to adjust our licensing to support contributor-based schemes as we approach production readiness.
5. **Memes**: If ur PR doesn't have a meme in it like idk sry bro, gg wp glhf.

## Documentation

If you want to build on top of Polaris, take a look at our [documentation](http://polaris.berachain.dev/).
If you want to help contribute to the framework, check out the [Framework Specs](./specs/).

## Directory Structure

> Polaris utilizes [go workspaces](https://go.dev/doc/tutorial/workspaces) to break up the repository into logical sections, helping to reduce cognitive overhead.

<pre>
🔭 Polaris 🔭
├── <a href="./build">build</a>: Build scripts and developer tooling.
├── <a href="./contracts">contracts</a>: Contracts and bindings for Polaris (and hosts).
├── <a href="./cosmos">cosmos</a>: Polaris integrated into a Cosmos-SDK based chain.
├── <a href="./e2e">e2e</a>: End-to-end testing utilities.
├── <a href="./eth">eth</a>: The Core of the Polaris Ethereum Framework.
├── <a href="./lib">lib</a>: A collection of libraries used throughout the repo.
├── <a href="./proto">proto</a>: Protobuf definitions.
</pre>

## Build & Test

[Golang 1.20+](https://go.dev/doc/install) and [Foundry](https://book.getfoundry.sh/getting-started/installation) are required for Polaris.

1. Install [go 1.21+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

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

2. Install Foundry:

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

3. Clone, Setup and Test:

   ```sh
   cd $HOME
   git clone https://github.com/berachain/polaris
   cd polaris
   git checkout main
   make test-unit
   ```

4. Start a local development network:

   ```sh
   make start
   ```

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

The network will have an Ethereum JSON-RPC server running at `http://localhost:8545` and a Tendermint RPC server running at `http://localhost:26657`.
