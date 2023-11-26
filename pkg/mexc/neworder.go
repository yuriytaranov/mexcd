package mexc

import "time"

type NewOrderRequest struct {
	Symbol           string
	Side             OrderSide
	OT               OrderType
	Quantity         float64
	QuoteOrderQty    float64
	Price            float64
	NewClientOrderID string
	RecvWindow       int
	Timestamp        time.Time
}

type NewOrderResponse struct {
	Symbol       string
	OrderID      string
	OrderListID  int
	Price        string
	OrigQTY      string
	OT           OrderType
	OS           OrderSide
	TransactTime time.Time
}

func (c *SpotClient) NewOrder(request NewOrderRequest) (*NewOrderResponse, error) {
	return nil, nil
}
