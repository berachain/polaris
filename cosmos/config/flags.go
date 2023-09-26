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

package config

const (
	flagGraphqlVirtualHosts = "polaris.node.graphql-virtual-hosts"
	flagIpcPath             = "polaris.node.ipc-path"

	//#nosec: G101 // not a secret.
	flagJwtSecret             = "polaris.node.jwt-secret"
	flagWsPort                = "polaris.node.ws-port"
	flagBatchRequestLimit     = "polaris.node.batch-request-limit"
	flagKeyStoreDir           = "polaris.node.key-store-dir"
	flagDBEngine              = "polaris.node.db-engine"
	flagReadTimeout           = "polaris.node.http-timeouts.read-timeout"
	flagDataDir               = "polaris.node.data-dir"
	flagUserIdent             = "polaris.node.user-ident"
	flagBlocks                = "polaris.polar.gpo.blocks"
	flagGraphqlCors           = "polaris.node.graphql-cors"
	flagSmartCardDaemonPath   = "polaris.node.smart-card-daemon-path"
	flagWsModules             = "polaris.node.ws-modules"
	flagHTTPCors              = "polaris.node.http-cors"
	flagIdleTimeout           = "polaris.node.http-timeouts.idle-timeout"
	flagAuthAddr              = "polaris.node.auth-addr"
	flagAllowUnprotectedTxs   = "polaris.node.allow-unprotected-txs"
	flagHTTPHost              = "polaris.node.http-host"
	flagUseLightweightKdf     = "polaris.node.use-lightweight-kdf"
	flagWsExposeAll           = "polaris.node.ws-expose-all"
	flagMaxBlockHistory       = "polaris.polar.gpo.max-block-history"
	flagPercentile            = "polaris.polar.gpo.percentile"
	flagInsecureUnlockAllowed = "polaris.node.insecure-unlock-allowed"
	flagWsPathPrefix          = "polaris.node.ws-path-prefix"
	flagWsHost                = "polaris.node.ws-host"
	flagName                  = "polaris.node.name"
	flagRPCEvmTimeout         = "polaris.polar.rpc-evm-timeout"
	flagAuthVirtualHosts      = "polaris.node.auth-virtual-hosts"
	flagAuthPort              = "polaris.node.auth-port"
	flagUsb                   = "polaris.node.usb"
	flagHTTPPort              = "polaris.node.http-port"
	flagBatchResponseMaxSize  = "polaris.node.batch-response-max-size"
	flagVersion               = "polaris.node.version"
	flagHTTPVirtualHosts      = "polaris.node.http-virtual-hosts"
	flagRPCTxFeeCap           = "polaris.polar.rpc-tx-fee-cap"
	flagMaxHeaderHistory      = "polaris.polar.gpo.max-header-history"
	flagExternalSigner        = "polaris.node.external-signer"
	flagHTTPPathPrefix        = "polaris.node.http-path-prefix"
	flagWriteTimeout          = "polaris.node.http-timeouts.write-timeout"
	flagReadHeaderTimeout     = "polaris.node.http-timeouts.read-header-timeout"
	flagHTTPModules           = "polaris.node.http-modules"
	flagRPCGasCap             = "polaris.polar.rpc-gas-cap"
	flagWsOrigins             = "polaris.node.ws-origins"
	flagDefault               = "polaris.node.http-timeouts.default"
	flagMaxPrice              = "polaris.node.http-timeouts.max-price"
	flagIgnorePrice           = "polaris.node.http-timeouts.ignore-price"
	flagLocals                = "polaris.polar.legacy-tx-pool.locals"
	flagNoLocals              = "polaris.polar.legacy-tx-pool.no-locals"
	flagJournal               = "polaris.polar.legacy-tx-pool.journal"
	flagReJournal             = "polaris.polar.legacy-tx-pool.rejournal"
	flagPriceLimit            = "polaris.polar.legacy-tx-pool.price-limit"
	flagPriceBump             = "polaris.polar.legacy-tx-pool.price-bump"
	flagAccountSlots          = "polaris.polar.legacy-tx-pool.account-slots"
	flagGlobalSlots           = "polaris.polar.legacy-tx-pool.global-slots"
	flagAccountQueue          = "polaris.polar.legacy-tx-pool.account-queue"
	flagGlobalQueue           = "polaris.polar.legacy-tx-pool.global-queue"
	flagLifetime              = "polaris.polar.legacy-tx-pool.lifetime"
)
