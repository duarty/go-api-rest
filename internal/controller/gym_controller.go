package controller

import (
	"goapirest/internal/dto"
	"goapirest/internal/usecases"
	"net/http"
)


type GymController struct {
	gymUseCase usecases.GymUsecase
}

func NewGymController (u usecases.GymUsecase) *GymController {
	return	&GymController{
		gymUseCase: u,
	}
}

func (c *GymController) createGym(w http.ResponseWriter, r http.Request){
	var gymDTO *dto.GymDTO
}
