package tsdbserver

import (
	"github.com/influxdata/influxdb/rpc"
	"github.com/influxdata/influxdb/tsdb"
	"net"
)

type Server struct {
	name   string
	dir    string
	store  *tsdb.Store
	rpcSvc *rpc.Service

	ln net.Listener
}

func New(c *Config) *Server {
	s := &Server{}
	s.store = tsdb.NewStore(c.dir)
	return s
}
func (s *Server) Open() error {
	s.store.Open()
	return nil
}
func (s *Server) Close() error {
	return s.store.Close()
}
