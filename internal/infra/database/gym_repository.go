package database

import (
	"fmt"
	"goapirest/internal/entity"
	"log"

	"github.com/jackc/pgx"
)

type GymRepository struct {
	Db *pgx.Conn
}

func NewGymRepository(db *pgx.Conn) *GymRepository {
	return &GymRepository{Db: db}
}

func (r *GymRepository) GetSavedGyms(gym *entity.GymRequest) []entity.Gym {
	//puxar academias no raio passando longitude e latitude
}

func (r *GymRepository) SaveGym(gym *entity.Gym) error {
	pstmt, err := r.Db.Prepare("saveGym", "INSERT INTO gyms (name, address, placeID, longitude, latitude) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatalf("Prepare failed: %v\n", err)
	}
	_, err = r.Db.Exec(pstmt.SQL, gym.Name, gym.Address, gym.PlaceID, gym.Longitude, gym.Latitude)
	if err != nil {
		return fmt.Errorf("unable to insert gym: %w", err)
	}
	return nil
}
