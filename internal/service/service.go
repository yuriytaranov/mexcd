package service

import "github.com/yuriytaranov/mexcd/pkg/mexc"

type Service interface {
	NewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error)
	TestNewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error)
	OrderSubmit(request mexc.OrderSubmitRequest) (*mexc.OrderSubmitResponse, error)
}

type Application struct {
	spot    mexc.SpotAPI
	futures mexc.FuturesAPI
}

func NewApplication(spot mexc.SpotAPI, futures mexc.FuturesAPI) *Application {
	return &Application{
		spot:    spot,
		futures: futures,
	}
}

func (a *Application) OrderSubmit(request mexc.OrderSubmitRequest) (*mexc.OrderSubmitResponse, error) {
	return a.futures.OrderSubmit(request)
}

func (a *Application) NewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error) {
	return a.spot.NewOrder(request)
}

func (a *Application) TestNewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error) {
	return a.spot.TestNewOrder(request)
}
