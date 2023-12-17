package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yuriytaranov/mexcd/internal/service"
	"github.com/yuriytaranov/mexcd/pkg/mexc"
)

type Handlers struct {
	app service.Service
}

func NewHanders(app service.Service) *Handlers {
	return &Handlers{
		app: app,
	}
}

func (h *Handlers) Root() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(chi.URLParam(r, "test_param")))
	}
}

func (h *Handlers) PostTestOrder() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		symbol := r.Form.Get("symbol")
		side := r.Form.Get("side")
		ot := r.Form.Get("type")
		quantity := r.Form.Get("quantity")
		quoteOrderQty := r.Form.Get("quoteOrderQty")
		price := r.Form.Get("price")
		newClientOrderID := r.Form.Get("newClientOrderId")
		recvWindow := r.Form.Get("recvWindow")

		if symbol == "" || side == "" || ot == "" {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "one of symbol side ot empty")))
			return
		}

		req := mexc.NewOrderRequest{
			Symbol:    symbol,
			Side:      mexc.OrderSide(side),
			OT:        mexc.OrderType(ot),
			Timestamp: time.Now(),
		}

		if quantity != "" {
			val, err := strconv.ParseFloat(quantity, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse quantity")))
				return
			}

			req.Quantity = &val
		}

		if quoteOrderQty != "" {
			val, err := strconv.ParseFloat(quoteOrderQty, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse quoteOrderQty")))
				return
			}

			req.QuoteOrderQty = &val
		}

		if price != "" {
			val, err := strconv.ParseFloat(price, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse price")))
				return
			}

			req.Price = &val
		}

		if newClientOrderID != "" {
			req.NewClientOrderID = &newClientOrderID
		}

		if recvWindow != "" {
			val, err := strconv.Atoi(recvWindow)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse recvWindow")))
				return
			}

			req.RecvWindow = &val
		}

		resp, err := h.app.TestNewOrder(req)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", err.Error())))
			return
		}

		respj, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", err.Error())))
			return
		}

		w.Write(respj)
	}
}

func (h *Handlers) PostNewOrder() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		symbol := r.Form.Get("symbol")
		side := r.Form.Get("side")
		ot := r.Form.Get("type")
		quantity := r.Form.Get("quantity")
		quoteOrderQty := r.Form.Get("quoteOrderQty")
		price := r.Form.Get("price")
		newClientOrderID := r.Form.Get("newClientOrderId")
		recvWindow := r.Form.Get("recvWindow")

		if symbol == "" || side == "" || ot == "" {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "one of symbol side ot empty")))
			return
		}

		req := mexc.NewOrderRequest{
			Symbol:    symbol,
			Side:      mexc.OrderSide(side),
			OT:        mexc.OrderType(ot),
			Timestamp: time.Now(),
		}

		if quantity != "" {
			val, err := strconv.ParseFloat(quantity, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse quantity")))
				return
			}

			req.Quantity = &val
		}

		if quoteOrderQty != "" {
			val, err := strconv.ParseFloat(quoteOrderQty, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse quoteOrderQty")))
				return
			}

			req.QuoteOrderQty = &val
		}

		if price != "" {
			val, err := strconv.ParseFloat(price, 64)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse price")))
				return
			}

			req.Price = &val
		}

		if newClientOrderID != "" {
			req.NewClientOrderID = &newClientOrderID
		}

		if recvWindow != "" {
			val, err := strconv.Atoi(recvWindow)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", "failed to parse recvWindow")))
				return
			}

			req.RecvWindow = &val
		}

		resp, err := h.app.NewOrder(req)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", err.Error())))
			return
		}

		respj, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", err.Error())))
			return
		}

		w.Write(respj)
	}
}
