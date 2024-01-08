<h1 align="center"> Polaris Monorepo ❄️🔭 </h1>

*프로젝트는 아직 진행 중이며, 아래의 [경고문](#-경고-공사-중-)을 참조하십시오.*

<div>
  <a href="https://codecov.io/gh/berachain/polaris" target="_blank">
    <img src="https://codecov.io/gh/berachain/polaris/branch/main/graph/badge.svg?token=5SYYGUS8GW"/>
  </a>
  <a href="https://pkg.go.dev/github.com/berachain/polaris" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/berachain/polaris.svg" alt="Go Reference">
  </a>
  <a href="https://t.me/polaris_devs" target="_blank">
    <img alt="Telegram Chat" src="https://img.shields.io/endpoint?color=neon&logo=telegram&label=chat&url=https%3A%2F%2Ftg.sumanjay.workers.dev%2Fpolaris_devs">
  </a>
  <a href="https://twitter.com/berachain" target="_blank">
    <img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/berachain">
  <a href="https://discord.gg/berachain">
   <img src="https://img.shields.io/discord/984015101017346058?color=%235865F2&label=Discord&logo=discord&logoColor=%23fff" alt="Discord">
  </a>
</div>

## Polaris란 무엇인가요?

Polaris를 소개합니다, 이것은 이더리움 가상 머신(EVM)을 애플리케이션에 통합을 단순화하도록 설계된 혁신적인 프레임워크입니다. Polaris는 개발자가 자신의 EVM 통합 솔루션을 함께 해킹하는 데 시간을 보내는 것을 제거하는 깔끔하고 쉽게 통합할 수 있는 API로 구축되었습니다. 우리의 프레임워크는 매우 모듈화되어 있어, 당신이 가장 필요로 하는 구성 요소를 선택하고 거의 모든 애플리케이션에 EVM 환경을 통합할 수 있습니다.

Polaris는 몇 가지 핵심 원칙을 염두에 두고 구축되었습니다:

1. **모듈성**: 각 구성 요소는 완전한 테스트, 문서화, 벤치마킹과 함께 별도의 패키지로 개발됩니다. 이러한 구성 요소를 개별적으로 사용하거나 결합하여 혁신적인 EVM 통합을 만들 수 있습니다.
2. **구성 가능성**: 우리는 Polaris가 가능한 많은 팀과 사용 사례에 접근할 수 있도록 하고 싶습니다. 이를 지원하기 위해, 우리의 프레임워크는 매우 구성 가능하며, 당신이 특정 필요에 맞게 조정할 수 있습니다.
3. **성능**: 오늘날의 경쟁력 있는 암호화 풍경에서 성능은 핵심입니다. Polaris는 최고 수준의 성능과 효율성을 제공하도록 최적화되었습니다.
4. **기여자 친화성**: 우리는 블록체인 개발에서 혁신을 추진하는 데 열린 협업이 핵심이라고 믿습니다. Polaris는 현재 BUSL-1.1에 따라 라이선스가 부여되지만, 우리는 생산 준비를 접근함에 따라 기여자 기반의 체계를 지원하기 위해 우리의 라이선싱을 조정할 계획입니다.
1. **Modularity**: Each component is developed as a distinct package, complete with thorough testing, documentation, and benchmarking. You can use these components individually or combine them to create innovative EVM integrations.
2. **Configurability**: We want Polaris to be accessible to as many teams and use cases as possible. To support this, our framework is highly configurable, allowing you to tailor it to your specific needs.
3. **Performance**: In today's competitive crypto landscape, performance is key. Polaris is optimized to deliver the highest levels of performance and efficiency.
4. **Contributor Friendliness**: We believe that open collaboration is key to driving innovation in blockchain development. While Polaris is currently licensed under BUSL-1.1, we plan to adjust our licensing to support contributor-based schemes as we approach production readiness.
5. **Memes**: If ur PR doesn't have a meme in it like idk sry bro, gg wp glhf.

## Documentation

If you want to build on top of Polaris, take a look at our [documentation](http://polaris.berachain.dev/).
If you want to help contribute to the framework, check out the [Framework Specs](./specs/).

## Directory Structure

> Polaris utilizes [go workspaces](https://go.dev/doc/tutorial/workspaces) to break up the repository into logical sections, helping to reduce cognitive overhead.

<pre>
🔭 Polaris 🔭
├── <a href="./build">build</a>: Build scripts and developer tooling.
├── <a href="./contracts">contracts</a>: Contracts and bindings for Polaris (and hosts).
├── <a href="./cosmos">cosmos</a>: Polaris integrated into a Cosmos-SDK based chain.
├── <a href="./e2e">e2e</a>: End-to-end testing utilities.
├── <a href="./eth">eth</a>: The Core of the Polaris Ethereum Framework.
├── <a href="./lib">lib</a>: A collection of libraries used throughout the repo.
├── <a href="./proto">proto</a>: Protobuf definitions.
</pre>

## Build & Test

[Golang 1.20+](https://go.dev/doc/install) and [Foundry](https://book.getfoundry.sh/getting-started/installation) are required for Polaris.

1. Install [go 1.21+ from the official site](https://go.dev/dl/) or the method of your choice. Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Ubuntu:

   ```sh
   cd $HOME
   sudo apt-get install golang jq -y
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   cd $HOME
   brew install go jq
   export PATH=$PATH:/opt/homebrew/bin/go
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Install Foundry:

   ```sh
   curl -L https://foundry.paradigm.xyz | bash
   ```

3. Clone, Setup and Test:

   ```sh
   cd $HOME
   git clone https://github.com/berachain/polaris
   cd polaris
   git checkout main
   make test-unit
   ```

4. Start a local development network:

   ```sh
   make start
   ```

## 🚧 WARNING: UNDER CONSTRUCTION 🚧

This project is work in progress and subject to frequent changes as we are still working on wiring up the final system.
It has not been audited for security purposes and should not be used in production yet.

The network will have an Ethereum JSON-RPC server running at `http://localhost:8545` and a Tendermint RPC server running at `http://localhost:26657`.
