package controller

import (
	"parking_lot/models"
	"parking_lot/service"
)

type EntryController struct {
	ParkingService *service.ParkingService
}

func NewEntryController(vehicle models.Vehicle, parkingService *service.ParkingService) *EntryController {
	return &EntryController{
		ParkingService: parkingService,
	}
}
