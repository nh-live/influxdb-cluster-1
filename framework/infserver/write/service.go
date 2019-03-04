package write

import "github.com/influxdata/influxdb/dp/point"

type WriteService struct {
	MetaClient interface {
		MapShard(points point.Points) (map[uint64]point.Points, error) //map[shardID]Points
		GetDnsIDBySID(shardID uint64) (dnsID []uint64, err error)
	}
	Transmitter interface {
		WriteToDataNode(dnID uint64, pts point.Points) error
	}
}

func (s *WriteService) writeToShard(sid uint64, pts point.Points) error {
	dnsID, err := s.MetaClient.GetDnsIDBySID(sid)
	for _, dn := range dnsID {
		//TODO 高可用考虑是否重写，以及是否可以重写，若不可重写，如何保证一致性
		err := s.Transmitter.WriteToDataNode(dn, pts)
	}

	return err
}

func (s *WriteService) Write(pts point.Points) error {
	spts, err := s.MetaClient.MapShard(pts)
	for sid, pts := range spts {
		s.writeToShard(sid, pts)
	}
	return err
}
