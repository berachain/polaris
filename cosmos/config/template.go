package config

const (
	PolarisConfigTemplate = `

###############################################################################
###                                 Polaris                                 ###
###############################################################################
[[polaris.polar]]
rpc_gas_cap = "{{ .Polaris.Polar.RPCGasCap }}"
rpc_evm_timeout = "{{ .Polaris.Polar.RPCEVMTimeout }}"
rpc_tx_fee_cap = "{{ .Polaris.Polar.RPCTxFeeCap }}"

[[[polaris.polar.gpo]]]
Blocks = {{ .Polaris.Polar.GPO.Blocks }}
Percentile = {{ .Polaris.Polar.GPO.Percentile }}
MaxHeaderHistor = {{ .Polaris.Polar.GPO.MaxHeaderHistory }}
MaxBlockHistory = {{ .Polaris.Polar.GPO.MaxBlockHistory }}
Default = {{ .Polaris.Polar.GPO.Default }}
MaxPrice = {{ .Polaris.Polar.GPO.MaxPrice }}
IgnorePrice = {{ .Polaris.Polar.GPO.IgnorePrice }}
`
)
