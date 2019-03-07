package infserver

import (
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/platform/infserver/metaservice"
	"github.com/influxdata/influxdb/platform/infserver/metaservice/raft_store"
	"github.com/influxdata/influxdb/platform/infserver/queryservice"
	"github.com/influxdata/influxdb/platform/infserver/transervice"
	"github.com/influxdata/influxdb/platform/infserver/writeservice"
	"github.com/influxdata/influxdb/query"
	"github.com/influxdata/influxdb/tcp"
	"net"
)

type Server struct {
	node   Node
	nodes  map[string]Node
	closed chan struct{}

	Listener net.Listener

	Mux      *tcp.Mux
	MetaSvc  *metaservice.MetaService
	TranSvc  *transervice.TranService
	WriteSvc *writeservice.WriteService
	QuerySvc *queryservice.QueryService
}

func New(c *Config) *Server {
	node := c.Nodes[c.Name]
	s := &Server{
		node:   node,
		closed: make(chan struct{}),
	}

	s.Mux = tcp.NewMux()
	s.MetaSvc = metaservice.New(c.Meta, c.Name, c.Nodes, c.Path, s.Mux.Listen(raft_store.MuxHeader))
	s.TranSvc = transervice.New()
	s.WriteSvc = writeservice.New(s.MetaSvc, s.TranSvc)
	s.QuerySvc = queryservice.New()
	return s
}

func (s *Server) Open() error {
	if err := s.MetaSvc.Open(); err != nil {
		return err
	}

	if err := s.TranSvc.Open(); err != nil {
		return err
	}
	if err := s.WriteSvc.Open(); err != nil {
		return nil
	}
	if err := s.QuerySvc.Open(); err != nil {
		return nil
	}
	return nil
}

func (s *Server) Write(pts models.Points) error {
	return s.WriteSvc.Write(pts)
}

func (s *Server) Query(qa string) (query.Result, error) {
	return s.QuerySvc.Query(qa)
}
