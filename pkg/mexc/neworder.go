package mexc

import (
	"fmt"
	"strconv"
	"time"
)

type NewOrderRequest struct {
	Symbol           string    `json:"symbol"`
	Side             OrderSide `json:"side"`
	OT               OrderType `json:"type"`
	Quantity         *float64  `json:"quantity"`
	QuoteOrderQty    *float64  `json:"quoteOrderQty"`
	Price            *float64  `json:"price"`
	NewClientOrderID *string   `json:"newClientOrderId"`
	RecvWindow       *int      `json:"recvWindow"`
	Timestamp        time.Time `json:"timestamp"`
}

type NewOrderResponse struct {
	Symbol       string    `json:"symbol"`
	OrderID      string    `json:"orderId"`
	OrderListID  int       `json:"orderListId"`
	Price        string    `json:"price"`
	OrigQTY      string    `json:"origQty"`
	OT           OrderType `json:"type"`
	OS           OrderSide `json:"side"`
	TransactTime time.Time `json:"transactTime"`
}

func (c *SpotClient) newOrder(endpoint string, request NewOrderRequest) (*NewOrderResponse, error) {
	req := make(map[string]string, 0)
	req["symbol"] = request.Symbol
	req["side"] = string(request.Side)
	req["type"] = string(request.OT)
	req["timestamp"] = fmt.Sprintf("%d", request.Timestamp.UnixMilli())
	if request.Quantity != nil {
		req["quantity"] = strconv.FormatFloat(*request.Quantity, 'f', -1, 64)
	}
	if request.QuoteOrderQty != nil {
		req["quoteOrderQty"] = strconv.FormatFloat(*request.QuoteOrderQty, 'f', -1, 64)
	}
	if request.Price != nil {
		req["price"] = strconv.FormatFloat(*request.Price, 'f', -1, 64)
	}
	if request.NewClientOrderID != nil {
		req["newClientOrderId"] = *request.NewClientOrderID
	}
	if request.RecvWindow != nil {
		req["recvWindow"] = fmt.Sprintf("%d", *request.RecvWindow)
	}
	response, err := postSignedRequest[NewOrderResponse](
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
