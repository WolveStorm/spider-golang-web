package sys_init

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"spider-golang-web/grpc_server"
)

var GRPCPool *grpc_server.Pool

func InitGrpcPool() {
	var err error
	GRPCPool, err = grpc_server.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial(":8770", grpc.WithInsecure())
	}, 1, 10, 0)
	if err != nil {
		zap.S().Errorf("error to init grpc pool,detail:%v", err)
		return
	}
}
