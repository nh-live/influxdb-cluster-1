package infserver

import (
	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/platform/infserver/metaservice"
	"github.com/influxdata/influxdb/platform/netserver"
)

type Config struct {
	BindAddress string                       `toml:"bing-address"`
	Name        string                       `toml:"name"`
	Nodes       map[string]influxdb.MetaNode `toml:"nodes"`
	Meta        *metaservice.Config          `toml:"meta"`
	Httpd       *netserver.Config            `toml:"http"`
}
