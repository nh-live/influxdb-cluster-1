package raft_store

import (
	"github.com/hashicorp/raft"
	"github.com/influxdata/influxdb/models/meta"
	"github.com/influxdata/influxdb/platform/infserver"
	"io"
	"net"
	"sync"
)

type RaftStore struct {
	mu        sync.RWMutex
	data      *meta.Data
	raftState *raftState
}

func NewRaftStore(localID string, nodes map[string]infserver.Node, path string, ln net.Listener) *RaftStore {
	return &RaftStore{
		raftState: NewRaftState(localID, nodes, path, ln),
	}
}
func (rs *RaftStore) Open() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.raftState.open(rs)
}
func (rs *RaftStore) Apply(log *raft.Log) interface{} {
	return nil
}

func (rs *RaftStore) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
func (rs *RaftStore) Restore(ir io.ReadCloser) error {
	return nil
}
