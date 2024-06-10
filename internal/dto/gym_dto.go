package dto

type GymInputDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	PlaceID   string  `json:"placeId"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GymRequest struct {
	IncludedTypes       []string             `json:"includedTypes"`
	MaxResultCount      uint8                `json:"maxResultCount"`
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
	DisplayName      *DisplayName `json:"displayName"`
	PlaceID          string       `json:"id"`
	FormattedAddress string       `json:"formattedAddress"`
	Location         *Location    `json:"location"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type DisplayName struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}
