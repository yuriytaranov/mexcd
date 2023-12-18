package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yuriytaranov/mexcd/pkg/mexc"
)

func (h *Handlers) PostOrderSubmit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mexc.OrderSubmitRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Write([]byte(fmt.Sprintf("{error: \"%s\"}", err.Error())))
			return
		}

		resp, err := h.app.OrderSubmit(req)
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
