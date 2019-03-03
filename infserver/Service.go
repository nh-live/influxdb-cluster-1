package infserver

import (
	"github.com/influxdata/influxdb/dp/point"
	"github.com/influxdata/influxdb/dp/result"
)

type Service struct {
	MetaClient interface {
		mapShard(points point.Points) (map[uint64]point.Points, error) //map[shardID]Points
	}
}

func (s *Service) writeToShard() {

}
func (s *Service) Write(pts point.Points) error {
	spts, err := s.MetaClient.mapShard(pts)
	for sid, pts := range spts {
		s.writeToShard(sid, pts)
	}
}

func (s *Service) Query(qa string) (result.Result, error) {

}
