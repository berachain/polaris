package graphql

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/node"
)

// Resolver is the top-level object in the GraphQL hierarchy.
type Resolver struct {
	backend      ethapi.Backend
	filterSystem *filters.FilterSystem
}

func RegisterGraphQLService(stack *node.Node, backend ethapi.Backend, filterSystem *filters.FilterSystem, cfg *node.Config) {
	utils.RegisterGraphQLService(stack, backend, filterSystem, cfg)
}
