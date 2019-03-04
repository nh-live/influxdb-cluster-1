package metaservice

import (
	"github.com/influxdata/influxdb/models/meta"
	"sync"
)

type Store struct {
	mu        sync.RWMutex
	cacheData *meta.Data

	change chan interface{}
	close  chan interface{}
}
