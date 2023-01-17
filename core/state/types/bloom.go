package types

type BloomBuilder struct {
}

// We utilize transient stores to cache the bloom filter for the current block. This is done to
// prevent the bloom filter from having to be fully recalculated from scratch for every transaction
// thus turning an O(n^2) operation into an O(n) operation. By nature of storing the bloom filter in a transient store
// the bloom filter is reset at the end of every block.
func NewBloomBuilder() *BloomBuilder {
	return &BloomBuilder{}
}

// Todo: store completed eth blocks in app state?
