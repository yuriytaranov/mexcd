package mexc

import (
	"fmt"
	"strconv"
)

type OrderSubmitRequest struct {
	Symbol          string               `json:"symbol"`
	Price           float64              `json:"price"`
	Vol             float64              `json:"vol"`
	Leverage        *int                 `json:"leverage,omitempty"`
	Side            FuturesOrderSide     `json:"side"`
	OrderType       FuturesOrderType     `json:"type"`
	OpenType        FuturesOpenType      `json:"openType"`
	PositionID      *int64               `json:"positionId,omitempty"`
	ExternalOID     *string              `json:"externalOid,omitempty"`
	StopLossPrice   *float64             `json:"stopLossPrice,omitempty"`
	TakeProfitPrice *float64             `json:"takeProfitPrice,omitempty"`
	PositionMode    *FuturesPositionMode `json:"positionMode,omitempty"`
	ReduceOnly      *bool                `json:"reduceOnly,omitempty"`
}

type OrderSubmitResponse struct {
	Success bool  `json:"success"`
	Code    int   `json:"code"`
	Data    int64 `json:"data"`
}

func (c *FuturesClient) OrderSubmit(request OrderSubmitRequest) (*OrderSubmitResponse, error) {
	req := make(map[string]string, 0)
	req["symbol"] = request.Symbol
	req["price"] = strconv.FormatFloat(request.Price, 'f', -1, 64)
	req["vol"] = strconv.FormatFloat(request.Vol, 'f', -1, 64)
	if request.Leverage != nil {
		req["leverage"] = fmt.Sprintf("%d", *request.Leverage)
	}
	req["side"] = fmt.Sprintf("%d", request.Side)
	req["type"] = fmt.Sprintf("%d", request.OrderType)
	req["openType"] = fmt.Sprintf("%d", request.OpenType)
	if request.PositionID != nil {
		req["positionId"] = fmt.Sprintf("%d", *request.PositionID)
	}
	if request.ExternalOID != nil {
		req["externalOid"] = *request.ExternalOID
	}
	if request.StopLossPrice != nil {
		req["stopLossPrice"] = strconv.FormatFloat(*request.StopLossPrice, 'f', -1, 64)
	}
	if request.TakeProfitPrice != nil {
		req["takeProfitPrice"] = strconv.FormatFloat(*request.TakeProfitPrice, 'f', -1, 64)
	}
	if request.PositionMode != nil {
		req["positionMode"] = fmt.Sprintf("%d", *request.PositionMode)
	}
	if request.ReduceOnly != nil {
		req["reduceOnly"] = strconv.FormatBool(*request.ReduceOnly)
	}
	response, err := postSignedFuturesRequest[OrderSubmitResponse](
		fmt.Sprintf("%s%s", c.base, "/api/v1/private/order/submit"),
		req,
		c.key,
		c.secret,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
