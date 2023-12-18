package mexc

type FuturesAPI interface {
	OrderSubmit(request OrderSubmitRequest) (*OrderSubmitResponse, error)
}

type FuturesClient struct {
	base   string
	key    string
	secret string
}

func NewFuturesClient(base, key, secret string) *FuturesClient {
	return &FuturesClient{
		base:   base,
		key:    key,
		secret: secret,
	}
}
