package infserver

import (
	"github.com/influxdata/influxdb/platform/infserver/metaservice"
	"github.com/influxdata/influxdb/platform/netserver"
)

type Config struct {
	NodeName string             `toml:"name"`
	Meta     metaservice.Config `toml:"meta"`
	HTTPD    netserver.Config   `toml:meta"`
}
