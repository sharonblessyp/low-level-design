package main

import (
	"fmt"
	"parking_lot/controller"
	"parking_lot/models"
	"parking_lot/persistence"
	"parking_lot/service"
	"time"
)

func main() {
	// repositories
	slotrepo := persistence.NewParkingSlotRepository()
	ticketrepo := persistence.NewTicketRepository()
	floorRepo := persistence.NewFloorRepository()

	// service layer - repositories are created internally for the service
	parkingService := service.NewParkingService(ticketrepo, slotrepo, floorRepo, models.Hourly)

	/* controllers
	entry controllers
	1. finds a free parking slot
	2. if available adds vehicle to the slot
	3. returns the tickets to the user
	*/

	adminService := service.NewAdminService(floorRepo, slotrepo)
	adminService.AddFloor(1, 10) // Add a floor with 10 slots

	vehicle := models.Vehicle{
		ID:   "ABC123",
		Type: models.Car,
	}
	entryController := controller.NewEntryController(vehicle, parkingService)
	ticket, err := entryController.ParkingService.ParkVehicle(vehicle.ID)
	if err != nil {
		fmt.Println("Error parking vehicle:", err)
		return
	}
	fmt.Println("Ticket generated:", ticket)

	time.Sleep(5 * time.Second)
	// unpark vehicle
	receipt, err := entryController.ParkingService.UnParkVehicle(vehicle.ID, ticket.SlotID)
	if err != nil {
		fmt.Println("Error parking vehicle:", err)
		return
	}
	fmt.Println("Receipt generated:", receipt)

}
