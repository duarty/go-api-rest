package api

type GoogleMapsClient interface {
    GetGymsFromCoordinates(lat, lng float64) (string, error)
}
