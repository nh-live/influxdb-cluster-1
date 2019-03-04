package netserver

import (
	"net"
	"net/http"
)

type Service struct {
	ln      net.Listener
	handler *Handler
	addr    string
}

func New(c *Config) *Service {
	s := &Service{
		addr: c.Addr,
	}
}

func (s *Service) Open() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.ln = ln

	s.serve()
}

func (s *Service) serve() {
	err := http.Serve(s.ln, s.handler)
}
