package rpc

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/gorilla/mux"

	ethrpc "pkg.berachain.dev/stargazer/eth/rpc"
	ethrpcconfig "pkg.berachain.dev/stargazer/eth/rpc/config"
)

const (
	httpPath = "/eth/rpc"
	wsPath   = "/eth/rpc"
)

type Provider interface {
	GetHTTP() *ethrpc.Server
	GetWS() *ethrpc.Server
	Ready() bool
}

type provider struct {
	ethrpc.Service
}

// `NewProvider` returns a new `Provider` object. The provider object is used to
// register the JSON-RPC servers with the API server.
func NewProvider(cfg ethrpcconfig.Server, backend ethapi.Backend) Provider {
	service, err := ethrpc.NewService(cfg, backend)
	if err != nil {
		panic(err)
	}
	return &provider{
		*service,
	}
}

func (p *provider) Ready() bool {
	// TODO: there is likely a race condition.
	return true
}

// RegisterSwaggerAPI provides a common function which registers swagger route with API Server
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
	return nil
}
