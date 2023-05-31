package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/hive/hivesim"
)

type testSpec struct {
	Name  string
	About string
	Run   func(*TestEnv)
}

var (
	// parameters used for signing transactions
	chainID   = big.NewInt(7)
	gasPrice  = big.NewInt(30 * params.GWei)
	networkID = big.NewInt(7)

	files = map[string]string{
		"/genesis.json": "./init/genesis.json",
	}

	clientEnv = hivesim.Params{
		"HIVE_NETWORK_ID": networkID.String(),
		"HIVE_CHAIN_ID":   chainID.String(),
	}
)

var tests = []testSpec{
	{Name: "http/ConsistentChainIDTest", Run: consistentChainIDTest},
}

func main() {
	suite := hivesim.Suite{
		Name: "rpc",
		Description: `The RPC test suite runs a set of RPC related tests against a running node. It tests
several real-world scenarios such as sending value transactions, deploying a contract or
interacting with one.`[1:],
	}

	suite.Add(&hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "client launch",
		Description: `This test launches the client and collects its logs.`,
		Parameters:  clientEnv,
		Files:       files,
		Run:         func(t *hivesim.T, c *hivesim.Client) { runAllTests(t, c, c.Type) },
		AlwaysRun:   true,
	})

	sim := hivesim.New()
	hivesim.MustRunSuite(sim, suite)
}

// runAllTests runs the tests against a client instance.
// Most tests simply wait for tx inclusion in a block so we can run many tests concurrently.
func runAllTests(t *hivesim.T, c *hivesim.Client, clientName string) {
	vault := newVault()

	s := newSemaphore(16)
	for _, test := range tests {
		test := test
		s.get()
		go func() {
			defer s.put()
			t.Run(hivesim.TestSpec{
				Name:        fmt.Sprintf("%s (%s)", test.Name, clientName),
				Description: test.About,
				Run: func(t *hivesim.T) {
					switch test.Name[:strings.IndexByte(test.Name, '/')] {
					case "http":
						runHTTP(t, c, vault, test.Run)
					case "ws":
						runWS(t, c, vault, test.Run)
					default:
						panic("bad test prefix in name " + test.Name)
					}
				},
			})
		}()
	}
	s.drain()
}

type semaphore chan struct{}

func newSemaphore(n int) semaphore {
	s := make(semaphore, n)
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
	return s
}

func (s semaphore) get() { <-s }
func (s semaphore) put() { s <- struct{}{} }

func (s semaphore) drain() {
	for i := 0; i < cap(s); i++ {
		<-s
	}
}
