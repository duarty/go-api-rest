package controller

import (
	"goapirest/internal/usecases"
	"net/http"
)


type GymController struct {
	gymUseCase usecases.GymUseCase
}

func NewGymController (u usecases.GymUseCase) *GymController {
	return	&GymController{
		gymUseCase: u,
	}
}

func (c *GymController) createGym(w http.ResponseWriter, r http.Request){

}
