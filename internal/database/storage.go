package database

import "gvozdevLabApi/internal/models"

type Storage interface {
	GetVideo(video *models.Video, id string) error
	AddVideo(url string) error
	DeleteVideo(id string) error

	GetPreview(preview *models.Preview, id string) error
	AddPreview(url string) error
	DeletePreview(id string) error
}
