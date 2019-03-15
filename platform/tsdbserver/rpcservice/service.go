package rpcservice

import (
	"context"
	"github.com/influxdata/influxdb/models"
	pb "github.com/influxdata/influxdb/rpc/tsdb"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type RpcService struct {
	mu     sync.RWMutex
	addr   string
	ln     net.Listener
	closed chan struct{}

	TSDBStore interface {
		WriteToShard(shardID uint64, points []models.Point) error
	}
}

func New(c *Config) *RpcService {
	return &RpcService{
		addr:   c.BindAddress,
		closed: make(chan struct{}),
	}
}
func (rs *RpcService) Open() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	ln, err := net.Listen("tcp", rs.addr)
	if err != nil {
		return err
	}

	gp := grpc.NewServer()
	pb.RegisterTsdbRPCServer(gp, rs)
	go gp.Serve(ln)
	return nil
}

func (rs *RpcService) WritePoints(ctx context.Context, request *pb.WritePointsRequest) (*pb.WritePointsReply, error) {
	if err := func() error {
		pts, err := models.ParsePoints(request.Points)
		if err != nil {
			return err
		}
		if err := rs.TSDBStore.WriteToShard(request.ShardID, pts); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return &pb.WritePointsReply{Err: err.Error()}, err
	}
	return &pb.WritePointsReply{Err: ""}, nil
}
