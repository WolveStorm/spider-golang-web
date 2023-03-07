package router

import (
	"github.com/gin-gonic/gin"
	"spider-golang-web/api"
)

func InitGameStoreRouter(engine *gin.Engine) {
	// test
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "ok",
		})
		return
	})
	group := engine.Group("/game_store")
	{
		group.POST("/game_list", api.GameList)
		group.POST("/game_detail", api.GameDetail)
	}
}
