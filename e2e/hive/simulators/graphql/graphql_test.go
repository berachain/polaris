package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/ethereum/hive/hivesim"
)

type testResponse struct {
	Data gasPrice `json:"data"`
}

type gasPrice struct {
	GasPrice string `json:"gasPrice"`
}

// Test_responseMatch tests whether the graphql tests are able
// to successfully compare a response to an array of valid expected
// responses.
func Test_responseMatch(t *testing.T) {
	// create hivesim tester
	hivesimT := &hivesim.T{}
	// unmarshal JSON test file
	fp := "./testcases/07_eth_gasPrice.json"
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("Warning: can't read test file %s: %v", fp, err)
	}
	var gqlTest graphQLTest
	if err = json.Unmarshal(data, &gqlTest); err != nil {
		t.Fatalf("Warning: can't unmarshal test file %s: %v", fp, err)
	}
	// build test case
	tc := testCase{
		name:    "test1",
		gqlTest: &gqlTest,
	}
	// create valid tests
	var tests = []struct {
		resp            testResponse
		status          string
		expectedFailure bool // true == failure expected
	}{
		{
			resp: testResponse{
				Data: gasPrice{GasPrice: "0x1"},
			},
			status: "200",
		},
		{
			resp: testResponse{
				Data: gasPrice{GasPrice: "0x10"},
			},
			status: "200",
		},
		{
			resp: testResponse{
				Data: gasPrice{GasPrice: "0x12"},
			},
			status:          "400",
			expectedFailure: true,
		},
		{
			resp: testResponse{
				Data: gasPrice{GasPrice: "0x11"},
			},
			status:          "400",
			expectedFailure: true,
		},
		{
			resp: testResponse{
				Data: gasPrice{GasPrice: "failfailfail"},
			},
			status:          "200",
			expectedFailure: true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			resp, err := json.Marshal(tt.resp)
			if err != nil {
				t.Fatal("could not marshal data: ", err)
			}
			err = tc.responseMatch(hivesimT, "200", resp)
			if err != nil && !tt.expectedFailure {
				t.Fatal(err)
			}
		})
	}
}
