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
