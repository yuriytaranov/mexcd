package mexc

import (
	"fmt"
	"time"
)

type HistoryPositionsRequest struct {
	Symbol       *string
	PositionType *int
	PageNum      int
	PageSize     int
}

type (
	HistoryPositionsDataResponse struct {
		PositionId     int       `json:"positionId"`
		Symbol         string    `json:"symbol"`
		PositionType   int       `json:"positionType"`
		OpenType       int       `json:"openType"`
		State          int       `json:"state"`
		HoldVol        float64   `json:"holdVol"`
		FrozenVol      float64   `json:"frozenVol"`
		CloseVol       float64   `json:"closeVol"`
		HoldAvgPrice   float64   `json:"holdAvgPrice"`
		OpenAvgPrice   float64   `json:"openAvgPrice"`
		CloseAvgPrice  float64   `json:"closeAvgPrice"`
		LiquidatePrice float64   `json:"liquidatePrice"`
		OIM            float64   `json:"oim"`
		IM             float64   `json:"im"`
		HoldFee        float64   `json:"holdFee"`
		Realised       float64   `json:"realised"`
		AdlLevel       int       `json:"adlLevel"`
		Leverage       int       `json:"leverage"`
		CreateTime     time.Time `json:"createTime"`
		UpdateTime     time.Time `json:"updateTime"`
		AutoAddIM      bool      `json:"autoAddIm"`
	}

	HistoryPositionsResponse struct {
		Success bool                           `json:"success"`
		Code    int                            `json:"code"`
		Message string                         `json:"message"`
		Data    []HistoryPositionsDataResponse `json:"data"`
	}
)

func (c *FuturesClient) HistoryPositions(request HistoryPositionsRequest) (*HistoryPositionsResponse, error) {
	req := make(map[string]string, 2)
	if request.Symbol != nil {
		req["symbol"] = *request.Symbol
	}
	if request.PositionType != nil {
		req["type"] = fmt.Sprintf("%d", *request.PositionType)
	}
	req["page_num"] = fmt.Sprintf("%d", request.PageNum)
	req["page_size"] = fmt.Sprintf("%d", request.PageSize)

	return nil, nil
}
