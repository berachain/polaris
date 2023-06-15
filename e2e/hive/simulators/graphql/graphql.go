// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/hive/hivesim"

	"github.com/ethereum/go-ethereum/params"
)

func main() {
	suite := hivesim.Suite{
		Name: "graphql",
		Description: `Test suite covering the graphql API surface.
The GraphQL tests were initially imported from the Besu codebase.`,
	}
	suite.Add(hivesim.ClientTestSpec{
		Role: "eth1",
		Name: "client launch",
		Description: `This is a meta-test. It launches the client with the test chain
and reads the test case files. The individual test cases are run as sub-tests against
the client launched by this test.`,
		Parameters: hivesim.Params{
			// The graphql chain comes from the Besu codebase, and is built on Frontier.
			"HIVE_CHAIN_ID":             "1",
			"HIVE_GRAPHQL_ENABLED":      "1",
			"HIVE_ALLOW_UNPROTECTED_TX": "1",
		},
		Files: map[string]string{
			"/genesis.json": "./init/testGenesis.json",
		},
		Run: graphqlTest,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func graphqlTest(t *hivesim.T, c *hivesim.Client) {
	parallelism := 16
	if val, ok := os.LookupEnv("HIVE_PARALLELISM"); ok {
		if p, err := strconv.Atoi(val); err != nil {
			t.Logf("Warning: invalid HIVE_PARALLELISM value %q", val)
		} else {
			parallelism = p
		}
	}

	// wait for blocks...
	delay := 3
	time.Sleep(time.Duration(delay) * time.Second)
	var wg sync.WaitGroup
	testCh := deliverTests(t, &wg, -1)
	for i := 0; i < parallelism; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for test := range testCh {
				url := "https://github.com/ethereum/hive/blob/master/simulators/ethereum/graphql/testcases"
				t.Run(hivesim.TestSpec{
					Name:        fmt.Sprintf("%s (%s)", test.name, c.Type),
					Description: fmt.Sprintf("Test case source: %s/%v.json", url, test.name),
					Run:         func(t *hivesim.T) { test.run(t, c) },
				})
			}
		}()
	}
	wg.Wait()
}

// deliverTests reads the test case files, sending them to the output channel.
func deliverTests(t *hivesim.T, wg *sync.WaitGroup, limit int) <-chan *testCase {
	out := make(chan *testCase)
	var i = 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := filepath.Walk("./testcases", func(filePath string, info os.FileInfo, err error) error {
			switch {
			case limit >= 0 && i >= limit:
				return nil
			case info.IsDir():
				return nil
			case !strings.HasSuffix(info.Name(), ".json"):
				return nil
			case err != nil:
				return err
			}
			filePath = filepath.Join("./", filepath.Clean(filePath))
			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Logf("Warning: can't read test file %s: %v", filePath, err)
				return nil
			}
			var gqlTest graphQLTest
			if err = json.Unmarshal(data, &gqlTest); err != nil {
				t.Logf("Warning: can't unmarshal test file %s: %v", filePath, err)
				return nil
			}
			i++
			t := testCase{
				name:    strings.TrimSuffix(info.Name(), path.Ext(info.Name())),
				gqlTest: &gqlTest,
			}
			out <- &t
			return nil
		})
		close(out)
		if err != nil {
			t.Logf("Warning: can't read test files: %v", err)
		}
	}()
	return out
}

type testCase struct {
	name    string
	gqlTest *graphQLTest
}

// graphQLTest is the JSON object structure of a test case file.
type graphQLTest struct {
	Request    string        `json:"request"`
	Responses  []interface{} `json:"responses"`
	StatusCode int           `json:"statusCode"`
}

type qlQuery struct {
	Query string `json:"query"`
}

// prepareRunTest administers the hive-specific test stuff, registering the suite and reporting back the suite results.
func (tc *testCase) run(t *hivesim.T, c *hivesim.Client) {
	// Example of working queries:
	// curl 'http://127.0.0.1:8545/graphql'
	//--data-binary '{"query":"query blockNumber {\n
	// block {\n
	//    number\n
	//  }\n
	// }\n
	// "}'

	// curl 'http://127.0.0.1:8545/graphql' --data-binary '{"query":"query blockNumber {\n
	//  block {\n
	//    number\n
	//  }\n
	// }\n
	// ","variables":null,"operationName":"blockNumber"}'
	postData, err := json.Marshal(qlQuery{Query: tc.gqlTest.Request})
	if err != nil {
		t.Fatal("can't marshal query:", err)
	}
	u := fmt.Sprintf("http://%v:8545/graphql", c.IP)
	parsedURL, err := url.Parse(u)
	if err != nil {
		t.Fatal("can't parse URL:", err)
	}
	resp, err := http.Post(parsedURL.String(), "application/json", bytes.NewReader(postData)) //nolint: noctx, lll // hive team wrote this
	if err != nil {
		t.Fatal("HTTP post failed:", err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("can't read HTTP response:", err)
	}

	err = resp.Body.Close()
	if err != nil {
		t.Fatal("can't close HTTP response body:", err)
	}

	if resp.StatusCode != tc.gqlTest.StatusCode {
		t.Errorf("HTTP response code is %d, want %d \n response body: %s",
			resp.StatusCode, tc.gqlTest.StatusCode, string(respBytes))
	}
	if resp.StatusCode != http.StatusOK {
		// Test expects HTTP error, and the client sent one, test done.
		// We don't bother to check the exact error messages, those aren't fully specified.
		return
	}

	err = tc.responseMatch(t, resp.Status, respBytes)
	if err != nil {
		t.Errorf("Could not run tests. Error: %v", err)
	}
}

func (tc *testCase) responseMatch(t *hivesim.T, respStatus string, respBytes []byte) error {
	// Check that the response matches.
	var got interface{}
	if err := json.Unmarshal(respBytes, &got); err != nil {
		t.Fatal("can't decode response:", err)
	}
	// return if a response matches. If not, error out.
	for _, response := range tc.gqlTest.Responses {
		if reflect.DeepEqual(response, got) {
			return nil
		} else if err := assertGasPrice(t, got); err == nil {
			return nil
		} else if err := assertBlockNumber(t, got); err == nil {
			return nil
		}
	}

	// this is to make sure that gasPrice is above the initialBaseFee
	// TODO: move this out, very specific to the gasPrice test

	prettyQuery, ok := reindentJSON(tc.gqlTest.Request)
	prettyResponse, _ := json.MarshalIndent(got, "", "  ")

	t.Log("Test failed.")
	t.Log("HTTP response code:", respStatus)
	if ok {
		t.Log("query:", prettyQuery)
	}
	t.Log("expected value(s):")

	for _, expected := range tc.gqlTest.Responses {
		prettyExpected, _ := json.MarshalIndent(expected, "", "  ")
		t.Log(string(prettyExpected), "\n_____________________\n")
	}

	t.Log("got:", string(prettyResponse))
	t.Fail()

	return fmt.Errorf("test failed")
}

func reindentJSON(text string) (string, bool) {
	var obj interface{}
	if json.Unmarshal([]byte(text), &obj) != nil {
		return "", false
	}
	indented, _ := json.MarshalIndent(&obj, "", "  ")
	return string(indented), true
}

func assertBlockNumber(t *hivesim.T, got interface{}) error {
	if data, ok := got.(map[string]interface{}); ok {
		inner, dataOk := data["data"].(map[string]interface{})
		if !dataOk {
			t.Fail()
		}
		blockData, ok := inner["block"].(map[string]interface{})
		if !ok {
			t.Fail()
		}
		number, ok := blockData["number"].(string)
		if !ok {
			t.Fail()
		}
		bn, err := strconv.ParseInt(number, 0, 64)
		if err != nil {
			t.Fail()
		}
		if bn > 0 {
			return nil
		}
	}
	return fmt.Errorf("block height is not greater than 0")
}

func assertGasPrice(t *hivesim.T, got interface{}) error {
	var initialBaseFee = int64(params.InitialBaseFee)
	if data, ok := got.(map[string]interface{}); ok {
		inner, dataOk := data["data"].(map[string]interface{})
		if !dataOk {
			t.Fail()
		}
		gasPrice, gasPriceOk := inner["gasPrice"].(string)
		if !gasPriceOk {
			t.Fail()
		}
		gp, err := strconv.ParseInt(gasPrice, 0, 64)
		if err != nil {
			t.Fail()
		}
		if gp > initialBaseFee {
			return nil
		}
	}
	return fmt.Errorf("gas price is not greater than initial base fee")
}
