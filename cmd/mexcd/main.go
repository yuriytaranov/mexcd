package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yuriytaranov/mexcd/internal/service"
	"github.com/yuriytaranov/mexcd/internal/webapi"
	"github.com/yuriytaranov/mexcd/pkg/mexc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	app := service.NewApplication(
		mexc.NewSpotClient(os.Getenv("MEXC_SPOT_BASE"), os.Getenv("MEXC_KEY"), os.Getenv("MEXC_SECRET")),
		mexc.NewFuturesClient(os.Getenv("MEXC_FUTURES_BASE"), os.Getenv("MEXC_KEY"), os.Getenv("MEXC_SECRET")),
	)

	api := webapi.NewWeb(app)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		api.Run(ctx, os.Getenv("WEBAPI_ADDR"))
	}()

	wg.Add(1)
	go func() {
		<-done
		cancel()
		defer wg.Done()
	}()

	wg.Wait()
}
