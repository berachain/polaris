# RPC-Compat Hive

The RPC-Compat Hive directory runs the tests specified by the `ethereum/rpc-compat` simulation test spec.

The rpc-compat simulation is automatically run as a part of the rpc simulation test, thus its exclusion from the CI. However, they can also be run in isolation if desired.

## Files Changed

`Dockerfile`: The tests are normally cloned into the repo at runtime by `ethereum/rpc-compat` in the docker image setup, but since we are running our own test cases, the dockerfile has been modified to accommodate for that.

`tests`: These tests come from `https://github.com/ethereum/execution-apis`, but since Polaris does not support `chain.rlp`, some tests have been modified locally.