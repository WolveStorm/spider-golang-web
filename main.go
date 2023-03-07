package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"spider-golang-web/grpc_server"
	"spider-golang-web/middleware"
	"spider-golang-web/router"
	"spider-golang-web/sys_init"
	"syscall"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.Cors)
	router.InitGameStoreRouter(engine)
	sys_init.InitMongo()
	sys_init.InitPostgres()
	sys_init.InitAllZap()
	sys_init.InitRedis()
	sys_init.InitGrpcPool()
	go func() {
		grpc_server.InitGrpcServer()
	}()
	zap.S().Info("gin server start at port 8660")
	go func() {
		if err := engine.Run("0.0.0.0:8660"); err != nil {
			zap.S().Error("error to start gin server")
			return
		}
	}()
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGQUIT)
	<-sign
}
