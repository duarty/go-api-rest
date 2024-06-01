package api

type GoogleMapsClient interface {
    GetGymsFromCoordinates(lat, lng float64) (string, error)
}

type googleMapsClient struct {
    apiKey string
}

func NewGoogleMapsClient(apiKey string) *googleMapsClient {
    return &googleMapsClient{apiKey: apiKey}
}
