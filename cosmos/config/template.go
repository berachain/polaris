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
RPCGasCap = "{{ .Polaris.Polar.RPCGasCap }}"
RPCEVMTimeout = "{{ .Polaris.Polar.RPCEVMTimeout }}"
RPCTxFeeCap = "{{ .Polaris.Polar.RPCTxFeeCap }}"

[polaris.polar.gpo]
Blocks = {{ .Polaris.Polar.GPO.Blocks }}
Percentile = {{ .Polaris.Polar.GPO.Percentile }}
MaxHeaderHistor = {{ .Polaris.Polar.GPO.MaxHeaderHistory }}
MaxBlockHistory = {{ .Polaris.Polar.GPO.MaxBlockHistory }}

[polaris.node]
Name = "{{ .Polaris.Node.Name }}"
UserIdent = "{{ .Polaris.Node.UserIdent }}"
Version = "{{ .Polaris.Node.Version }}"
DataDir = "{{ .Polaris.Node.DataDir }}"
KeyStoreDir = "{{ .Polaris.Node.KeyStoreDir }}"
ExternalSigner = "{{ .Polaris.Node.ExternalSigner }}"
UseLightweightKDF = {{ .Polaris.Node.UseLightweightKDF }}
InsecureUnlockAllowed = {{ .Polaris.Node.InsecureUnlockAllowed }}
NoUSB = {{ .Polaris.Node.NoUSB }}
USB = {{ .Polaris.Node.USB }}
SmartCardDaemonPath = "{{ .Polaris.Node.SmartCardDaemonPath }}"
IPCPath = "{{ .Polaris.Node.IPCPath }}"
HTTPHost = "{{ .Polaris.Node.HTTPHost }}"
HTTPPort = {{ .Polaris.Node.HTTPPort }}
HTTPCors = [{{ range $index, $element := .Polaris.Node.HTTPCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
HTTPVirtualHosts = [{{ range $index, $element := .Polaris.Node.HTTPVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
HTTPModules = [{{ range $index, $element := .Polaris.Node.HTTPModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
HTTPPathPrefix = "{{ .Polaris.Node.HTTPPathPrefix }}"
AuthAddr = "{{ .Polaris.Node.AuthAddr }}"
AuthPort = {{ .Polaris.Node.AuthPort }}
AuthVirtualHosts = [{{ range $index, $element := .Polaris.Node.AuthVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
WSHost = "{{ .Polaris.Node.WSHost }}"
WSPort = {{ .Polaris.Node.WSPort }}
WSPathPrefix = "{{ .Polaris.Node.WSPathPrefix }}"
WSOrigins = [{{ range $index, $element := .Polaris.Node.WSOrigins }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
WSModules = [{{ range $index, $element := .Polaris.Node.WSModules }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
WSExposeAll = {{ .Polaris.Node.WSExposeAll }}
GraphQLCors = [{{ range $index, $element := .Polaris.Node.GraphQLCors }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
GraphQLVirtualHosts = [{{ range $index, $element := .Polaris.Node.GraphQLVirtualHosts }}{{ if $index }}, {{ end }}"{{ $element }}"{{ end }}]
AllowUnprotectedTxs = {{ .Polaris.Node.AllowUnprotectedTxs }}
BatchRequestLimit = {{ .Polaris.Node.BatchRequestLimit }}
BatchResponseMaxSize = {{ .Polaris.Node.BatchResponseMaxSize }}
JWTSecret = "{{ .Polaris.Node.JWTSecret }}"
DBEngine = "{{ .Polaris.Node.DBEngine }}"

[polaris.node.http-timeouts]
ReadTimeout = "{{ .Polaris.Node.HTTPTimeouts.ReadTimeout }}"
ReadHeaderTimeout = "{{ .Polaris.Node.HTTPTimeouts.ReadHeaderTimeout }}"
WriteTimeout = "{{ .Polaris.Node.HTTPTimeouts.WriteTimeout }}"
IdleTimeout = "{{ .Polaris.Node.HTTPTimeouts.IdleTimeout }}"
`
)
