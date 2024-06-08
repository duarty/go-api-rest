package controller

import (
	"goapirest/internal/usecase"
	"net/http"
)

type GymController struct {
    userUseCase *usecase.CreateGymUseCase
}

func NewGymController(u *usecase.CreateGymUseCase) *GymController {
	return &GymController{userUseCase: u}
}

func (c *GymController) Create(r http.Request, w http.ResponseWriter){

}
