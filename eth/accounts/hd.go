package accounts

var (
	// `Bip44CoinType` was defined by EIP-84 (https://github.com/ethereum/EIPs/issues/84).
	Bip44CoinType uint32 = 60

	// `BIP44HDPath` is the default BIP44 HD path used on Ethereum.
	BIP44HDPath = DefaultBaseDerivationPath.String()
)

// `HDPathIterator` creates a generic iterator for a given base path.
func NewHDPathIterator(basePath string, ledgerIter bool) (HDPathIterator, error) {
	hdPath, err := ParseDerivationPath(basePath)
	if err != nil {
		return nil, err
	}

	if ledgerIter {
		return LedgerLiveIterator(hdPath), nil
	}

	return DefaultIterator(hdPath), nil
}
