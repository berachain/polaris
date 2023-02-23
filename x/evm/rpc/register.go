package rpc

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// `RegisterJSONRPCServer` provides a common function which registers the ethereum rpc servers
// with routes on the native Cosmos API Server.
func RegisterJSONRPCServer(ctx client.Context, rtr *mux.Router, provider Provider) error {
	if !provider.Ready() {
		return errors.New("JSONRPC provider not ready")
	}
	httpSrv := provider.GetHTTP()
	wsSrv := provider.GetWS()
	rtr.PathPrefix(httpPath).Handler(httpSrv)
	rtr.PathPrefix(httpPath + "/").Handler(httpSrv)
	rtr.PathPrefix(wsPath).Handler(wsSrv)
	rtr.PathPrefix(wsPath + "/").Handler(wsSrv)
	provider.SetClientContext(ctx)
	return nil
}
