package service

import (
	"gvozdevLabApi/internal/models"
	"net/url"
)

func (s *service) GetPreview(id string) (*models.Preview, error) {
	preview := &models.Preview{}
	err := s.gvozdevLabRepo.GetPreview(preview, id)
	if err != nil {
		return nil, err
	}
	return preview, nil
}

func (s *service) AddPreview(previewUrl string) error {
	_, err := url.ParseRequestURI(previewUrl)
	if err != nil {
		return err
	}
	err = s.gvozdevLabRepo.AddPreview(previewUrl)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeletePreview(id string) error {
	err := s.gvozdevLabRepo.DeletePreview(id)
	if err != nil {
		return err
	}
	return nil
}
