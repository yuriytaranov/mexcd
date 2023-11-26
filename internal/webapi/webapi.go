package webapi

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yuriytaranov/mexcd/internal/service"
)

type Web struct {
	app    service.Service
	server *http.Server
	router *chi.Mux
}

func NewWeb(app service.Service) *Web {
	return &Web{
		app: app,
	}
}

func (w *Web) Run(ctx context.Context, address string) {
	w.router = chi.NewRouter()

	w.server = &http.Server{
		Addr:    address,
		Handler: w.router,
	}

	go func() {
		if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to listen and serve %v\n", err)
		}
	}()
	log.Printf("web server started %s", address)
	<-ctx.Done()
	log.Println("web server stopping")

	waitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := w.server.Shutdown(waitCtx); err != nil {
		log.Printf("failed to shutdown web server %v\n", err)
	}

	log.Println("web server exited")
}
