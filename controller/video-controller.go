package controller

import "main.go/entity"

type VideoController interface {
	FindAll() []entity.Video
	
}