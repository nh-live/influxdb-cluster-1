package tsdbserver

import "github.com/influxdata/influxdb/tsdb"

type Server struct {
	ServerID uint64
	store    *tsdb.Store
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
