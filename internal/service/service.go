package service

import "github.com/yuriytaranov/mexcd/pkg/mexc"

type Service interface {
	NewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error)
}

type Application struct {
	client mexc.APISpot
}

func NewApplication(client mexc.APISpot) *Application {
	return &Application{
		client: client,
	}
}

func (a *Application) NewOrder(request mexc.NewOrderRequest) (*mexc.NewOrderResponse, error) {
	return a.client.NewOrder(request)
}
