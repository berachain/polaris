//nolint:lll
//go:generate abigen --abi ../out/staking.sol/IStakingModule.abi.json --pkg staking --type Interface --out ../staking/contract.abigen.go
//go:generate abigen --abi ../out/staking.sol/StakingEvents.abi.json --pkg staking --type Events --out ../staking/events.abigen.go
package staking
