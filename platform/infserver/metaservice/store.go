package metaservice

import (
	"github.com/hashicorp/raft"
	"github.com/influxdata/influxdb/models/meta"
	"github.com/influxdata/influxdb/platform/infserver"
	"github.com/influxdata/influxdb/platform/infserver/metaservice/raft_store"
	"net"
	"sync"
)

type Store struct {
	mu        sync.RWMutex
	cacheData *meta.Data

	raftStore *raft_store.RaftStore
	changed   chan interface{}

	closed chan interface{}
}

func NewStore(localID string, nodes map[string]infserver.Node, path string, ln net.Listener) *Store {
	s := &Store{}
	s.raftStore = raft_store.NewRaftStore(localID, nodes, path, raft.FSM(s), ln)
	return s
}

func (s *Store) Open() error {
	return s.raftStore.Open()
}
