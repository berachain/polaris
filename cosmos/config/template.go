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

//nolint:lll // template file.
package config

const (
	PolarisConfigTemplate = `
###############################################################################
###                                 Polaris                                 ###
###############################################################################
# General Polaris settings
[polaris]
optimistic-execution = {{ .Polaris.OptimisticExecution }}

[polaris.polar]
# Gas cap for RPC requests
rpc-gas-cap = "{{ .Polaris.Polar.RPCGasCap }}"

# Timeout setting for EVM operations via RPC
rpc-evm-timeout = "{{ .Polaris.Polar.RPCEVMTimeout }}"

# Transaction fee cap for RPC requests
rpc-tx-fee-cap = "{{ .Polaris.Polar.RPCTxFeeCap }}"

# Chain config
[polaris.polar.chain] 
chain-id = "{{ .Polaris.Polar.Chain.ChainID }}"

# Homestead switch block
homestead-block = "{{ .Polaris.Polar.Chain.HomesteadBlock }}"

# DAO fork switch block
dao-fork-block = "{{ .Polaris.Polar.Chain.DAOForkBlock }}"

# Whether to support DAO fork
dao-fork-support = {{ .Polaris.Polar.Chain.DAOForkSupport }}

# EIP150 switch block
eip150-block = "{{ .Polaris.Polar.Chain.EIP150Block }}"

# EIP155 switch block
eip155-block = "{{ .Polaris.Polar.Chain.EIP155Block }}"

# EIP158 switch block
eip158-block = "{{ .Polaris.Polar.Chain.EIP158Block }}"

# Byzanitum switch block
byzantium-block = "{{ .Polaris.Polar.Chain.ByzantiumBlock }}"

# Constantinople switch block
constantinople-block = "{{ .Polaris.Polar.Chain.ConstantinopleBlock }}"

# Petersburg switch block
petersburg-block = "{{ .Polaris.Polar.Chain.PetersburgBlock }}"

# Istanbul switch block
istanbul-block = "{{ .Polaris.Polar.Chain.IstanbulBlock }}"

# Muir Glacier switch block
muir-glacier-block = "{{ .Polaris.Polar.Chain.MuirGlacierBlock }}"

# Berlin switch block
berlin-block = "{{ .Polaris.Polar.Chain.BerlinBlock }}"

# London switch block
london-block = "{{ .Polaris.Polar.Chain.LondonBlock }}"

# Arrow Glacier switch block
arrow-glacier-block = "{{ .Polaris.Polar.Chain.ArrowGlacierBlock }}"

# Gray Glacier switch block
gray-glacier-block = "{{ .Polaris.Polar.Chain.GrayGlacierBlock }}"

# Merge Netsplit switch block
merge-netsplit-block = "{{ .Polaris.Polar.Chain.MergeNetsplitBlock }}"

# Shanghai switch time (nil == no fork, 0 = already on shanghai)
shanghai-time = "{{ .Polaris.Polar.Chain.ShanghaiTime }}"

# Cancun switch time (nil == no fork, 0 = already on cancun)
cancun-time = "{{ .Polaris.Polar.Chain.CancunTime }}"

# Prague switch time (nil == no fork, 0 = already on prague)
prague-time = "{{ .Polaris.Polar.Chain.PragueTime }}"

# Verkle switch time (nil == no fork, 0 = already on verkle)
verkle-time = "{{ .Polaris.Polar.Chain.VerkleTime }}"

# Terminal total difficulty
terminal-total-difficulty = "{{ .Polaris.Polar.Chain.TerminalTotalDifficulty }}"

# Whether terminal total difficulty has passed
terminal-total-difficulty-passed = "{{ .Polaris.Polar.Chain.TerminalTotalDifficultyPassed }}"


# Miner config
[polaris.polar.miner]
# The address to which mining rewards will be sent
etherbase = "{{.Polaris.Polar.Miner.Etherbase }}"

# Extra data included in mined blocks
extra-data = "{{.Polaris.Polar.Miner.ExtraData }}"

# Gas price for transactions included in blocks
gas-price = "{{.Polaris.Polar.Miner.GasPrice }}"

# Minimum gas limit for transactions included in blocks
gas-floor = "{{.Polaris.Polar.Miner.GasFloor }}"

# Maximum gas limit for transactions included in blocks
gas-ceil = "{{.Polaris.Polar.Miner.GasCeil }}"

# Whether to enable recommit feature
recommit = "{{.Polaris.Polar.Miner.Recommit }}"

# Timeout for creating a new payload
new-payload-timeout = "{{.Polaris.Polar.Miner.NewPayloadTimeout }}"


# Gas price oracle settings for Polaris
[polaris.polar.gpo]
# Number of blocks to check for gas prices
blocks = "{{ .Polaris.Polar.GPO.Blocks }}"

# Percentile of gas price to use
percentile = "{{ .Polaris.Polar.GPO.Percentile }}"

# Maximum header history for gas price determination
max-header-history = "{{ .Polaris.Polar.GPO.MaxHeaderHistory }}"

# Maximum block history for gas price determination
max-block-history = "{{ .Polaris.Polar.GPO.MaxBlockHistory }}"

# Default gas price value
default = "{{ .Polaris.Polar.GPO.Default }}"

# Maximum gas price value
max-price = "{{ .Polaris.Polar.GPO.MaxPrice }}"

# Prices to ignore for gas price determination
ignore-price = "{{ .Polaris.Polar.GPO.IgnorePrice }}"

# LegacyTxPool settings
[polaris.polar.legacy-tx-pool]

# Addresses that should be treated by default as local
locals = {{ .Polaris.Polar.LegacyTxPool.Locals }}

# Whether local transaction handling should be disabled
no-locals = {{ .Polaris.Polar.LegacyTxPool.NoLocals }}

# Journal of local transactions to survive node restarts
journal = "{{ .Polaris.Polar.LegacyTxPool.Journal }}"

#  Time interval to regenerate the local transaction journal
rejournal = "{{ .Polaris.Polar.LegacyTxPool.Rejournal }}"

# Minimum gas price to enforce for acceptance into the pool
price-limit = "{{ .Polaris.Polar.LegacyTxPool.PriceLimit }}"

# Minimum price bump percentage to replace an already existing transaction (nonce)
price-bump = "{{ .Polaris.Polar.LegacyTxPool.PriceBump }}"

# Number of executable transaction slots guaranteed per account
account-slots = "{{ .Polaris.Polar.LegacyTxPool.AccountSlots }}"

#  Maximum number of executable transaction slots for all accounts
account-queue = "{{.Polaris.Polar.LegacyTxPool.AccountQueue }}"

# Maximum number of non-executable transaction slots permitted per account
global-slots = "{{ .Polaris.Polar.LegacyTxPool.GlobalSlots }}"

# Maximum number of non-executable transaction slots for all accounts
global-queue = "{{ .Polaris.Polar.LegacyTxPool.GlobalQueue }}"

# Maximum amount of time non-executable transaction are queued
lifetime = "{{ .Polaris.Polar.LegacyTxPool.Lifetime }}"


# Node-specific settings
[polaris.node]
# Name of the node
name = "{{ .Polaris.Node.Name }}"

# User identity associated with the node
user-ident = "{{ .Polaris.Node.UserIdent }}"

# Version of the node
version = "{{ .Polaris.Node.Version }}"

# Directory for storing node data
data-dir = "{{ .Polaris.Node.DataDir }}"

# Directory for storing node keys
key-store-dir = "{{ .Polaris.Node.KeyStoreDir }}"

# Path to the external signer
external-signer = "{{ .Polaris.Node.ExternalSigner }}"

# Whether to use lightweight KDF
use-lightweight-kdf = {{ .Polaris.Node.UseLightweightKDF }}

# Allow insecure unlock
insecure-unlock-allowed = {{ .Polaris.Node.InsecureUnlockAllowed }}

# USB setting for the node
usb = {{ .Polaris.Node.USB }}

# Path to smart card daemon
smart-card-daemon-path = "{{ .Polaris.Node.SmartCardDaemonPath }}"

# IPC path for the node
ipc-path = "{{ .Polaris.Node.IPCPath }}"

# Host for HTTP requests
http-host = "{{ .Polaris.Node.HTTPHost }}"

# Port for HTTP requests
http-port = {{ .Polaris.Node.HTTPPort }}

# CORS settings for HTTP
http-cors = [{{ range $index, $element := .Polaris.Node.HTTPCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Virtual hosts for HTTP
http-virtual-hosts = [{{ range $index, $element := .Polaris.Node.HTTPVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Enabled modules for HTTP
http-modules = [{{ range $index, $element := .Polaris.Node.HTTPModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Path prefix for HTTP
http-path-prefix = "{{ .Polaris.Node.HTTPPathPrefix }}"

# Address for authentication
auth-addr = "{{ .Polaris.Node.AuthAddr }}"

# Port for authentication
auth-port = {{ .Polaris.Node.AuthPort }}

# Virtual hosts for authentication
auth-virtual-hosts = [{{ range $index, $element := .Polaris.Node.AuthVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Host for WebSockets
ws-host = "{{ .Polaris.Node.WSHost }}"

# Port for WebSockets
ws-port = {{ .Polaris.Node.WSPort }}

# Path prefix for WebSockets
ws-path-prefix = "{{ .Polaris.Node.WSPathPrefix }}"

# Origins allowed for WebSockets
ws-origins = [{{ range $index, $element := .Polaris.Node.WSOrigins }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Enabled modules for WebSockets
ws-modules = [{{ range $index, $element := .Polaris.Node.WSModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Expose all settings for WebSockets
ws-expose-all = {{ .Polaris.Node.WSExposeAll }}

# CORS settings for GraphQL
graphql-cors = [{{ range $index, $element := .Polaris.Node.GraphQLCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Virtual hosts for GraphQL
graphql-virtual-hosts = [{{ range $index, $element := .Polaris.Node.GraphQLVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]

# Allow unprotected transactions
allow-unprotected-txs = {{ .Polaris.Node.AllowUnprotectedTxs }}

# Limit for batch requests
batch-request-limit = {{ .Polaris.Node.BatchRequestLimit }}

# Maximum size for batch responses
batch-response-max-size = {{ .Polaris.Node.BatchResponseMaxSize }}

# JWT secret for authentication
jwt-secret = "{{ .Polaris.Node.JWTSecret }}"

# Database engine for the node
db-engine = "{{ .Polaris.Node.DBEngine }}"


# HTTP timeout settings for the node
[polaris.node.http-timeouts]
# Timeout for reading HTTP requests
read-timeout = "{{ .Polaris.Node.HTTPTimeouts.ReadTimeout }}"

# Timeout for reading HTTP request headers
read-header-timeout = "{{ .Polaris.Node.HTTPTimeouts.ReadHeaderTimeout }}"

# Timeout for writing HTTP responses
write-timeout = "{{ .Polaris.Node.HTTPTimeouts.WriteTimeout }}"

# Timeout for idle HTTP connections
idle-timeout = "{{ .Polaris.Node.HTTPTimeouts.IdleTimeout }}"
`
)
