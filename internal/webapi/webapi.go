package webapi

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yuriytaranov/mexcd/internal/service"
	"github.com/yuriytaranov/mexcd/internal/webapi/handlers"
)

type Web struct {
	app    service.Service
	server *http.Server
}

func NewWeb(app service.Service) *Web {
	return &Web{
		app: app,
	}
}

func newRouter(app service.Service) *chi.Mux {
	r := chi.NewRouter()
	h := handlers.NewHanders(app)
	r.Get("/", h.Root())
	r.Post("/spot/order/new", h.PostNewOrder())
	r.Post("/spot/order/test", h.PostTestOrder())
	r.Post("/futures/order/submit", h.PostOrderSubmit())

	return r
}

func (w *Web) Run(ctx context.Context, address string) {
	w.server = &http.Server{
		Addr:    address,
		Handler: newRouter(w.app),
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
