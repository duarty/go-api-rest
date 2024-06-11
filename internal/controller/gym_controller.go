package controller

//chamar primeiro GetGymController para ver se está no raio de alguma academia
//Se tiver no raio, retornar as academias desse raio para o usuário escolher a que gostaria de fazer check-in
//se não, s

import (
	"encoding/json"
	"fmt"
	"goapirest/internal/dto"
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
		var format dto.CheckInDTO
		err = json.Unmarshal(body, &format)
		if err != nil {
			panic(err)
		}

		fmt.Print(format.Latitude)
	}

	// if r.Method == "GET" {
	// 	io.Writer()
	// }
	// var gymDto dto.GymInputDTO
	// c.gymUseCase.Execute()
}
