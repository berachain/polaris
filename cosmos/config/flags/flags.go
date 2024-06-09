// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package flags

const (
	OptimisticExecution = "polaris.optimistic-execution"

	// Polar Root.
	RPCEvmTimeout = "polaris.polar.rpc-evm-timeout"
	RPCTxFeeCap   = "polaris.polar.rpc-tx-fee-cap"
	RPCGasCap     = "polaris.polar.rpc-gas-cap"

	// Miner.
	MinerEtherbase         = "polaris.polar.miner.etherbase"
	MinerExtraData         = "polaris.polar.miner.extra-data"
	MinerGasFloor          = "polaris.polar.miner.gas-floor"
	MinerGasCeil           = "polaris.polar.miner.gas-ceil"
	MinerGasPrice          = "polaris.polar.miner.gas-price"
	MinerRecommit          = "polaris.polar.miner.recommit"
	MinerNewPayloadTimeout = "polaris.polar.miner.new-payload-timeout"

	// GPO.
	Blocks           = "polaris.polar.gpo.blocks"
	MaxBlockHistory  = "polaris.polar.gpo.max-block-history"
	Percentile       = "polaris.polar.gpo.percentile"
	MaxHeaderHistory = "polaris.polar.gpo.max-header-history"

	// Node.
	JwtSecret             = "polaris.node.jwt-secret" //#nosec: G101 // not a secret.
	WsPort                = "polaris.node.ws-port"
	BatchRequestLimit     = "polaris.node.batch-request-limit"
	KeyStoreDir           = "polaris.node.key-store-dir"
	DBEngine              = "polaris.node.db-engine"
	ReadTimeout           = "polaris.node.http-timeouts.read-timeout"
	DataDir               = "polaris.node.data-dir"
	UserIdent             = "polaris.node.user-ident"
	GraphqlCors           = "polaris.node.graphql-cors"
	SmartCardDaemonPath   = "polaris.node.smart-card-daemon-path"
	WsModules             = "polaris.node.ws-modules"
	HTTPCors              = "polaris.node.http-cors"
	IdleTimeout           = "polaris.node.http-timeouts.idle-timeout"
	AuthAddr              = "polaris.node.auth-addr"
	AllowUnprotectedTxs   = "polaris.node.allow-unprotected-txs"
	HTTPHost              = "polaris.node.http-host"
	UseLightweightKdf     = "polaris.node.use-lightweight-kdf"
	WsExposeAll           = "polaris.node.ws-expose-all"
	InsecureUnlockAllowed = "polaris.node.insecure-unlock-allowed"
	WsPathPrefix          = "polaris.node.ws-path-prefix"
	WsHost                = "polaris.node.ws-host"
	Name                  = "polaris.node.name"
	AuthVirtualHosts      = "polaris.node.auth-virtual-hosts"
	AuthPort              = "polaris.node.auth-port"
	Usb                   = "polaris.node.usb"
	HTTPPort              = "polaris.node.http-port"
	BatchResponseMaxSize  = "polaris.node.batch-response-max-size"
	Version               = "polaris.node.version"
	HTTPVirtualHosts      = "polaris.node.http-virtual-hosts"
	ExternalSigner        = "polaris.node.external-signer"
	HTTPPathPrefix        = "polaris.node.http-path-prefix"
	WriteTimeout          = "polaris.node.http-timeouts.write-timeout"
	ReadHeaderTimeout     = "polaris.node.http-timeouts.read-header-timeout"
	HTTPModules           = "polaris.node.http-modules"
	WsOrigins             = "polaris.node.ws-origins"
	Default               = "polaris.node.http-timeouts.default"
	MaxPrice              = "polaris.node.http-timeouts.max-price"
	IgnorePrice           = "polaris.node.http-timeouts.ignore-price"
	GraphqlVirtualHosts   = "polaris.node.graphql-virtual-hosts"
	IpcPath               = "polaris.node.ipc-path"

	// Legacy TxPool.
	Locals       = "polaris.polar.legacy-tx-pool.locals"
	NoLocals     = "polaris.polar.legacy-tx-pool.no-locals"
	Journal      = "polaris.polar.legacy-tx-pool.journal"
	ReJournal    = "polaris.polar.legacy-tx-pool.rejournal"
	PriceLimit   = "polaris.polar.legacy-tx-pool.price-limit"
	PriceBump    = "polaris.polar.legacy-tx-pool.price-bump"
	AccountSlots = "polaris.polar.legacy-tx-pool.account-slots"
	GlobalSlots  = "polaris.polar.legacy-tx-pool.global-slots"
	AccountQueue = "polaris.polar.legacy-tx-pool.account-queue"
	GlobalQueue  = "polaris.polar.legacy-tx-pool.global-queue"
	Lifetime     = "polaris.polar.legacy-tx-pool.lifetime"

	// Chain Config.
	ChainID                       = "polaris.polar.chain.chain-id"
	HomesteadBlock                = "polaris.polar.chain.homestead-block"
	DAOForkBlock                  = "polaris.polar.chain.dao-fork-block"
	DAOForkSupport                = "polaris.polar.chain.dao-fork-support"
	EIP150Block                   = "polaris.polar.chain.eip150-block"
	EIP155Block                   = "polaris.polar.chain.eip155-block"
	EIP158Block                   = "polaris.polar.chain.eip158-block"
	ByzantiumBlock                = "polaris.polar.chain.byzantium-block"
	ConstantinopleBlock           = "polaris.polar.chain.constantinople-block"
	PetersburgBlock               = "polaris.polar.chain.petersburg-block"
	IstanbulBlock                 = "polaris.polar.chain.istanbul-block"
	MuirGlacierBlock              = "polaris.polar.chain.muir-glacier-block"
	BerlinBlock                   = "polaris.polar.chain.berlin-block"
	LondonBlock                   = "polaris.polar.chain.london-block"
	ArrowGlacierBlock             = "polaris.polar.chain.arrow-glacier-block"
	GrayGlacierBlock              = "polaris.polar.chain.gray-glacier-block"
	MergeNetsplitBlock            = "polaris.polar.chain.merge-netsplit-block"
	ShanghaiTime                  = "polaris.polar.chain.shanghai-time"
	CancunTime                    = "polaris.polar.chain.cancun-time"
	PragueTime                    = "polaris.polar.chain.prague-time"
	VerkleTime                    = "polaris.polar.chain.verkle-time"
	TerminalTotalDifficulty       = "polaris.polar.chain.terminal-total-difficulty"
	TerminalTotalDifficultyPassed = "polaris.polar.chain.terminal-total-difficulty-passed"
)
