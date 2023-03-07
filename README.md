<h1 align="center"> Polaris ❄️🔭 </h1>

<a href="https://berachain.com">
  <img src="./docs/web/public/bear_banner.png" width="100%">
</a>
<h1 class="center">
</h1>

<div>
  <a href="https://codecov.io/gh/berachain/polaris" > 
    <img src="https://codecov.io/gh/berachain/polaris/branch/main/graph/badge.svg?token=5SYYGUS8GW"/> 
  </a>
  <a href="https://pkg.go.dev/pkg.berachain.dev/polaris">
    <img src="https://pkg.go.dev/badge/pkg.berachain.dev/polaris.svg" alt="Go Reference">
  </a>
  <a href="https://magefile.org"> 
    <img alt="Built with Mage" src="https://magefile.org/badge.svg" />
  </a>
  <a href="https://twitter.com/berachain">
    <img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/berachain">
  </a>
</div>
&nbsp;

# Welcome to Polaris

Polaris introduces the new standard for EVM integrations. With improvements to speed, security, reliability, and an extended set of features, Polaris will be able to support the next generation of decentralized applications while offering a compelling alternative to existing implementations. 

This meant that we had to built Polaris with serveral core principles in mind:

1. **Modular**: Every component is built out as a distinct package, central to Polaris, where each component is designed as a library with thorough testing, documentation, and benchmarking. The goal is for developers to use these components by importing them and combine them creatively to integrate an EVM environment into their application.
2. **Configurabile**: We want as many different application frameworks, consensus engines and teams using Polaris as possible. In order to support a wide variety of use cases, Polaris has to be highly configurable.
3. **Performant**: Polaris must be able to perform at the highest level to remain competitive in today's fast-paced and demanding crypto space.
4. **Contributor Friendly**: Depsite currently being BUSL-1.1 licensed, the goal for Polaris is to attract high quality contributors in order to build adoption.
6. **Have Memes**: If ur PR doesn't have a meme in it like idk sry bro, gg wp.

# Repository Layout

> Polaris utilizes [go workspaces](https://go.dev/doc/tutorial/workspaces) to break up the repository into sections to help reduce cognitive overhead.

    .
    ├── build                   # Build scripts and utils
    ├── docs                    # Documentation files
    ├── eth                     # The core Polaris VM implementation
    ├── lib                     # Library files usable throughout the repo
    ├── pkg                     
    │   └── cosmos              # A Cosmos-SDK integration of Polaris.
    │         ├── ....
    │         ├── ....
    │         └── x/evm         # Cosmos-SDK `x/evm` module
    ├── testutil                # Various testing utilities
    ├── LICENSE                 # Licensing information
    └── README.md               # This README


## Build & Test

[Golang 1.20+](https://go.dev/doc/install) and [Foundry](https://book.getfoundry.sh/getting-started/installation) are required for Polaris.

1. Install [Go 1.20+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   cd $HOME
   sudo apt-get install golang -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   cd $HOME
   brew install go
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin

2. Install Foundry:
   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

3. Clone, Setup and Test:

```bash
cd $HOME
git clone https://github.com/berachain/polaris
cd polaris
git checkout main
go run build/setup.go
mage test
```


## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.
