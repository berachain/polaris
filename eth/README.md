# Polaris Ethereum

Welcome to Polaris Ethereum, a modular framework for injecting a Go-Ethereum (geth) EVM into any 
underlying consensus layer. This folder's directory structure closely resembles that of geth, as it
is meant to be a thin wrapper around the existing geth codebase. 

[insert image here]

## api

`api` includes the Chain API that Polaris Ethereum exports.
 
## core

`core` includes the Polaris Core logic that runs the EVM: process blocks, transactions, and state
transitions.

## rpc

`rpc` includes rpc service that can be injected into the host chain's JSON-RPC server.