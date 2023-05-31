package main

import "math/big"

func consistentChainIDTest(t *TestEnv) {
	var (
		expectedChainID = big.NewInt(7)
	)

	chainID, err := t.Eth.ChainID(t.Ctx())
	if err != nil {
		t.Fatalf("could not get chain ID: %v", err)
	}

	if expectedChainID.Cmp(chainID) != 0 {
		t.Fatalf("expected chain ID %d, got %d", expectedChainID, chainID)
	}
}
