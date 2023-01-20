package types

// TODO: Remove the need to do the whole conversion dance.

func NewProtoLogs(ethlogs []*EthLog) []*Log {
	logs := make([]*Log, 0, len(ethlogs))
	for _, ethlog := range ethlogs {
		logs = append(logs, NewLogFromEth(ethlog))
	}

	return logs
}

// NewLogFromEth creates a new Log instance from a Ethereum type Log.
func NewLogFromEth(log *EthLog) *Log {
	topics := make([][]byte, len(log.Topics))
	for i, topic := range log.Topics {
		topics[i] = topic[:]
	}

	return &Log{
		Address:     log.Address[:],
		Topics:      topics,
		Data:        log.Data,
		BlockNumber: log.BlockNumber,
		TxHash:      log.TxHash.String(),
		TxIndex:     uint64(log.TxIndex),
		Index:       uint64(log.Index),
		BlockHash:   log.BlockHash.String(),
		Removed:     log.Removed,
	}
}
