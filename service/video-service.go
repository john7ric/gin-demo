package service

import "github.com/john7ric/gin-demo/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// confirms to interface VideoService
type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

//save and returns a new video entity
func (vs *videoService) Save(v entity.Video) entity.Video {
	vs.videos = append(vs.videos, v)
	return v
}

// returns the existing videos
func (vs *videoService) FindAll() []entity.Video {
	return vs.videos
}
