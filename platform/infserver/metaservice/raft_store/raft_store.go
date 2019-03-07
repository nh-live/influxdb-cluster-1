package raft_store

import (
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/influxdata/influxdb/platform/infserver"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
)

type RaftStore struct {
	fsm       raft.FSM
	LocalID   raft.ServerID
	LocalAddr raft.ServerAddress
	LocalPath string

	peers     []raft.ServerAddress
	ln        net.Listener
	transport *raft.NetworkTransport
	rstore    *raftboltdb.BoltStore
	raft      *raft.Raft
}

func NewRaftStore(localID string, nodes map[string]infserver.Node, path string, fsm raft.FSM, ln net.Listener) *RaftStore {
	lnode := nodes[localID]
	return &RaftStore{
		LocalID:   raft.ServerID(lnode.Name),
		LocalAddr: raft.ServerAddress(lnode.RaftAddress),
		LocalPath: path,
		fsm:       fsm,
		ln:        ln,
	}
}
func (rs *RaftStore) Open() error {
	config := raft.DefaultConfig()
	config.LogOutput = ioutil.Discard

	raftlayer := newRaftLayer(rs.LocalAddr, rs.ln)

	rs.transport = raft.NewNetworkTransport(raftlayer, 3, 10*time.Second, config.LogOutput)

	store, err := raftboltdb.NewBoltStore(filepath.Join(rs.LocalPath, "raft.db"))
	if err != nil {
		return fmt.Errorf("new bolt store : %s", err)
	}
	rs.rstore = store

	snapshots, err := raft.NewFileSnapshotStore(rs.LocalPath, 2, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}

	ra, err := raft.NewRaft(config, rs.fsm, store, store, snapshots, rs.transport)
	if err != nil {
		return fmt.Errorf("new raft : %s", err)
	}
	rs.raft = ra
	return nil
}
