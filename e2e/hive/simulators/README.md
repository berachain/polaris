# Hive Simulators

Each Hive simulator maintained in the Polaris repo is a subset of the respective Hive simulator maintained in the Hive repo. The goal internally is to maintain the minimum number of files necessary to run the Hive simulator in Polaris.

## Usage

Since we do not maintain the whole directory, there is no way (at the moment) to run the Hive simulator directly. Instead, to run it locally, `hive` must be invoked via the `Makefile`.

### Steps
1. run `make hive-setup` to generate the hive clone and creates the polaris namespace simulations
2. run `make docker-build` to build the polard base image
3. run `make test-hive` to run the simulation on the given client

Note: polaris namespace simulations are called `polaris/<name>`, for the sake of maintaining consistency with `ethereum/<name>` simulations.

## Adding simulations

Currently, the existing makefile setup only supports maintaining simulations which are forks of an existing simulation in the Hive repo.


## .hive file extension

The `.hive` file extension used in the rpc simulation is the local replacement file for it's respective copy in the standard ethereum simulation folder.

The copy is used to make minor changes to the expected values of certain test cases to be compliant with the way that Polaris handles returns (ex. estimateGas).

Internally, this file replaces the destination in the Hive clone generated during `make hive-setup`.

The motivation to keeping the file in plain text is to eliminate errors generated by dependencies of the file and to remove the need to create an additional go package within Polaris.