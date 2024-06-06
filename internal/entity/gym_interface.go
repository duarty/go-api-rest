package entity

type GymRepositoryInterface interface {
	Save(gym *Gym) error
	//	Get()
	// GetTotal() (int, error)
}
