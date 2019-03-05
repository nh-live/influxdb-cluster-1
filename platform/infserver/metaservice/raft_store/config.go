package raft_store

type Config struct {
	LocalID   uint64
	LocalAddr string
	LocalPath string

	RaftServers []string

	Exist    bool
	JoinAddr string
}
