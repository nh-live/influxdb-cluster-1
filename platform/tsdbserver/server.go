package tsdbserver

import "github.com/influxdata/influxdb/tsdb"

type Server struct {
	ServerID uint64
	store    *tsdb.Store
}

func New() *Server {
	s := &Server{}
	s.store = tsdb.NewStore()
	tsdb.NewStore()
	return &Server{}
}
func (s *Server) Open() error {
	s.store.Open()
	return nil
}
func (s *Server) Close() error {
	return nil
}
