package mock

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
)

//go:generate moq -out ./host.mock.go -pkg mock ../ StargazerHostChain

func NewMockHost() *StargazerHostChainMock {
	// make and configure a mocked core.StargazerHostChain
	mockedStargazerHostChain := &StargazerHostChainMock{
		StargazerHeaderAtHeightFunc: func(contextMoqParam context.Context, v uint64) *types.StargazerHeader {
			return &types.StargazerHeader{
				Header: &types.Header{
					Number: big.NewInt(int64(v)),
				},
				CachedHash: common.Hash{},
			}
		},
	}
	return mockedStargazerHostChain
}
