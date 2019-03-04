package netserver

import (
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/query"
	"net/http"
)

type Handler struct {
	InfService interface {
		Write(pts models.Points) error
		Query(qa string) (query.Result, error)
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
