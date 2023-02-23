package crypto

import (
	"errors"
	"math/big"

	"pkg.berachain.dev/stargazer/eth/common"
)

// `RecoverPubkey` recovers the public key from the given signature.
func RecoverPubkey(sighash common.Hash, R, S, Vb *big.Int, homestead bool) ([]byte, error) {
	if Vb.BitLen() > 8 {
		return nil, errors.New("invalid signature length")
	}
	V := byte(Vb.Uint64() - 27)
	if !ValidateSignatureValues(V, R, S, homestead) {
		return nil, errors.New("invalid signature values")
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := Ecrecover(sighash[:], sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, errors.New("invalid public key")
	}
	return pub, nil
}
