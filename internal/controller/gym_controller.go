package controller

//chamar primeiro GetGymController para ver se está no raio de alguma academia
//Se tiver no raio, retornar as academias desse raio para o usuário escolher a que gostaria de fazer check-in
//se não, s

import (
	"encoding/json"
	"goapirest/internal/entity"
	"goapirest/internal/usecase"
	"io"
	"net/http"
)

type GymController struct {
	gymUseCase *usecase.GymUseCase
}

func NewGymController(u *usecase.GymUseCase) *GymController {
	return &GymController{gymUseCase: u}
}

func (c *GymController) CreateGymController(w http.ResponseWriter, r *http.Request) {
	// data := map[string]interface{}{
	// 	"includedTypes": []string{config.IncludedTypes},
	// 	"maxResultCount": 1,
	// 	"locationRestriction": map[string]interface{}{
	// 		"circle": map[string]interface{}{
	// 			"center": map[string]float64{
	// 				"latitude": 2.8135529076686567,
	// 				"longitude": -60.67046458758208,
	// 			},
	// 			"radius": config.NearbySearchRadius,
	// 		},
	// 	},
	// }

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatalf("Erro codificando dados JSON: %v", err)
	// }

	// body := bytes.NewBuffer(jsonData)
	if r.Method == "POST" {

		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var checkInCoordinates *entity.CheckInCoordinates
		err = json.Unmarshal(body, &checkInCoordinates)
		if err != nil {
			panic(err)
		}
		c.gymUseCase.Execute(checkInCoordinates)
	}

	// if r.Method == "GET" {
	// 	io.Writer()
	// }
	// var gymDto dto.GymInputDTO
	// c.gymUseCase.Execute()
}
