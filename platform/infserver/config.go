package infserver

import (
	"github.com/influxdata/influxdb/platform/infserver/metaservice"
	"github.com/influxdata/influxdb/platform/netserver"
)

type Config struct {
	Name  string              `toml:"name"`
	Path  string              `toml:"dir"`
	Nodes map[string]Node     `toml:"nodes"`
	Meta  *metaservice.Config `toml:"meta"`
	Httpd *netserver.Config   `toml:"http"`
}
