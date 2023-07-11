package codec

import (
	"cosmossdk.io/core/address"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ address.Codec = (*EIP55Address)(nil)

type EIP55Address struct{}

func (e EIP55Address) StringToBytes(text string) ([]byte, error) {
	return common.FromHex(text), nil
}

func (e EIP55Address) BytesToString(bz []byte) (string, error) {
	return common.BytesToAddress(bz).Hex(), nil
}
