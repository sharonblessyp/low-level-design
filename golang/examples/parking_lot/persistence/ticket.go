package persistence

import (
	"fmt"
	"parking_lot/models"
	"sync"
	"time"
)

type TicketRepository struct {
	// mutex
	mutex   sync.Mutex
	tickets map[string]*models.Ticket
	nextID  int
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		tickets: make(map[string]*models.Ticket),
		nextID:  1,
	}
}

// save ticket
func (tr *TicketRepository) SaveTicket(ticket *models.Ticket) error {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	ticket.ID = tr.nextID
	tr.nextID++
	tr.tickets[ticket.VehicleID] = ticket
	return nil
}

// return the ticket by vehicle id
func (tr *TicketRepository) GetTicket(vehicleID string) *models.Ticket {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	return tr.tickets[vehicleID]
}

// deactivate the ticket by vehicle id
func (tr *TicketRepository) DeactivateTicket(vehicleID string) (models.Ticket, error) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	if ticket, exists := tr.tickets[vehicleID]; exists {
		ticket.ExitTime = time.Now()
		return *ticket, nil
	}
	return models.Ticket{}, fmt.Errorf("ticket not found for vehicle ID: %s", vehicleID)
}
