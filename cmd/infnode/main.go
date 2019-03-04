package main

import (
	"fmt"
	"github.com/influxdata/influxdb/platform/infserver"
	"os"
)

const (
	SERVICENAME = "infnode"
)

func main() {
	s := infserver.New()
	if err := s.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "error open %s", SERVICENAME)
	}
}
