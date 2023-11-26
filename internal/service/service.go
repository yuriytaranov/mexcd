package service

import "github.com/yuriytaranov/mexcd/pkg/mexc"

type Service interface{}

type Application struct {
	client mexc.API
}

func NewApplication(client mexc.API) *Application {
	return &Application{
		client: client,
	}
}
