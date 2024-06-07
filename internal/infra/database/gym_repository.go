package database

import (
	"database/sql"
	"fmt"
	"goapirest/internal/entity"
	"log"
)

type GymRepository struct {
	Db *sql.DB
}

func NewGymRepository(db *sql.DB) *GymRepository {
	return &GymRepository{Db: db}
}

func (r *GymRepository) Save(gym *entity.Gym) error {

	pstmt, err := r.Db.Prepare("insertUser", "INSERT INTO gyms (name, address, placeID, longitude, latitude) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
        log.Fatalf("Prepare failed: %v\n", err)
    }

	_, err = r.Db.Exec(pstmt.SQL, gym.Name, gym.Address, gym.PlaceID, gym.Longitude, gym.Latitude)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}
	return nil
}
