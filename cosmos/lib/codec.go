package lib

import "cosmossdk.io/core/address"

type CodecProvider interface {
	AddressCodec() address.Codec
}
