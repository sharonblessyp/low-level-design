package service

import (
	"parking_lot/models"
	"parking_lot/persistence"
)

type AdminService struct {
	FloorRepo       *persistence.FloorRepository
	ParkingSlotRepo *persistence.ParkingSlotRepository
}

func NewAdminService(floorRepo *persistence.FloorRepository, slotRepo *persistence.ParkingSlotRepository) *AdminService {
	return &AdminService{
		FloorRepo:       floorRepo,
		ParkingSlotRepo: slotRepo,
	}
}

func (as *AdminService) AddFloor(floorNum int, numSlots int) {
	floor := &models.Floor{
		ID: floorNum,
	}

	for i := 0; i < numSlots; i++ {
		slot := models.ParkingSlot{
			ID:         i + 1,
			FloorNum:   floorNum,
			IsOccupied: false,
		}
		floor.Slots = append(floor.Slots, slot)
	}

	as.FloorRepo.SaveFloor(floor)
}
