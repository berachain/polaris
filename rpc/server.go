package main

import (
	"fmt"
	"net/http/httptest"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
)

// Have a type with some exported methods
type SimpleServerHandler struct {
	n int
}

func (h *SimpleServerHandler) AddGet(in int) int {
	h.n += in
	return h.n
}

func main() {
	// create a new server instance
	rpcServer := jsonrpc.NewServer()

	// create a handler instance and register it
	serverHandler := &SimpleServerHandler{}
	rpcServer.Register("SimpleServerHandler", serverHandler)
	// rpcServer is now http.Handler which will serve jsonrpc calls to SimpleServerHandler.AddGet
	// a method with a single int param, and an int response. The server supports both http and websockets.
	// serve the api

	testServ := httptest.NewServer(rpcServer)
	defer testServ.Close()
	fmt.Println("URL: ", "ws://"+testServ.Listener.Addr().String(), testServ.URL)
	select {}
}
