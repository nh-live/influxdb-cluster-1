package transervice

import "github.com/influxdata/influxdb/models"

type TranService struct {
}

func New() *TranService {
	return &TranService{}
}

func (ts *TranService) Open() error {
	return nil
}

func (ts *TranService) WriteToDataNode(dnID uint64, pts models.Points) error {
	return nil
}
