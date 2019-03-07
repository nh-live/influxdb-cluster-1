package infserver

type Node struct {
	Name        string `toml:"name"`
	RaftAddress string `toml:"raft-bind-address"`
}
