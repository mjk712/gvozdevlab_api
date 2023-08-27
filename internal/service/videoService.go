package service

import (
	"gvozdevLabApi/internal/models"
	"net/url"
)

func (s *service) GetVideo(id string) (*models.Video, error) {
	video := &models.Video{}
	err := s.gvozdevLabRepo.GetVideo(video, id)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (s *service) AddVideo(videoUrl string) error {
	_, err := url.ParseRequestURI(videoUrl)
	if err != nil {
		return err
	}
	err = s.gvozdevLabRepo.AddVideo(videoUrl)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteVideo(id string) error {
	err := s.gvozdevLabRepo.DeleteVideo(id)
	if err != nil {
		return err
	}
	return nil
}
