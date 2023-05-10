package graphql

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/node"
)

// use PolarisProvider

func RegisterGraphQLService(stack *node.Node, backend ethapi.Backend, filterSystem *filters.FilterSystem, cfg *node.Config) {
	utils.RegisterGraphQLService(stack, backend, filterSystem, cfg)
}

//func idk() *ethapi.Backend {
//return &ethapi.Backend{}
//}

//func ethApi() *ethapi.EthereumAPI {
//var t graphql.Long
//t.ImplementsGraphQLType("hello")
//ret := &ethapi.EthereumAPI{b: nil}

//return ret
//}

//type GraphQL interface {
//}
