package database

import (
	"gvozdevLabApi/internal/database/query"
	"gvozdevLabApi/internal/models"
)

func (r RepoDB) GetVideo(preview *models.Video, id string) error {
	err := r.db.Get(preview, query.GetVideo, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) AddVideo(url string) error {
	_, err := r.db.Query(query.AddVideo, url)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) DeleteVideo(id string) error {
	_, err := r.db.Query(query.DeleteVideo, id)
	if err != nil {
		return err
	}
	return nil
}
