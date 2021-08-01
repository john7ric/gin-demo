package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/john7ric/gin-demo/entity"
	"github.com/john7ric/gin-demo/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *videoController) Save(c *gin.Context) entity.Video {
	var video entity.Video
	c.BindJSON(&video)
	vc.service.Save(video)
	return video
}
