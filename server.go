package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goldenmonster/golang-gin-poc/controller"
	"github.com/goldenmonster/golang-gin-poc/service"
)

var (
	videoService service.VideoService = service.New()
	videoControoler controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func (ctx *gin.Context)  {
		ctx.JSON(200, videoControoler.FindAll) 
	})

	server.POST("/videos", func (ctx *gin.Context)  {
		ctx.JSON(200, videoControoler.Save(ctx)) 
	})
	

	server.Run(":8080")
}