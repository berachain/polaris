// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	jsonrpc "github.com/berachain/stargazer/jsonrpc"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "json-rpc",
	Args:  cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
	Short: "Foundry contract generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Retrieve the Cosmos Context from the cobra command.
		ctx := client.GetClientContextFromCmd(cmd)

		// Create a new Stargazer JSON-RPC server.
		rpcServer := jsonrpc.NewServer(ctx)

		// TODO: move into `./jsonrpc` and add configuration file.
		httpSrv := &http.Server{
			Addr:              "localhost:8545",
			Handler:           rpcServer,
			ReadHeaderTimeout: time.Second,
			WriteTimeout:      time.Second,
		}
		httpSrvDone := make(chan struct{}, 1)
		errCh := make(chan error)
		go func() {
			// TODO: proper logger
			fmt.Println("Starting JSON-RPC server at:", httpSrv.Addr) //nolint: forbidigo // temp.
			if err := httpSrv.ListenAndServe(); err != nil {
				if err == http.ErrServerClosed {
					close(httpSrvDone)
					return
				}
				// TODO: proper logger
				//nolint: forbidigo // temp.
				fmt.Println("failed to start JSON-RPC server", "error", err.Error())
				errCh <- err
			}
		}()
		return <-errCh
	},
}
