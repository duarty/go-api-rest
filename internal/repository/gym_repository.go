package repository

import (
	"fmt"
	"goapirest/internal/domain/model"
	"log"

	"github.com/jackc/pgx"
)

// type GymRepository interface {
//     CreateGym(gym *model.Gym) error
// }

type GymRepo struct {
    db *pgx.Conn
}

func NewGymRepo(db *pgx.Conn) *GymRepo {
    return &GymRepo{db: db}
}

func (r *GymRepo) CreateGym(gym *model.Gym) error {
	pstmt, err := r.db.Prepare("createGym", "INSERT INTO gyms (id, name, address, placeID, longitude, latitude) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
        log.Fatalf("Prepare failed: %v\n", err)
    }

	_, err =  r.db.Exec(pstmt.SQL, gym.ID, gym.Name, gym.Address, gym.PlaceID, gym.Longitude, gym.Latitude)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}
	return nil
}

