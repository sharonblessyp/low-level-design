package models

type ParkingSlot struct {
	ID         int
	FloorNum   int
	IsOccupied bool
	VehicleID  string
	NextID     int
}
