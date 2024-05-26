package configs

import (
	"fmt"

	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var config *configs

type configs struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        uint16 `mapstructure:"DB_PORT"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBUser        string `mapstructure:"DB_USER"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"SERVER_PORT"`
	JWTSecretKey  string `mapstructure:"SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_TIME"`
	TokenAuthKey  *jwtauth.JWTAuth
	GoogleMapsApiEndPoint string `mapstructure:"GOOGLE_MAPS_API_ENDPOINT"`
	GoogleMapsApiPlaces string `mapstructure:"GOOGLE_MAPS_API_PLACES"`
	GoogleMapsApiKey string `mapstructure:"GOOGLE_MAPS_API_KEY"`
	NearbySearchRadius float32 `mapstructure:"NEARBY_SEARCH_RADIUS"`
	IncludedTypes string `mapstructure:"INCLUDED_TYPES"`
	SearchFieldMask string `mapstructure:"SEARCH_FIELDMASK"`
}

func init() {
	LoadEnv("../.env")
}

func LoadEnv(path string) (*configs, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config.TokenAuthKey = jwtauth.New("HS256", []byte(config.JWTSecretKey), nil)
	return config, nil
}
