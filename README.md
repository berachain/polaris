# <h1 align="center"> Polaris ❄️🔭 </h1>

<a href="https://berachain.com">
  <img src="./docs/web/public/bear_banner.png" width="100%">
</a>

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

> Polaris introduces the new standard for EVM integrations. With improvements to speed, security, reliability, and an extended set of features, Polaris will be able to support the next generation of decentralized applications while offering a compelling alternative to existing implementations.



## Repository Layout

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

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.
