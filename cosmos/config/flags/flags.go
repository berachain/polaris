// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
