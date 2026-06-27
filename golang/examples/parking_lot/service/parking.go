package service

import (
	"errors"
	"fmt"
	strategy "parking_lot/fee_strategy"
	"parking_lot/models"
	"parking_lot/persistence"
	"time"
)

type ParkingService struct {
	FeeService strategy.FeeStartegy
	TicketRepo *persistence.TicketRepository
	FloorRepo  *persistence.FloorRepository
	SlotRepo   *persistence.ParkingSlotRepository
}

// ticket service returns repo
func NewParkingService(ticketRepo *persistence.TicketRepository, slotRepo *persistence.ParkingSlotRepository, floorRepo *persistence.FloorRepository, pricingStrategy models.PricingStrategy) *ParkingService {
	feeStartegy := strategy.ChooseFeeStrategy(pricingStrategy)

	return &ParkingService{
		FeeService: feeStartegy,
		TicketRepo: ticketRepo,
		SlotRepo:   slotRepo,
		FloorRepo:  floorRepo,
	}
}

// park vehicle
func (ps *ParkingService) ParkVehicle(vehicleID string) (models.Ticket, error) {
	// allocate a slot for the vehicle
	slot, err := ps.AllocateSlot(vehicleID)
	if err != nil {
		// handle error
		fmt.Println("Error allocating slot:", err)
		return models.Ticket{}, err
	}

	// generate ticket
	fmt.Print(slot)
	ticket := ps.GenerateTicket(slot)
	err = ps.TicketRepo.SaveTicket(&ticket)
	if err != nil {
		return models.Ticket{}, nil
	}
	return ticket, nil
}

func (ps *ParkingService) AllocateSlot(vehicleID string) (models.ParkingSlot, error) {
	// allocate slot if available else send error
	floors := ps.FloorRepo.ListFloors()
	for _, floor := range floors {
		for _, slot := range floor.Slots {
			if !slot.IsOccupied {
				// allocate the slot to the vehicle
				slot.IsOccupied = true
				slot.VehicleID = vehicleID
				ps.SlotRepo.SaveSlot(&slot)
				return slot, nil
			}
		}
	}
	return models.ParkingSlot{}, errors.New("no available slot found")
}

// generate the ticket for given details of vehicle entry time and slot
func (ps *ParkingService) GenerateTicket(slot models.ParkingSlot) models.Ticket {
	ticket := models.Ticket{
		VehicleID: slot.VehicleID,
		EntryTime: time.Now(),
		SlotID:    slot.ID,
		Floor:     slot.FloorNum,
	}
	return ticket
}

func (ps *ParkingService) UnParkVehicle(vehicleID string, slotID int) (models.Receipt, error) {
	// release slot
	// deactivate the ticket
	// calculate fees
	// generate the receipt
	ps.SlotRepo.ReleaseSlot(slotID)
	ticket, err := ps.TicketRepo.DeactivateTicket(vehicleID)
	if err != nil {
		return models.Receipt{}, err
	}
	receipt := ps.FeeService.CalculateFee(ticket)
	return receipt, nil
}
