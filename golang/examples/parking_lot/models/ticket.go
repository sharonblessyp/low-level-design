package models

import "time"

type Ticket struct {
	ID        int
	VehicleID string
	SlotID    int
	EntryTime time.Time
	ExitTime  time.Time
	Floor     int
}
