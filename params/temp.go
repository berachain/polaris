package params

// ExtraEIPs represents extra EIPs for the vm.Config
type ExtraEIPs struct {
	// eips defines the additional EIPs for the vm.Config
	EIPs []int64 `protobuf:"varint,1,rep,packed,name=eips,proto3" json:"eips,omitempty" yaml:"eips"`
}
