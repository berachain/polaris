# Hive GraphQL Execution Tests

`init` contains the genesis file that the Hive instance tests on. 
`testcases` contains all of the requests and expected responses. 

When running `mage hive:setup hive:testv polaris/graphql polard`, on your machine, an instance of Hive starts and tests locally against your `testcases` folder and checks that each GraphQL request results in the expected response. 

No additional files are needed as most of the core logic for running the tests are contained within the Hive instance, which, as previously mentioned, is already set up for you after running `mage hive:setup`.

# How to run the tests

You must build a Polard image first, do so by running `mage cosmos:testHive polaris/graphql`. This builds a fresh Docker image for Polard in which to run your tests on. The tests will run after this. 

After building an image, if you want to modify a testcase or the genesis, simply run `mage hive:setup hive:test polaris/graphql polard` after making your changes. 

NOTE: If you change any of the core Polaris EVM logic in `./eth/...`, you must rebuild your Polard image. 


