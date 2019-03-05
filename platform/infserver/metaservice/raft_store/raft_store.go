package raft_store

import (
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
)

type RaftStore struct {
	c         *Config
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

func NewRaftStore(c *Config, fsm raft.FSM, ln net.Listener) *RaftStore {
	return &RaftStore{
		LocalID:   raft.ServerID(c.LocalID),
		LocalAddr: raft.ServerAddress(c.LocalAddr),
		LocalPath: c.LocalPath,
		fsm:       fsm,
		ln:        ln,
	}
}
func (rs *RaftStore) Open(ln net.Listener) error {
	rs.ln = ln
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
	rs.raft.BootstrapCluster()
	return nil
}
