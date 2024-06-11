package controller

//chamar primeiro GetGymController para ver se está no raio de alguma academia
//Se tiver no raio, retornar as academias desse raio para o usuário escolher a que gostaria de fazer check-in
//se não, s

import (
	"goapirest/internal/usecase"
	"net/http"
)

type GymController struct {
	gymUseCase *usecase.GymUseCase
}

func NewGymController(u *usecase.GymUseCase) *GymController {
	return &GymController{gymUseCase: u}
}

func (c *GymController) CreateGymController(w http.ResponseWriter, r *http.Request) {

	// var gymDto dto.GymInputDTO
	// c.gymUseCase.Execute()
}
