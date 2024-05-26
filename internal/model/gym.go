package model


type GymResponse struct {
	Places *[]Place `json:"places"`
}

type Place struct {
	DisplayName *DisplayName `json:"displayName"`
	PlaceID string `json:"id"`
	FormattedAddress string `json:"formattedAddress"`
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

type GymData struct {
	ID uint
	Name string
	Address string
	PlaceID string
	Longitude float64
	Latitude float64
}

