package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/hive/hivesim"
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
		filepath.Walk("./testcases", func(filepath string, info os.FileInfo, err error) error {
			if limit >= 0 && i >= limit {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if fname := info.Name(); !strings.HasSuffix(fname, ".json") {
				return nil
			}
			data, err := ioutil.ReadFile(filepath)
			if err != nil {
				t.Logf("Warning: can't read test file %s: %v", filepath, err)
				return nil
			}
			var gqlTest graphQLTest
			if err = json.Unmarshal(data, &gqlTest); err != nil {
				t.Logf("Warning: can't unmarshal test file %s: %v", filepath, err)
				return nil
			}
			i = i + 1
			t := testCase{
				name:    strings.TrimSuffix(info.Name(), path.Ext(info.Name())),
				gqlTest: &gqlTest,
			}
			out <- &t
			return nil
		})
		close(out)
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

// prepareRunTest administers the hive-specific test stuff, registering the suite and reporting back the suite results
func (tc *testCase) run(t *hivesim.T, c *hivesim.Client) {
	// Example of working queries:
	// curl 'http://127.0.0.1:8545/graphql' --data-binary '{"query":"query blockNumber {\n  block {\n    number\n  }\n}\n"}'
	// curl 'http://127.0.0.1:8545/graphql' --data-binary '{"query":"query blockNumber {\n  block {\n    number\n  }\n}\n","variables":null,"operationName":"blockNumber"}'
	postData, err := json.Marshal(qlQuery{Query: tc.gqlTest.Request})
	if err != nil {
		t.Fatal("can't marshal query:", err)
	}
	url := fmt.Sprintf("http://%v:8545/graphql", c.IP)
	resp, err := http.Post(url, "application/json", bytes.NewReader(postData))
	if err != nil {
		t.Fatal("HTTP post failed:", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("can't read HTTP response:", err)
	}
	resp.Body.Close()

	if resp.StatusCode != tc.gqlTest.StatusCode {
		t.Errorf("HTTP response code is %d, want %d \n response body: %s", resp.StatusCode, tc.gqlTest.StatusCode, string(respBytes))
	}
	if resp.StatusCode != 200 {
		// Test expects HTTP error, and the client sent one, test done.
		// We don't bother to check the exact error messages, those aren't fully specified.
		return
	}

	tc.responseMatch(t, resp.Status, respBytes)
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
		}
	}

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
