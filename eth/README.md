# Polaris Ethereum

Welcome to Polaris Ethereum, a modular framework for injecting a Go-Ethereum (geth) EVM into any 
underlying consensus layer. This folder's directory structure closely resembles that of geth, as it
is meant to be a thin wrapper around the existing geth codebase. The following architecture diagram
shows how Polaris Ethereum integrates into the application level of a host chain.

![polaris_architecture.png](../docs/web/public/polaris_architecture.png)

## api

`api` includes the public Chain API that Polaris Ethereum exports.
 
## core

`core` includes the Polaris Core logic that runs the EVM: process blocks, transactions, and state
transitions. This encapsulates **State Processor** and **Embedded Host Chain** in the architecture
diagram.

## rpc

`rpc` includes rpc service that can be injected into the host chain's JSON-RPC server. This 
encapsulates **RPC Backend** in the architecture diagram. 

## [provider.go](https://github.com/berachain/polaris/blob/main/eth/provider.go) 

The `PolarisProvider` can be exported and used by the host chain.
