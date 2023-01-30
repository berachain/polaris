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

package jsonrpc

import (
	"github.com/cosmos/cosmos-sdk/client"
	jsonrpc "github.com/filecoin-project/go-jsonrpc"
)

// Have a type with some exported methods.
type SimpleServerHandler struct {
	n int
}

func (h *SimpleServerHandler) AddGet(in int) int {
	h.n += in
	return h.n
}

type Server struct {
	*jsonrpc.RPCServer
	clientCtx client.Context
}

func NewServer(clientCtx client.Context) *Server {
	rpcServer := jsonrpc.NewServer()
	serverHandler := &SimpleServerHandler{}
	rpcServer.Register("SimpleServerHandler", serverHandler)
	return &Server{
		RPCServer: rpcServer,
		clientCtx: clientCtx,
	}
}

// func main() {
// 	// create a handler instance and register it
// 	serverHandler := &SimpleServerHandler{}
// 	rpcServer.Register("SimpleServerHandler", serverHandler)
// 	// rpcServer is now http.Handler which will serve jsonrpc calls to SimpleServerHandler.AddGet
// 	// a method with a single int param, and an int response. The server supports both http and websockets.
// 	// serve the api

// 	testServ := httptest.NewServer(rpcServer)
// 	defer testServ.Close()
// 	fmt.Println("URL: ", "ws://"+testServ.Listener.Addr().String(), testServ.URL)
// 	select {}
// }
