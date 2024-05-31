package usecases

import (
	"goapirest/internal/domain/model"
)

type GymService interface {
	CreateGym(gym *model.Gym) error
}

type GymUseCase struct {
	service GymService
}

func NewGymUseCase(service GymService) *GymUseCase {
	return &GymUseCase{
		service: service,
	}
}

func (u *GymUseCase) CreateGym(gym *model.Gym) error {
	return u.service.CreateGym(gym)
}
