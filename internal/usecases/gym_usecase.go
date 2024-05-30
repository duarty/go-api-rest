package usecases

import (
	"goapirest/internal/domain/model"
)

type GymService interface {
	CreateGym(model.Gym) error
}

type GymUseCase struct {
	service GymService
}

func NewGymUsecase(s GymService) *GymUseCase {
    return &GymUseCase{
        service: s,
    }
}

func (u *GymUseCase) CreateGym(gym model.Gym)	error {
	return u.service.CreateGym(gym)
}
