package usecases

import (
	"goapirest/internal/domain/model"
)

type GymService interface {
    CreateGym(gym *model.Gym) error
}

type GymUsecase struct {
    gymService GymService
}

func NewGymUsecase(s GymService) *GymUsecase {
    return &GymUsecase{
        gymService: s,
    }
}

func (u *GymUsecase) CreateGym(gym *model.Gym) error {
    return u.gymService.CreateGym(gym)
}
