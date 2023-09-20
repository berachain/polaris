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
[polaris]
[polaris.polar]
rpc-gas-cap = "{{ .Polaris.Polar.RPCGasCap }}"
rpc-evm-timeout = "{{ .Polaris.Polar.RPCEVMTimeout }}"
rpc-tx-fee-cap = "{{ .Polaris.Polar.RPCTxFeeCap }}"

[polaris.polar.gpo]
blocks = {{ .Polaris.Polar.GPO.Blocks }}
percentile = {{ .Polaris.Polar.GPO.Percentile }}
max-header-history = {{ .Polaris.Polar.GPO.MaxHeaderHistory }}
max-block-history = {{ .Polaris.Polar.GPO.MaxBlockHistory }}
default = "{{ .Polaris.Polar.GPO.Default }}"
max-price = "{{ .Polaris.Polar.GPO.MaxPrice }}"
ignore-price = "{{ .Polaris.Polar.GPO.IgnorePrice }}"

[polaris.node]
name = "{{ .Polaris.Node.Name }}"
user-ident = "{{ .Polaris.Node.UserIdent }}"
version = "{{ .Polaris.Node.Version }}"
data-dir = "{{ .Polaris.Node.DataDir }}"
key-store-dir = "{{ .Polaris.Node.KeyStoreDir }}"
external-signer = "{{ .Polaris.Node.ExternalSigner }}"
use-lightweight-kdf = {{ .Polaris.Node.UseLightweightKDF }}
insecure-unlock-allowed = {{ .Polaris.Node.InsecureUnlockAllowed }}
usb = {{ .Polaris.Node.USB }}
smart-card-daemon-path = "{{ .Polaris.Node.SmartCardDaemonPath }}"
ipc-path = "{{ .Polaris.Node.IPCPath }}"
http-host = "{{ .Polaris.Node.HTTPHost }}"
http-port = {{ .Polaris.Node.HTTPPort }}
http-cors = [{{ range $index, $element := .Polaris.Node.HTTPCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
http-virtual-hosts = [{{ range $index, $element := .Polaris.Node.HTTPVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
http-modules = [{{ range $index, $element := .Polaris.Node.HTTPModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
http-path-prefix = "{{ .Polaris.Node.HTTPPathPrefix }}"
auth-addr = "{{ .Polaris.Node.AuthAddr }}"
auth-port = {{ .Polaris.Node.AuthPort }}
auth-virtual-hosts = [{{ range $index, $element := .Polaris.Node.AuthVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
ws-host = "{{ .Polaris.Node.WSHost }}"
ws-port = {{ .Polaris.Node.WSPort }}
ws-path-prefix = "{{ .Polaris.Node.WSPathPrefix }}"
ws-origins = [{{ range $index, $element := .Polaris.Node.WSOrigins }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
ws-modules = [{{ range $index, $element := .Polaris.Node.WSModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
ws-expose-all = {{ .Polaris.Node.WSExposeAll }}
graphql-cors = [{{ range $index, $element := .Polaris.Node.GraphQLCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
graphql-virtual-hosts = [{{ range $index, $element := .Polaris.Node.GraphQLVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
allow-unprotected-txs = {{ .Polaris.Node.AllowUnprotectedTxs }}
batch-request-limit = {{ .Polaris.Node.BatchRequestLimit }}
batch-response-max-size = {{ .Polaris.Node.BatchResponseMaxSize }}
jwt-secret = "{{ .Polaris.Node.JWTSecret }}"
db-engine = "{{ .Polaris.Node.DBEngine }}"

[polaris.node.http-timeouts]
read-timeout = "{{ .Polaris.Node.HTTPTimeouts.ReadTimeout }}"
read-header-timeout = "{{ .Polaris.Node.HTTPTimeouts.ReadHeaderTimeout }}"
write-timeout = "{{ .Polaris.Node.HTTPTimeouts.WriteTimeout }}"
idle-timeout = "{{ .Polaris.Node.HTTPTimeouts.IdleTimeout }}"
`
)
