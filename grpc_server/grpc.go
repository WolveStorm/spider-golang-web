package grpc_server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"spider-golang-web/proto"
)

func InitGrpcServer() {
	server := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8770")
	if err != nil {
		zap.S().Error("error to listen tcp on 8770")
		return
	}
	InitGameGrpcService(server)
	if err = server.Serve(listen); err != nil {
		zap.S().Error("error to start grpc server")
		return
	}
	zap.S().Info("grpc server start at port 8770")
}

func InitGameGrpcService(server *grpc.Server) {
	proto.RegisterGameServer(server, &proto.ServiceGameImpl{})

}
