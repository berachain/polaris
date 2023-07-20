# Precompiles

Polaris Ethereum supports the execution of **stateless** and **stateful** precompiled contracts. 

## Stateless Precompiles

Stateless precompiles run completely in the EVM and do not rely on state from the native host chain.
Go-ethereum provides implementations of several stateless precompiles at hardcoded addresses [here](https://github.com/berachain/polaris-geth/blob/stateful-v1.11.4/core/vm/contracts.go). More stateless
precompiles can be added by adhering to the `StatelessImpl`, defined in [interfaces.go](https://github.com/berachain/polaris/blob/main/eth/core/precompile/interfaces.go#L48).

If no custom precompiles are added by the host chain, the [default precompile plugin](https://github.com/berachain/polaris/blob/main/eth/core/precompile/default_plugin.go) will execute 
the stateless precompiles.

## Stateful Precompiles

Stateful Precompiles are run in the host chain's native execution environment. This is enabled via 
injecting a [Precompile Plugin](https://github.com/berachain/polaris/blob/main/eth/core/precompile/interfaces.go#L33) from the host chain.

Stateful Precompiles can be implemented by adhering to the `StatefulImpl` interface, defined in 
[interfaces.go](https://github.com/berachain/polaris/blob/main/eth/core/precompile/interfaces.go#L56). Below are the suggested steps to follow (more details in [method.go](https://github.com/berachain/polaris/blob/main/eth/core/precompile/method.go)):

    1) Define a Solidity interface with the methods that you want implemented via a precompile.
    2) Build a Go precompile contract, which implements the interface's methods.
        A) This precompile contract should expose the ABI's `Methods`, which can be generated via
        Go-Ethereum's abi package. These methods are of type `abi.Method`.
 	    B) This precompile contract should also expose the `Method`s. A `Method` includes the
        `executable`, which is the direct implementation of a corresponding ABI method, and the ABI signature. Do NOT provide the `AbiMethod` as
        this field will be automatically populated.

Examples of stateful precompiles that run in a Cosmos SDK-based host chain can be found in the
[precompile](https://github.com/berachain/polaris/tree/main/cosmos/precompile) directory.


