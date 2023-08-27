package service

import (
	"gvozdevLabApi/internal/database"
	"gvozdevLabApi/internal/models"
)

type service struct {
	gvozdevLabRepo database.Storage
}

type Service interface {
	GetVideo(id string) (*models.Video, error)
	AddVideo(url string) error
	DeleteVideo(id string) error

	GetPreview(id string) (*models.Preview, error)
	AddPreview(url string) error
	DeletePreview(id string) error
}

func NewService(gvozdevLabRepo database.Storage) Service {
	return &service{gvozdevLabRepo}
}
