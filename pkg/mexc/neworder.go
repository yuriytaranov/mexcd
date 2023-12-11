package mexc

import (
	"fmt"
	"time"
)

type NewOrderRequest struct {
	Symbol           string
	Side             OrderSide
	OT               OrderType
	Quantity         *float64
	QuoteOrderQty    *float64
	Price            *float64
	NewClientOrderID *string
	RecvWindow       *int
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

func (c *SpotClient) newOrder(endpoint string, request NewOrderRequest) (*NewOrderResponse, error) {
	req := make(map[string]any, 0)
	req["symbol"] = request.Symbol
	req["side"] = request.Side
	req["type"] = request.OT
	req["timestamp"] = request.Timestamp.Unix()
	if request.Quantity != nil {
		req["quantity"] = *request.Quantity
	}
	if request.QuoteOrderQty != nil {
		req["quoteOrderQty"] = *request.QuoteOrderQty
	}
	if request.Price != nil {
		req["price"] = *request.Price
	}
	if request.NewClientOrderID != nil {
		req["newClientOrderId"] = *request.Price
	}
	if request.RecvWindow != nil {
		req["recvWindow"] = *request.RecvWindow
	}
	response, err := postRequest[NewOrderResponse](
		endpoint,
		req,
		c.key,
		c.secret,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *SpotClient) TestNewOrder(request NewOrderRequest) (*NewOrderResponse, error) {
	return c.newOrder(fmt.Sprintf("%s%s", c.base, "/api/v3/order/test"), request)
}

func (c *SpotClient) NewOrder(request NewOrderRequest) (*NewOrderResponse, error) {
	return c.newOrder(fmt.Sprintf("%s%s", c.base, "/api/v3/order"), request)
}
