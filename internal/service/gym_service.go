package service

import (
	"goapirest/internal/model"
	"goapirest/internal/repository"
)

type GymService struct {
    repo repository.GymRepository
}

func NewGymService(repo repository.GymRepository) *GymService {
    return &GymService{repo: repo}
}

func (s *GymService) CreateGym(gymData *model.GymData) error {
    return s.repo.CreateGym(gymData)
}
