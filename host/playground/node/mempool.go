package node

type Mempool interface{}

type mempool struct {
}

func NewMempool() Mempool {
	return &mempool{}
}
