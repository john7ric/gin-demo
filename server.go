package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/john7ric/gin-demo/controller"
	"github.com/john7ric/gin-demo/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/save", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.Save(c))
	})

	server.Run(":8080")
}
