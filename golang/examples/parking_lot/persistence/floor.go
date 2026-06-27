package persistence

import (
	"sync"

	"parking_lot/models"
)

type FloorRepository struct {
	// MUTEX
	mutex        sync.Mutex
	parkingslots map[int][]*models.Floor
}

/*
1. saveFloor
2. listAllFloors
*/
func NewFloorRepository() *FloorRepository {
	return &FloorRepository{
		parkingslots: make(map[int][]*models.Floor),
	}
}

// list floors
func (fr *FloorRepository) ListFloors() []*models.Floor {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	var floors []*models.Floor
	for _, floorList := range fr.parkingslots {
		floors = append(floors, floorList...)
	}
	return floors
}

func (fr *FloorRepository) SaveFloor(floor *models.Floor) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	fr.parkingslots[floor.ID] = append(fr.parkingslots[floor.ID], floor)
}
