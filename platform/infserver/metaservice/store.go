package metaservice

import (
	"github.com/influxdata/influxdb/models/meta"
	"github.com/influxdata/influxdb/platform/infserver/metaservice/raft_store"
	"sync"
)

type Store struct {
	mu        sync.RWMutex
	cacheData *meta.Data

	raftStore *raft_store.RaftStore
	change    chan interface{}
	close     chan interface{}
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Open() error {
}
