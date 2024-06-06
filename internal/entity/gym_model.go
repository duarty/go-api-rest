package entity

import "errors"

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
	//ID uint
	Name string
	Address string
	PlaceID string
	Longitude float64
	Latitude float64
}

func NewGym (name, address, placeId string, longitude, latitude float64 ) (*Gym, error) {
	gym := &Gym{
		//ID: 1,
		Name: name,
		Address: address,
		PlaceID: placeId,
		Longitude: longitude,
		Latitude: latitude,
	}

	err := gym.isValid()
	if err != nil {
		return nil, err
	}
	return gym, nil
}

func (g *Gym) isValid() error {
	if g.Address == "" {
		return errors.New("Address is invalid")
	}
	if g.Name == "" {
		return errors.New("Name is invalid")
	}
	if g.PlaceID == "" {
		return errors.New("PlaceID is invalid")
	}
	if g.Latitude < -90.0 && g.Latitude > 90.0 {
		return errors.New("Latitude is invalid")
	}
	if g.Longitude < -180.0 && g.Longitude > 180.0  {
		return errors.New("Longitude is invalid")
	}
	return nil
}


