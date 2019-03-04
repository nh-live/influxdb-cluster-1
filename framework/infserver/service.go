package infserver

import (
	"github.com/influxdata/influxdb/dp/point"
	"github.com/influxdata/influxdb/dp/result"
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/query"
)

type Service struct {
	WriteService interface {
		Write(pts models.Points) error
	}

	QueryService interface {
		Query(qa string) (rst query.Result, err error)
	}
}

func (s *Service) Write(pts models.Points) error {
	return s.WriteService.Write(pts)
}

func (s *Service) Query(qa string) (models.Points, error) {

}
