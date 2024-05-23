package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goapirest/configs"
	"io"
	"log"
	"net/http"
)

type GymRequest struct {
	IncludedTypes       []string             `json:"includedTypes"`
	MaxResultCount      uint8                  `json:"maxResultCount"`
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
	Latitude  float64 `json:"latitude"`
}



type GymResponse struct {
	Places *[]Place `json:"places"`
}

type Place struct {
	DisplayName *DisplayName `json:"displayName"`
	PlaceID *string `json:"id"`
	FormattedAddress *string `json:"formattedAddress"`
	Location *Location `json:"location"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type DisplayName struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

func main() {
	config, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// data := &GymRequest{
	// 	IncludedTypes:  []string{"gym"},
	// 	MaxResultCount: 1,
	// 	LocationRestriction: &LocationRestriction{
	// 		Circle: &Circle{
	// 			Center: &Center{
	// 				Longitude: 2.814704860187976,
	// 				Latitude:  -60.72387482874884,
	// 			},
	// 			Radius: 1000.0,
	// 		},
	// 	},
	// }
	data := map[string]interface{}{
		"includedTypes": []string{"gym"},
		"maxResultCount": 1,
		"locationRestriction": map[string]interface{}{
			"circle": map[string]interface{}{
				"center": map[string]float64{
					"latitude": 2.8135529076686567,
					"longitude": -60.67046458758208,
				},
				"radius": 5,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Erro codificando dados JSON: %v", err)
	}

	body := bytes.NewBuffer(jsonData)

	URL := config.GoogleMapsApiEndPoint + config.GoogleMapsApiPlaces
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", config.GoogleMapsApiKey)
	req.Header.Set("X-Goog-FieldMask", config.SearchFieldMask)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var gymResponse GymResponse
	err = json.Unmarshal(responseBody, &gymResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	for _,v := range *gymResponse.Places {
		fmt.Printf("Deserialized Response: %+v\n", *v.Location)
	}



}


