package main

import (
	"fmt"
	"goapirest/configs"
	"goapirest/internal/app"
	"goapirest/internal/infra/database"
	"log"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	_, err := configs.LoadEnv(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	db, err := database.DBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	app.Server()

	// defer db.Close()
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

	// URL := config.GoogleMapsApiEndPoint + config.GoogleMapsApiPlaces
	// req, err := http.NewRequest("POST", URL, body)
	// if err != nil {
	// 	panic(err)
	// }

	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-Goog-Api-Key", config.GoogleMapsApiKey)
	// req.Header.Set("X-Goog-FieldMask", config.SearchFieldMask)

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response:", err)
	// 	return
	// }

	// var gymResponse *model.GymResponse
	// err = json.Unmarshal(responseBody, &gymResponse)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling response:", err)
	// 	return
	// }

	// var name *string
	// var address *string
	// var placeId *string
	// var longitude *float64
	// var latitude *float64

	// for _,v := range *gymResponse.Places {
	// 	name = &v.DisplayName.Text
	// 	address = &v.FormattedAddress
	// 	placeId = &v.PlaceID
	// 	longitude = &v.Location.Longitude
	// 	latitude = &v.Location.Latitude
	// }

	// newGym := CreateGym(*name, *address, *placeId, *longitude, *latitude)

	// err = insertGym(db, newGym)
	// if err != nil {
	// 	panic(err)
	// }

}

// func insertGym ( db *pgx.Conn, gym *model.GymData) error {
// 	pstmt, err := db.Prepare("insertUser", "INSERT INTO gyms (id, name, address, placeID, longitude, latitude) VALUES ($1, $2, $3, $4, $5, $6)")
// 	if err != nil {
//         log.Fatalf("Prepare failed: %v\n", err)
//     }

// 	_, err = db.Exec(pstmt.SQL, gym.ID, gym.Name, gym.Address, gym.PlaceID, gym.Longitude, gym.Latitude)
// 	if err != nil {
// 		return fmt.Errorf("unable to insert user: %w", err)
// 	}
// 	return nil
// }

// func NewGym (name string, address string, placeId string, longitude float64, latitude float64 ) *model.GymData {
// 	return &model.GymData{
// 		ID: 1,
// 		Name: name,
// 		Address: address,
// 		PlaceID: placeId,
// 		Longitude: longitude,
// 		Latitude: latitude,
// 	}
// }
