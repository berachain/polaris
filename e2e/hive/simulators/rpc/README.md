# RPC Hive

The RPC Hive directory runs the tests specified by the `ethereum/rpc` simulation test spec.

These tests are run as go functions, and required a few changes to comply with Polaris logic.

For example, the estimateGas test needed to be changed to account for the intended overestimation by 20%.

## Files Changed

For the reasons above, we have our own `ethclient.hive`. Since the directory is a subset, this means that not all functions called by ethclient.hive are included in this directory, and proper editing should involve referencing the original file in the Hive repo.

We also maintain our own `init/genesis.json`. This is because Polaris does not support pre-byzantium forks, which is the chain config that the standard `ethereum/rpc` simulation specifies.