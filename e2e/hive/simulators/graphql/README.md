# GraphQL Hive Integration Tests

Tests were taken from https://github.com/ethereum/hive/tree/master/simulators/ethereum/graphql/testcases

NOTE: these tests do not test against `hash`, `miner`, or any other state root dependent fields because we do not have a state root or `chain.rlp`. It is important to note that these do work, but there is not a good way to test them (yet) as these fields change with every test run.