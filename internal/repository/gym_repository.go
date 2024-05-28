package repository

import (
	"goapirest/internal/model"

	"github.com/jackc/pgx"
)

type GymRepository interface {
    CreateGym(gym *model.GymData) error
}

type GymRepoImpl struct {
    db *pgx.Conn
}

func NewGymRepoImpl(db *pgx.Conn) *GymRepoImpl {
    return &GymRepoImpl{db: db}
}

func (r *GymRepoImpl) CreateGym(gym *model.GymData) error {
    // Implementação para criar uma academia no banco de dados
    // Exemplo:
    // _, err := r.db.Exec("INSERT INTO gyms (name, address, place_id, longitude, latitude) VALUES (?, ?, ?, ?, ?)", gym.Name, gym.Address, gym.PlaceID, gym.Longitude, gym.Latitude)
    // return err
    return nil // Substitua pelo código real de inserção
}
