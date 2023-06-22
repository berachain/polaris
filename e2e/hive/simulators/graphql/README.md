# Hive GraphQL Execution Tests

`init` contains the genesis file that the Hive instance tests on. 
`testcases` contains all of the requests and expected responses. 

When running `mage hive:setup hive:testv polaris/graphql polard`, on your machine, an instance of Hive starts and tests locally against your `testcases` folder and checks that each GraphQL request results in the expected response. 

No additional files are needed as most of the core logic for running the tests are contained within the Hive instance, which, as previously mentioned, is already set up for you after running `mage hive:setup`.