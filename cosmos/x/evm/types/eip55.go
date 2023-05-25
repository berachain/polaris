package types

import (
	fmt "fmt"

	"cosmossdk.io/core/address"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ address.Codec

type EIP55AddressEncoder struct {
	// contains filtered or unexported fields
}

func NewEIP55AddressEncoder() *EIP55AddressEncoder {
	return &EIP55AddressEncoder{}
}

func (e EIP55AddressEncoder) BytesToString(address []byte) (string, error) {
	fmt.Println("LETS GOOOO 1")
	return common.Bytes2Hex(address), nil
}

func (e EIP55AddressEncoder) StringToBytes(address string) ([]byte, error) {
	fmt.Println("LETS GOOOO 2")
	return common.Hex2Bytes(address), nil
}
