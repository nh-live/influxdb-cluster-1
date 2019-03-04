package infserver

import "github.com/influxdata/influxdb/platform/infserver/metaservice"

type Config struct {
	NodeName string              `toml:"name"`
	Meta     *metaservice.Config `toml:"meta"`
}
