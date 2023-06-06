package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"gotest.tools/assert"
)

type HTTPReq struct {
	body string
	want string
	code int
}

func graphQLChainIdSupport(t *TestEnv) {
	var (
		expectedChainID = big.NewInt(7) //nolint:gomnd // TODO: REFACTOR.
	)

	//	query := `
	//		query {
	//			block {
	//				number
	//			}
	//		}`

	cID, err := t.Eth.ChainID(t.Ctx())
	assert.NilError(t, err, "could not get chain ID: %w", err)

	if expectedChainID.Cmp(cID) != 0 {
		t.Fatalf("expected chain ID %d, got %d", expectedChainID, cID)
	}
}

func graphQLGetLatestBlockSupport(t *TestEnv) {
	query := `{"query": "{block{number}}","variables": null}`
	var result interface{}
	t.CallContext(t.Ctx(), &result, query)
	fmt.Println("RESULT: ", result)
	sendHTTP(t, HTTPReq{body: query, want: `{"data":{"block":{"number":"0x1"}}}`, code: http.StatusOK})
}

//func graphQLBlockNumberSupport(t *TestEnv) {
//	var (
//		expectedBlockNumber = big.NewInt(1) //nolint:gomnd // TODO: REFACTOR.
//	)
//
//	bN, err := t.Eth.BlockNumber(t.Ctx())
//	assert.NilError(t, err, "could not get block number: %w", err)
//
//	if expectedBlockNumber.Cmp(bN) != 0 {
//		t.Fatalf("expected block number %d, got %d", expectedBlockNumber, bN)
//	}
//}

func graphQLGasPriceSupport(t *TestEnv) {
	// query := `{"query": "{block{number}}","variables": null}`
	// var result interface{}
	// fmt.Println("RESULT:", result)
}

func graphQLGetBlockByHashSupport(t *TestEnv) {}

func graphQLAccountDataSupport(t *TestEnv) {}

func graphQLGetTransactionByBlockHashAndIndex(t *TestEnv) {}

func graphQLGetLogsSupport(t *TestEnv) {}

func graphQLSendRawTransactionSupport(t *TestEnv) {}

func graphQLSyncingSupport(t *TestEnv) {}

func graphQLFailMalformatted(t *TestEnv) {}

// TODO use t.CallContext(...)
func sendHTTP(t *TestEnv, tt HTTPReq) (*http.Response, error) {

	resp, err := http.Post("http://localhost:8545/graphql", "application/json", strings.NewReader(tt.body))
	if err != nil {
		t.Fatalf("could not post: %v", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("could not read from response body: %v", err)
	}
	if have := string(bodyBytes); have != tt.want {
		t.Errorf("testcase %s,\nhave:\n%v\nwant:\n%v", tt.body, have, tt.want)
	}
	if tt.code != resp.StatusCode {
		t.Errorf("testcase %s,\nwrong statuscode, have: %v, want: %v", tt.body, resp.StatusCode, tt.code)
	}

	return resp, nil
}
