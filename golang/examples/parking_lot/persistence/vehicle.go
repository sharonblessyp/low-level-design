package persistence

import (
	"parking_lot/models"
	"sync"
)

type VehicleRepository struct {
	// mutex
	mutex    sync.Mutex
	vehicles map[string]*models.Vehicle
}

func NewVehicleRepository() *VehicleRepository {
	return &VehicleRepository{
		vehicles: make(map[string]*models.Vehicle),
	}
}
