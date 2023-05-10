package mock

//go:generate moq -out ./geth.mock.go -pkg mock ./ PrecompileEVM MessageRouter

func NewPrecompileEVMMock() *PrecompileEVMMock {
	return &PrecompileEVMMock{}
}

func NewMsgRouterMock() *MessageRouterMock {
	return &MessageRouterMock{}
}
