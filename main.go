package main

import (
	"encoding/json"
	"fmt"
	"goapirest/configs"
	"io"
	"net/http"
	"strings"
)

type GymRequest struct{
	IncludedTypes []string `json:"includedTypes"`
	MaxResultCount  int `json:"maxResultCount"`
	LocationRestriction *LocationRestriction `json:"locationRestriction"`
}

type LocationRestriction struct {
	Circle *Circle `json:"circle"`
}

type Circle struct {
	Center *Center `json:"center"`
	Radius float32 `json:"radius"`
}

type Center struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

func main() {
	_, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var gymRequestData = &GymRequest{
		IncludedTypes: []string{"gym"},
		MaxResultCount: 1,
		LocationRestriction: &LocationRestriction{
			Circle: &Circle{
				Center: &Center{
					Longitude: 2.814704860187976,
					Latitude: -60.72387482874884,
				},
				Radius: 50.0,
			},
		},
	}

	gymRequestDataJson, err := json.Marshal(gymRequestData)
	if err != nil {
		panic(err)
	}
	body := strings.NewReader(string(gymRequestDataJson))
	req, err := http.Post("https://places.googleapis.com/v1/places:searchNearby", "application/json",body)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	responseBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response Body:", string(responseBody))

}
