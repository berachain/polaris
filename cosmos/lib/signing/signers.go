package signing

import (
	"google.golang.org/protobuf/proto"

	"cosmossdk.io/x/tx/signing"
)

// ProvideNoopGetSigners returns a CustomGetSigner that always returns 0x0.
func ProvideNoopGetSigners[T proto.Message]() signing.CustomGetSigner {
	var t T
	return signing.CustomGetSigner{
		MsgType: proto.MessageName(t),
		Fn: func(msg proto.Message) ([][]byte, error) {
			// Return the signer in the required format.
			return [][]byte{{0x0}}, nil
		},
	}
}
