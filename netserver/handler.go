package netserver

import (
	"github.com/influxdata/influxdb/dp/point"
	"github.com/influxdata/influxdb/dp/result"
	"net/http"
)

type Handler struct {
	InfService interface {
		Write(pts point.Points) error
		Query(qa string) (result.Result, error)
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("method") {
	case "write":
		h.serveWrite(w, r)
	case "query":
		h.serveQuery(w, r)
	}
}

func (h *Handler) serveWrite(w http.ResponseWriter, r *http.Request) {
	err := h.InfService.Write()
}
func (h *Handler) serveQuery(w http.ResponseWriter, r *http.Request) {
	res, err := h.InfService.Query()
}
