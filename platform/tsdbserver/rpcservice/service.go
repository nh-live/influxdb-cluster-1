package rpcservice

import (
	"context"
	"github.com/influxdata/influxdb/platform/tsdbserver/internal"
	"google.golang.org/grpc"
	"net"
)

type RpcService struct {
	addr string
	ln   net.Listener

	TSDBStore interface {
	}
}

func New() *RpcService {
	return &RpcService{}
}
func (rs *RpcService) Open() error {
	ln, err := net.Listen("tcp", rs.addr)
	grpcServer := grpc.NewServer()
	internal.RegisterTsdbRPCServer(grpcServer, rs)
	go func() {
		grpcServer.Serve(ln)
	}()
	return nil
}

func (rs *RpcService) WritePoints(ctx context.Context, request *internal.WritePointsRequest) (*internal.WritePointsReply, error) {

}
