package service

import (
	"goapirest/internal/domain/model"
	"goapirest/internal/repository"
)

type GymService struct {
	repo repository.GymRepository
}

func NewGymService(repo repository.GymRepository) *GymService {
	return &GymService{repo: repo}
}

func (s *GymService) CreateGym(gym *model.Gym) error {
	return s.repo.CreateGym(gym)
}
