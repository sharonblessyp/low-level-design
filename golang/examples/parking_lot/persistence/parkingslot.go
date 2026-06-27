package persistence

import (
	"sync"

	"parking_lot/models"
)

type ParkingSlotRepository struct {
	// MUTEX
	mutex        sync.Mutex
	Parkingslots map[int]*models.ParkingSlot
}

func NewParkingSlotRepository() *ParkingSlotRepository {
	return &ParkingSlotRepository{
		Parkingslots: make(map[int]*models.ParkingSlot),
	}
}

func (psr *ParkingSlotRepository) SaveSlot(slot *models.ParkingSlot) error {
	psr.mutex.Lock()
	defer psr.mutex.Unlock()

	psr.Parkingslots[slot.ID] = slot
	return nil
}

func (psr *ParkingSlotRepository) GetSlot(slotID int) (*models.ParkingSlot, error) {
	psr.mutex.Lock()
	defer psr.mutex.Unlock()
	return psr.Parkingslots[slotID], nil
}

func (psr *ParkingSlotRepository) ReleaseSlot(slotID int) {
	psr.mutex.Lock()
	defer psr.mutex.Unlock()

	psr.Parkingslots[slotID].IsOccupied = false
	psr.Parkingslots[slotID].VehicleID = ""
	return
}
