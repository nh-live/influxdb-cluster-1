package transervice

import (
	"github.com/influxdata/influxdb/models"
	pb "github.com/influxdata/influxdb/rpc/tsdb"
	"google.golang.org/grpc"
)

type TranService struct {
}

func New() *TranService {
	return &TranService{}
}

func (ts *TranService) Open() error {
	return nil
}

func (ts *TranService) WriteToDataNode(dnID uint64, sID uint64, pts models.Points) error {
	conn := grpc.Dial()
	pb.NewTsdbRPCClient(conn).WritePoints()
	return nil
}
