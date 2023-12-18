package mexc

type OrderSide string

const (
	OrderSideSell OrderSide = "SELL"
	OrderSideBuy  OrderSide = "BUY"
)

type OrderType string

const (
	OrderTypeLimit             OrderType = "LIMIT"
	OrderTypeMarket            OrderType = "MARKET"
	OrderTypeLimitMaker        OrderType = "LIMIT_MAKER"
	OrderTypeImmediateOrCancel OrderType = "IMMEDIATE_OR_CANCEL"
	OrderTypeFillOrKill        OrderType = "FILL_OR_KILL"
)

type FuturesOpenType int

const (
	FuturesOpenTypeIsolated FuturesOpenType = 1
	FuturesOpenTypeCross    FuturesOpenType = 2
)

type FuturesPositionMode int

const (
	FuturesPositionModeHedge  FuturesPositionMode = 1
	FuturesPositionModeOneWay FuturesPositionMode = 2
)

type FuturesOrderSide int

const (
	FuturesOrderSideOpenLong   FuturesOrderSide = 1
	FuturesOrderSideCloseShort FuturesOrderSide = 2
	FuturesOrderSideOpenShort  FuturesOrderSide = 3
	FuturesOrderSideCloseLong  FuturesOrderSide = 4
)

type FuturesOrderType int

const (
	FuturesOrderTypePriceLimited                         FuturesOrderType = 1
	FuturesOrderTypePostOnlyMarker                       FuturesOrderType = 2
	FuturesOrderTypeTransactOrCancelInstantly            FuturesOrderType = 3
	FuturesOrderTypeTransactCompletelyOrCancelCompletely FuturesOrderType = 4
	FuturesOrderTypeMarketOrders                         FuturesOrderType = 5
	FuturesOrderTypeConvertMarketPriceToCurrentPrice     FuturesOrderType = 6
)
