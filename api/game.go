package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"spider-golang-web/proto"
	"spider-golang-web/sys_init"
	"spider-golang-web/util"
)

func GameList(c *gin.Context) {
	var req struct {
		Page     int32  `json:"page"`
		PageSize int32  `json:"page_size"`
		Keyword  string `json:"keyword"`
	}
	if err := c.ShouldBind(&req); err != nil {
		util.RspBindingError(c, err)
		return
	}
	dial, err := sys_init.GRPCPool.Get(context.Background())
	defer dial.Close()
	client := proto.NewGameClient(dial)
	// 从数据库拿信息
	list, err := client.GameList(context.Background(), &proto.GameListFilterRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		zap.S().Errorf("error to call gamelist,args:%v", req)
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, list)
}

func GameDetail(c *gin.Context) {
	var req struct {
		GameName string `json:"game_name"`
	}
	if err := c.ShouldBind(&req); err != nil {
		util.RspBindingError(c, err)
		return
	}
	dial, err := sys_init.GRPCPool.Get(context.Background())
	defer dial.Close()
	client := proto.NewGameClient(dial)
	detail, err := client.GameDetail(context.Background(), &proto.GameDetailRequest{
		GameName: req.GameName,
	})
	if err != nil {
		zap.S().Errorf("error to call gamelist,args:%v", req)
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, detail)
}
