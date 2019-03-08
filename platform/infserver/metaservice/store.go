package metaservice

import (
	"github.com/hashicorp/raft"
	"github.com/influxdata/influxdb/models/meta"
	"github.com/influxdata/influxdb/platform/infserver"
	"github.com/influxdata/influxdb/platform/infserver/metaservice/raft_store"
	"io"
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
	s.raftStore = raft_store.NewRaftStore(localID, nodes, path, ln)
	return s
}

func (s *Store) Open() error {
	return s.raftStore.Open()
}

func (s *Store) Apply(log *raft.Log) interface{} {
	return nil
}
func (s *Store) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
func (s *Store) Restore(ir io.ReadCloser) error {
	return nil
}
