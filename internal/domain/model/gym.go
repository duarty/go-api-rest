package model



type GymRequest struct {
	IncludedTypes       []string
	MaxResultCount      uint8
	LocationRestriction *LocationRestriction
}

type LocationRestriction struct {
	Circle *Circle
}

type Circle struct {
	Center *Center
	Radius float32
}

type Center struct {
	Longitude float64
	Latitude  float64
}


type GymResponse struct {
	Places *[]Place
}

type Place struct {
	DisplayName *DisplayName
	PlaceID string
	FormattedAddress string
	Location *Location
}

type Location struct {
	Longitude float64
	Latitude  float64
}

type DisplayName struct {
	Text         string
	LanguageCode string
}

type Gym struct {
	ID uint
	Name string
	Address string
	PlaceID string
	Longitude float64
	Latitude float64
}

func NewGym (name, address, placeId string, longitude, latitude float64 ) *Gym {
	return &Gym{
		ID: 1,
		Name: name,
		Address: address,
		PlaceID: placeId,
		Longitude: longitude,
		Latitude: latitude,
	}
}


