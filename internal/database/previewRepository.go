package database

import (
	"gvozdevLabApi/internal/database/query"
	"gvozdevLabApi/internal/models"
)

func (r RepoDB) GetPreview(preview *models.Preview, id string) error {
	err := r.db.Get(preview, query.GetPreview, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) AddPreview(url string) error {
	_, err := r.db.Query(query.AddPreview, url)
	if err != nil {
		return err
	}
	return nil
}

func (r RepoDB) DeletePreview(id string) error {
	_, err := r.db.Query(query.DeletePreview, id)
	if err != nil {
		return err
	}
	return nil
}
